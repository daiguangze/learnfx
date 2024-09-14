// Code generated by Kitex v0.11.3. DO NOT EDIT.

package learnfxservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	item "learnfx/kitex_gen/learn/fx/item"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateItem": kitex.NewMethodInfo(
		createItemHandler,
		newLearnFxServiceCreateItemArgs,
		newLearnFxServiceCreateItemResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"FilterVisibleItems": kitex.NewMethodInfo(
		filterVisibleItemsHandler,
		newLearnFxServiceFilterVisibleItemsArgs,
		newLearnFxServiceFilterVisibleItemsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	learnFxServiceServiceInfo                = NewServiceInfo()
	learnFxServiceServiceInfoForClient       = NewServiceInfoForClient()
	learnFxServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return learnFxServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return learnFxServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return learnFxServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "LearnFxService"
	handlerType := (*item.LearnFxService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "item",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func createItemHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*item.LearnFxServiceCreateItemArgs)
	realResult := result.(*item.LearnFxServiceCreateItemResult)
	success, err := handler.(item.LearnFxService).CreateItem(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLearnFxServiceCreateItemArgs() interface{} {
	return item.NewLearnFxServiceCreateItemArgs()
}

func newLearnFxServiceCreateItemResult() interface{} {
	return item.NewLearnFxServiceCreateItemResult()
}

func filterVisibleItemsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*item.LearnFxServiceFilterVisibleItemsArgs)
	realResult := result.(*item.LearnFxServiceFilterVisibleItemsResult)
	success, err := handler.(item.LearnFxService).FilterVisibleItems(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLearnFxServiceFilterVisibleItemsArgs() interface{} {
	return item.NewLearnFxServiceFilterVisibleItemsArgs()
}

func newLearnFxServiceFilterVisibleItemsResult() interface{} {
	return item.NewLearnFxServiceFilterVisibleItemsResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateItem(ctx context.Context, req *item.CreateItemReq) (r *item.CreateItemResp, err error) {
	var _args item.LearnFxServiceCreateItemArgs
	_args.Req = req
	var _result item.LearnFxServiceCreateItemResult
	if err = p.c.Call(ctx, "CreateItem", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FilterVisibleItems(ctx context.Context, req *item.FilterVisibleItemsReq) (r *item.FilterVisibleItemsResp, err error) {
	var _args item.LearnFxServiceFilterVisibleItemsArgs
	_args.Req = req
	var _result item.LearnFxServiceFilterVisibleItemsResult
	if err = p.c.Call(ctx, "FilterVisibleItems", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
