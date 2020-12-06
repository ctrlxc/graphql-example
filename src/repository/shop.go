package repository

import (
	"app/models"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (r *Repository) ShopByID(ctx context.Context, id int64) (*models.Shop, error) {
	return models.FindShop(ctx, r.db, id)
}

func (r *Repository) ShopsByIDs(ctx context.Context, ids []int64) ([]*models.Shop, error) {
	return models.Shops(models.ShopWhere.ID.IN(ids)).All(ctx, r.db)
}

func (r *Repository) ShopsByName(ctx context.Context, name string) ([]*models.Shop, error) {
	return models.Shops(
		qm.Where(fmt.Sprintf("%s like ?", models.ShopColumns.ShopName), fmt.Sprintf("%%%s%%", name)),
	).All(ctx, r.db)
}
