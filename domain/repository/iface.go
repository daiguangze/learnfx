package repository

import (
	"context"
	"learnfx/domain/entity"
)


type ItemRepo interface{
	CreateItem(ctx context.Context, item *entity.Item) (int64, error)
	BatchQueryItemByID(ctx context.Context, itemIDs []int64) ([]*entity.Item, error)
}
