package repository

import (
	"app/models"
	"app/pagination"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *Repository) BookByID(ctx context.Context, id int64) (*models.Book, error) {
	return models.FindBook(ctx, r.db, id)
}

func (r *Repository) BooksByIDs(ctx context.Context, ids []int64) ([]*models.Book, error) {
	return models.Books(models.BookWhere.ID.IN(ids)).All(ctx, r.db)
}

func (r *Repository) BooksByShopID(ctx context.Context, id int64) ([]*models.Book, error) {
	shops, _ := models.Shops(
		models.ShopWhere.ID.EQ(id),
		qm.Load(qm.Rels(models.ShopRels.Stocks, models.StockRels.Book)),
	).All(ctx, r.db)

	books := make([]*models.Book, 0)

	for _, shop := range shops {
		for _, stock := range shop.R.Stocks {
			books = append(books, stock.R.Book)
		}
	}

	return books, nil
}

func (r *Repository) BooksByShopIDs(ctx context.Context, ids []int64) (map[int64][]*models.Book, error) {
	shops, _ := models.Shops(
		models.ShopWhere.ID.IN(ids),
		qm.Load(qm.Rels(models.ShopRels.Stocks, models.StockRels.Book)),
	).All(ctx, r.db)

	books := make(map[int64][]*models.Book, 0)

	for _, shop := range shops {
		for _, stock := range shop.R.Stocks {
			books[shop.ID] = append(books[shop.ID], stock.R.Book)
		}
	}

	return books, nil
}

func (r *Repository) BoolsByTitle(ctx context.Context, title string, paginator *pagination.Paginator) ([]*models.Book, error) {
	condition := qm.Where(fmt.Sprintf("%s like ?", models.BookColumns.BookTitle), fmt.Sprintf("%%%s%%", title))

	if paginator != nil {
		return models.Books(paginator.Queries(condition)...).All(ctx, r.db)
	}

	return models.Books(condition).All(ctx, r.db)
}

func (r *Repository) BooksCountByTitle(ctx context.Context, title string) (int64, error) {
	condition := qm.Where(fmt.Sprintf("%s like ?", models.BookColumns.BookTitle), fmt.Sprintf("%%%s%%", title))
	return models.Books(condition).Count(ctx, r.db)
}
