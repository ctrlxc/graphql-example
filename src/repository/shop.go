package repository

import (
	"app/models"
	"app/pagination"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

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
