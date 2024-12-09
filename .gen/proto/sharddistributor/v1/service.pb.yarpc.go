// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/sharddistributor/v1/service.proto

package sharddistributorv1

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/api/x/restriction"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// ShardDistributorAPIYARPCClient is the YARPC client-side interface for the ShardDistributorAPI service.
type ShardDistributorAPIYARPCClient interface {
	GetShardOwner(context.Context, *GetShardOwnerRequest, ...yarpc.CallOption) (*GetShardOwnerResponse, error)
}

func newShardDistributorAPIYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) ShardDistributorAPIYARPCClient {
	return &_ShardDistributorAPIYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.cadence.sharddistributor.v1.ShardDistributorAPI",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewShardDistributorAPIYARPCClient builds a new YARPC client for the ShardDistributorAPI service.
func NewShardDistributorAPIYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) ShardDistributorAPIYARPCClient {
	return newShardDistributorAPIYARPCClient(clientConfig, nil, options...)
}

// ShardDistributorAPIYARPCServer is the YARPC server-side interface for the ShardDistributorAPI service.
type ShardDistributorAPIYARPCServer interface {
	GetShardOwner(context.Context, *GetShardOwnerRequest) (*GetShardOwnerResponse, error)
}

type buildShardDistributorAPIYARPCProceduresParams struct {
	Server      ShardDistributorAPIYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildShardDistributorAPIYARPCProcedures(params buildShardDistributorAPIYARPCProceduresParams) []transport.Procedure {
	handler := &_ShardDistributorAPIYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "uber.cadence.sharddistributor.v1.ShardDistributorAPI",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "GetShardOwner",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.GetShardOwner,
							NewRequest:  newShardDistributorAPIServiceGetShardOwnerYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildShardDistributorAPIYARPCProcedures prepares an implementation of the ShardDistributorAPI service for YARPC registration.
func BuildShardDistributorAPIYARPCProcedures(server ShardDistributorAPIYARPCServer) []transport.Procedure {
	return buildShardDistributorAPIYARPCProcedures(buildShardDistributorAPIYARPCProceduresParams{Server: server})
}

// FxShardDistributorAPIYARPCClientParams defines the input
// for NewFxShardDistributorAPIYARPCClient. It provides the
// paramaters to get a ShardDistributorAPIYARPCClient in an
// Fx application.
type FxShardDistributorAPIYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxShardDistributorAPIYARPCClientResult defines the output
// of NewFxShardDistributorAPIYARPCClient. It provides a
// ShardDistributorAPIYARPCClient to an Fx application.
type FxShardDistributorAPIYARPCClientResult struct {
	fx.Out

	Client ShardDistributorAPIYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxShardDistributorAPIYARPCClient provides a ShardDistributorAPIYARPCClient
// to an Fx application using the given name for routing.
//
//	fx.Provide(
//	  sharddistributorv1.NewFxShardDistributorAPIYARPCClient("service-name"),
//	  ...
//	)
func NewFxShardDistributorAPIYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxShardDistributorAPIYARPCClientParams) FxShardDistributorAPIYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxShardDistributorAPIYARPCClientResult{
			Client: newShardDistributorAPIYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxShardDistributorAPIYARPCProceduresParams defines the input
// for NewFxShardDistributorAPIYARPCProcedures. It provides the
// paramaters to get ShardDistributorAPIYARPCServer procedures in an
// Fx application.
type FxShardDistributorAPIYARPCProceduresParams struct {
	fx.In

	Server      ShardDistributorAPIYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxShardDistributorAPIYARPCProceduresResult defines the output
// of NewFxShardDistributorAPIYARPCProcedures. It provides
// ShardDistributorAPIYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxShardDistributorAPIYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxShardDistributorAPIYARPCProcedures provides ShardDistributorAPIYARPCServer procedures to an Fx application.
// It expects a ShardDistributorAPIYARPCServer to be present in the container.
//
//	fx.Provide(
//	  sharddistributorv1.NewFxShardDistributorAPIYARPCProcedures(),
//	  ...
//	)
func NewFxShardDistributorAPIYARPCProcedures() interface{} {
	return func(params FxShardDistributorAPIYARPCProceduresParams) FxShardDistributorAPIYARPCProceduresResult {
		return FxShardDistributorAPIYARPCProceduresResult{
			Procedures: buildShardDistributorAPIYARPCProcedures(buildShardDistributorAPIYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: ShardDistributorAPIReflectionMeta,
		}
	}
}

// ShardDistributorAPIReflectionMeta is the reflection server metadata
// required for using the gRPC reflection protocol with YARPC.
//
// See https://github.com/grpc/grpc/blob/master/doc/server-reflection.md.
var ShardDistributorAPIReflectionMeta = reflection.ServerMeta{
	ServiceName:     "uber.cadence.sharddistributor.v1.ShardDistributorAPI",
	FileDescriptors: yarpcFileDescriptorClosure0055bfd59dff1f95,
}

type _ShardDistributorAPIYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_ShardDistributorAPIYARPCCaller) GetShardOwner(ctx context.Context, request *GetShardOwnerRequest, options ...yarpc.CallOption) (*GetShardOwnerResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "GetShardOwner", request, newShardDistributorAPIServiceGetShardOwnerYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*GetShardOwnerResponse)
	if !ok {
		return nil, protobuf.CastError(emptyShardDistributorAPIServiceGetShardOwnerYARPCResponse, responseMessage)
	}
	return response, err
}

type _ShardDistributorAPIYARPCHandler struct {
	server ShardDistributorAPIYARPCServer
}

func (h *_ShardDistributorAPIYARPCHandler) GetShardOwner(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *GetShardOwnerRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*GetShardOwnerRequest)
		if !ok {
			return nil, protobuf.CastError(emptyShardDistributorAPIServiceGetShardOwnerYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.GetShardOwner(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newShardDistributorAPIServiceGetShardOwnerYARPCRequest() proto.Message {
	return &GetShardOwnerRequest{}
}

func newShardDistributorAPIServiceGetShardOwnerYARPCResponse() proto.Message {
	return &GetShardOwnerResponse{}
}

var (
	emptyShardDistributorAPIServiceGetShardOwnerYARPCRequest  = &GetShardOwnerRequest{}
	emptyShardDistributorAPIServiceGetShardOwnerYARPCResponse = &GetShardOwnerResponse{}
)

var yarpcFileDescriptorClosure0055bfd59dff1f95 = [][]byte{
	// uber/cadence/sharddistributor/v1/service.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x4d, 0x4a, 0x2d,
		0xd2, 0x4f, 0x4e, 0x4c, 0x49, 0xcd, 0x4b, 0x4e, 0xd5, 0x2f, 0xce, 0x48, 0x2c, 0x4a, 0x49, 0xc9,
		0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a, 0x2d, 0xc9, 0x2f, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0x4e, 0x2d,
		0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x00, 0xa9, 0xd7, 0x83,
		0xaa, 0xd7, 0x43, 0x57, 0xaf, 0x57, 0x66, 0xa8, 0x14, 0xc8, 0x25, 0xe2, 0x9e, 0x5a, 0x12, 0x0c,
		0x92, 0xf1, 0x2f, 0xcf, 0x4b, 0x2d, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe6,
		0xe2, 0x04, 0x2b, 0x8f, 0xcf, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x00,
		0x0b, 0x78, 0xa7, 0x56, 0x0a, 0xc9, 0x70, 0x71, 0xe6, 0x25, 0xe6, 0xa6, 0x16, 0x17, 0x24, 0x26,
		0xa7, 0x4a, 0x30, 0x81, 0x25, 0x11, 0x02, 0x4a, 0xde, 0x5c, 0xa2, 0x68, 0x46, 0x16, 0x17, 0xe4,
		0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x70, 0xb1, 0xe6, 0x83, 0x04, 0xa0, 0xe6, 0x41, 0x38, 0xf8, 0x0d,
		0x33, 0x9a, 0xc1, 0xc8, 0x25, 0x0c, 0x36, 0xca, 0x05, 0xe1, 0x6e, 0xc7, 0x00, 0x4f, 0xa1, 0x06,
		0x46, 0x2e, 0x5e, 0x14, 0x5b, 0x84, 0xcc, 0xf4, 0x08, 0x79, 0x56, 0x0f, 0x9b, 0x4f, 0xa5, 0xcc,
		0x49, 0xd6, 0x07, 0xf1, 0x8e, 0x93, 0x77, 0x94, 0x67, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
		0x72, 0x7e, 0xae, 0x3e, 0x4a, 0xcc, 0xe8, 0xa5, 0xa7, 0xe6, 0xe9, 0x83, 0xa3, 0x00, 0x5b, 0x24,
		0x59, 0xa3, 0x8b, 0x95, 0x19, 0x26, 0xb1, 0x81, 0x55, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
		0x55, 0x06, 0xa6, 0x2a, 0xe2, 0x01, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) ShardDistributorAPIYARPCClient {
			return NewShardDistributorAPIYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}