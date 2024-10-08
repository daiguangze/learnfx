// Code generated by Kitex v0.11.3. DO NOT EDIT.

package learnfxservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	item "learnfx/kitex_gen/learn/fx/item"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateItem(ctx context.Context, req *item.CreateItemReq, callOptions ...callopt.Option) (r *item.CreateItemResp, err error)
	FilterVisibleItems(ctx context.Context, req *item.FilterVisibleItemsReq, callOptions ...callopt.Option) (r *item.FilterVisibleItemsResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kLearnFxServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kLearnFxServiceClient struct {
	*kClient
}

func (p *kLearnFxServiceClient) CreateItem(ctx context.Context, req *item.CreateItemReq, callOptions ...callopt.Option) (r *item.CreateItemResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateItem(ctx, req)
}

func (p *kLearnFxServiceClient) FilterVisibleItems(ctx context.Context, req *item.FilterVisibleItemsReq, callOptions ...callopt.Option) (r *item.FilterVisibleItemsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FilterVisibleItems(ctx, req)
}
