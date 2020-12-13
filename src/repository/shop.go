package repository

import (
	"app/models"
	"context"
)

func (r *Repository) ShopByID(ctx context.Context, id int64) (*models.Shop, error) {
	return models.FindShop(ctx, r.Db, id)
}

func (r *Repository) ShopsByIDs(ctx context.Context, ids []int64) ([]*models.Shop, error) {
	return models.Shops(models.ShopWhere.ID.IN(ids)).All(ctx, r.Db)
}
