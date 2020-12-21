# graphql-example

## GraphQLプロジェクトの作成

GraphQLサーバーは[gqlgen](https://github.com/99designs/gqlgen)を使用して作成します。  
[gqlgen - Getting Started](https://gqlgen.com/getting-started/) の通りにプロジェクトを作成します。

### Goモジュールのプロジェクトを作成

```sh
cd /go
mkdir src
cd /go/src
mkdir app
cd app
go mod init app
```

### gqlgenでプロジェクトのスケルトンを生成

```sh
go get -u github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init
```

gqlgen推奨のプロジェクトのスケルトンが作成されます。

```text
app
├── go.mod
├── go.sum
├── gqlgen.yml               - gqlgenの設定ファイル。後で編集していきます
├── graph
│   ├── generated            - `schema.graphql`からクエリのパース処理などを行うruntimeが自動生成されます。`DO NOT EDIT.`
│   │   └── generated.go
│   ├── model                - `schema.graphql`からGraphQLのモデルが自動生成されます。 `DO NOT EDIT.`
│   │   └── models_gen.go
│   ├── resolver.go          - リゾルバー本体です。ここを主に実装していきます。
│   ├── schema.graphqls      - ここに`Schema`を定義していきます。
│   └── schema.resolvers.go  - `schema.graphql`からエンドポイントのスケルトンが自動生成されます。
└── server.go                - エントリーポイント。個人的な好みでmain.goにリネームしました。
```

個人的に`server.go`より`main.go`の方がエントリーポイントとして分かりやすいと考えたのでリネームしました。

```
mv server.go main.go
```

### Schemaを定義

`graph/schema.graphqls` にスキーマを定義していきます。

### GraphQLのモデルとリゾルバーの生成

```sh
gqlgen generate
```

`schema.resolvers.go` にエンドポイントが自動生成されます。

以下は `type Query { node(id: ID!): Node! }` に対して生成されたエンドポイントの例です。  
`panic(fmt.Errorf("not implemented"))` の部分を本来の実装に置き換えていきます。

```go
func (r *queryResolver) Node(ctx context.Context, id string) ([]*model.Node, error) {
	panic(fmt.Errorf("not implemented"))
}
```

`schema.resolvers.go` はエンドポイントとして半自動生成されるのでシンプルな状態を保ちたいです。  
そのため `resolver.go` でデータ取得やデータ加工などのリゾルバーとしての本来の実装を行います。

- schema.resolvers.go  - 半自動生成されるのでシンプルな実装を保つ。本来の処理は`resolver.go`に移譲する
- resolver.go          - リゾルバー本体。データ取得やデータ加工などのリゾルバーとしての本来の実装は主にここで行う

以下のような実装になります。

`graph/schema.resolvers.go`

```diff
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
-	panic(fmt.Errorf("not implemented"))
+	return r.node(ctx, id) // resolver.go に処理を移譲
}
```

`graph/resolvers.go`

```go
func (r *Resolver) node(ctx context.Context, id string) (model.Node, error) {
	// データ取得やデータ加工などをして　`mode.Node` を生成して返す
	return model.Node{}, nil
}
```

## データベースへのアクセス

ORMの[sqlboiler](https://github.com/volatiletech/sqlboiler)を使用します。

データベースドリブンでORMコードを自動生成するため  
データベースやテーブルは事前に作成しておいてください。

### データベースからORMコードを自動生成

`sqlboiler.toml` にDB接続情報を設定します。

```toml
[psql]
  dbname = "graphql"
  host   = "postgres"
  port   = 5432
  user   = "graphql"
  pass   = "graphql"
  sslmode = "disable"
  schema = "graphql"
```

以下のコマンドでORMコードを自動生成します。

```sh
go get -u github.com/lib/pq
go get -u github.com/volatiletech/sqlboiler
go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
sqlboiler --wipe --no-tests psql
```

```text
app
└── models
    └── *.go     - `models`以下にORM用のモデルが自動生成されます。`DO NOT EDIT.`
```

### リポジトリパターン

クエリの構築などデータベース固有の処理がアプリケーションのコード中に散在しないようにリポジトリパターンを採用します。  
sqlboilerで自動生成されたデータ取得／更新のAPIの呼び出しは基本的には全てリポジトリを経由して行います。

```
アプリケーションコード -> リポジトリコード -> sqlboilerの自動生成コード -> データベース
```

以下にアプリケーションで必要なデータベースへのアクセス用のAPIを追加します。

```txt
app
└── repository
    ├── repository.go  - Repositoryの共通関数を定義します。
    └── *.go           - テーブルやエンティティ毎に分けて作成します。
```

以下のような実装になります。

```go
func (r *Repository) ShopByID(ctx context.Context, id int64) (*models.Shop, error) {
	return models.FindShop(ctx, r.db, id)
}

func (r *Repository) ShopsByIDs(ctx context.Context, ids []int64) ([]*models.Shop, error) {
	return models.Shops(models.ShopWhere.ID.IN(ids)).All(ctx, r.db)
}

func (r *Repository) ShopsByName(ctx context.Context, name string, paginator *pagination.Paginator) ([]*models.Shop, error) {
	condition := qm.Where(fmt.Sprintf("%s like ?", models.ShopColumns.ShopName), fmt.Sprintf("%%%s%%", name))

	if paginator != nil {
		return models.Shops(paginator.Queries(condition)...).All(ctx, r.db)
	}

	return models.Shops(condition).All(ctx, r.db)
}

func (r *Repository) ShopsCountByName(ctx context.Context, name string) (int64, error) {
	condition := qm.Where(fmt.Sprintf("%s like ?", models.ShopColumns.ShopName), fmt.Sprintf("%%%s%%", name))
	return models.Shops(condition).Count(ctx, r.db)
}
```

## DataLoader

N+1問題への対応として[graph-gophers/dataloader](https://github.com/graph-gophers/dataloader)を使用します。  

### DataLoaderの作成

以下にアプリケーションで必要なLoaderを追加します。

```txt
app
└── loader
    ├── loader.go   - Loaderの共通関数を定義します。
    └── *.go        - テーブルやエンティティ毎に分けて作成します。
```

以下のような実装になります。

```go
package loader

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"app/models"
	"app/repository"

	"github.com/graph-gophers/dataloader"
)

const shopLoaderKey = "shopLoader"

type shopIDKey struct {
	id int64
}

func (key shopIDKey) String() string {
	return fmt.Sprintf("%s/%v", reflect.TypeOf(key).Name(), key.id) // should be global unique
}

func (key shopIDKey) Raw() interface{} {
	return key.id
}

func LoadShop(ctx context.Context, id int64) (*models.Shop, error) {
	ldr, err := getLoader(ctx, shopLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, shopIDKey{id: id})()
	if err != nil {
		return nil, err
	}

	return data.(*models.Shop), nil
}

func LoadShops(ctx context.Context, ids []int64) ([]*models.Shop, error) {
	ldr, err := getLoader(ctx, shopLoaderKey)
	if err != nil {
		return nil, err
	}

	shopIDs := make(dataloader.Keys, len(ids))
	for i, id := range ids {
		shopIDs[i] = shopIDKey{id: id}
	}

	datas, errs := ldr.LoadMany(ctx, shopIDs)()
	if len(errs) != 0 {
		return nil, errs[0]
	}

	shops := make([]*models.Shop, len(datas))
	for i, data := range datas {
		shops[i] = data.(*models.Shop)
	}

	return shops, nil
}

func newShopLoader(repo *repository.Repository) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		shopIDs := make([]int64, len(keys))

		for i, key := range keys {
			shopIDs[i] = key.(shopIDKey).id
		}

		shops, _ := repo.ShopsByIDs(ctx, shopIDs)

		for i, key := range keys {
			results[i] = &dataloader.Result{Data: nil, Error: nil}

			for _, shop := range shops {
				if key.(shopIDKey).id == shop.ID {
					results[i].Data = shop
					continue
				}
			}

			if results[i].Data == nil {
				results[i].Error = errors.New("Shop not found")
			}
		}

		return results
	}
}
```

### DataLoaderのインスタンスの生成

作成したDataLoaderのインスタンスを以下のように生成します。

`loader/loader.go`

```go
func NewLoaders(repo *repository.Repository) *Loaders {
	return &Loaders{
		batchFuncs: map[string]dataloader.BatchFunc{
			// ここに作成したDataLoaderのインスタンスを追加していきます。
			shopLoaderKey: newShopLoader(repo),
			bookLoaderKey: newBookLoader(repo),
		},
	}
}
```

### リクエスト単位でキャッシュ

[リクエスト単位のキャッシュ](https://github.com/graph-gophers/dataloader#cache)とするため  
MiddlewareでcontextにDataLoaderのインスタンスを保持させます。　　

`main.go` 

```go
// Middleware for attaching data loaders for GraphQL
func loaderMiddleware(next http.Handler, repo *repository.Repository) http.Handler {
	loaders := loader.NewLoaders(repo)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(loaders.Attach(r.Context())))
	})
}
```

### ページネーション

カーソルベースのページネーションを行います。  
ライブラリは使用せずに独自に実装しました。

```txt
app
└── pagination
    ├── cursor.go         - カーソルのエンコード／デコード(base64)を行います。
    └── pagination.go     - ページネーションを行うための`sqlboiler｀のクエリを生成します。
```

以下のような使い方になります。

```go
// graph/resolver.go
func (r *Resolver) shops(ctx context.Context, after *string, before *string, first *int, last *int, query string, orderBy []*model.ShopOrder) (*model.ShopConnection, error) {
	paginator := pagination.NewPaginator(
		after,
		before,
		first,
		last,
		model.ShopOrderToPaginationOrders(orderBy),
	)

	shops, err := r.repo.ShopsByName(ctx, query, paginator)

	if err != nil {
		return nil, err
    }
    
    // ... 省略　...
}

// repository/shop.go
func (r *Repository) ShopsByName(ctx context.Context, name string, paginator *pagination.Paginator) ([]*models.Shop, error) {
	condition := qm.Where(fmt.Sprintf("%s like ?", models.ShopColumns.ShopName), fmt.Sprintf("%%%s%%", name))

	if paginator != nil {
		return models.Shops(paginator.Queries(condition)...).All(ctx, r.db)
	}

	return models.Shops(condition).All(ctx, r.db)
}
```

## 未対応

- 複雑なクエリの制限

https://gqlgen.com/reference/complexity/#limiting-query-complexity

- 認証

https://gqlgen.com/recipes/authentication/
