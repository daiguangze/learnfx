package main

import (
	"context"
	"learnfx/domain/entity"
	"learnfx/domain/service"
)

type LearnFxServiceImpl struct {
	ItemDomainService service.ItemDomainService
}

func NewLearnFxServiceImpl(s service.ItemDomainService) service.ItemDomainService {
	return &LearnFxServiceImpl{
		ItemDomainService: s,
	}
}

func (i *LearnFxServiceImpl) FilterVisibleItems(ctx context.Context, req *item.FilterVisibleItemsReq) (*item.FilterVisibleItemsResp, error) {
	visibleItemIDs, err := i.ItemDomainService.FilterVisibleItems(ctx, req.ItemIDs, req.UserID)
	if err != nil {
		return nil, err
	}
	resp := &item.FilterVisibleItemsResp{
		VisibleItemIDs: visibleItemIDs,
	}
	return resp, nil
}

func (i *LearnFxServiceImpl) CreateItem(ctx context.Context, req *item.CreateItemReq) (resp *item.CreateItemResp, err error) {
	itemID, err := i.ItemDomainService.CreateItem(ctx, &entity.Item{
		ID:           req.Item.ID,
		Name:         req.Item.Name,
		Desc:         req.Item.Desc,
		VisibleUsers: req.Item.VisibleUsers,
	})
	if err != nil {
		return nil, err
	}
	resp = &item.CreateItemResp{
		ID: itemID,
	}
	return resp, nil
}
