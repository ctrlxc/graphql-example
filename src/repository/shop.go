package repository

import (
	"context"
	"app/models"
)

func (r *Repository) ShopByID(ctx context.Context, id int64) (*models.Shop, error) {
	return models.FindShop(ctx, r.db, id)
}

func (r *Repository) ShopsByIDs(ctx context.Context, ids []int64) ([]*models.Shop, error) {
	return models.Shops(models.ShopWhere.ID.IN(ids)).All(ctx, r.db)
}
