// Code generated by protoc-gen-triple. DO NOT EDIT.
//
// Source: comment_api.proto
package api

import (
	"context"
)

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/protocol/triple/triple_protocol"
	"dubbo.apache.org/dubbo-go/v3/server"
)

// This is a compile-time assertion to ensure that this generated file and the Triple package
// are compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of Triple newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of Triple or updating the Triple
// version compiled into your binary.
const _ = triple_protocol.IsAtLeastVersion0_1_0

const (
	// CommentName is the fully-qualified name of the Comment service.
	CommentName = "org.apache.dubbogo.samples.shop.comment.api.Comment"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CommentGetCommentProcedure is the fully-qualified name of the Comment's GetComment RPC.
	CommentGetCommentProcedure = "/org.apache.dubbogo.samples.shop.comment.api.Comment/GetComment"
)

var (
	_ Comment = (*CommentImpl)(nil)
)

// Comment is a client for the org.apache.dubbogo.samples.shop.comment.api.Comment service.
type Comment interface {
	GetComment(ctx context.Context, req *CommentReq, opts ...client.CallOption) (*CommentResp, error)
}

// NewComment constructs a client for the api.Comment service.
func NewComment(cli *client.Client, opts ...client.ReferenceOption) (Comment, error) {
	conn, err := cli.DialWithInfo("org.apache.dubbogo.samples.shop.comment.api.Comment", &Comment_ClientInfo, opts...)
	if err != nil {
		return nil, err
	}
	return &CommentImpl{
		conn: conn,
	}, nil
}

func SetConsumerService(srv common.RPCService) {
	dubbo.SetConsumerServiceWithInfo(srv, &Comment_ClientInfo)
}

// CommentImpl implements Comment.
type CommentImpl struct {
	conn *client.Connection
}

func (c *CommentImpl) GetComment(ctx context.Context, req *CommentReq, opts ...client.CallOption) (*CommentResp, error) {
	resp := new(CommentResp)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "GetComment", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

var Comment_ClientInfo = client.ClientInfo{
	InterfaceName: "org.apache.dubbogo.samples.shop.comment.api.Comment",
	MethodNames:   []string{"GetComment"},
	ConnectionInjectFunc: func(dubboCliRaw interface{}, conn *client.Connection) {
		dubboCli := dubboCliRaw.(*CommentImpl)
		dubboCli.conn = conn
	},
}

// CommentHandler is an implementation of the org.apache.dubbogo.samples.shop.comment.api.Comment service.
type CommentHandler interface {
	GetComment(context.Context, *CommentReq) (*CommentResp, error)
}

func RegisterCommentHandler(srv *server.Server, hdlr CommentHandler, opts ...server.ServiceOption) error {
	return srv.Register(hdlr, &Comment_ServiceInfo, opts...)
}

func SetProviderService(srv common.RPCService) {
	dubbo.SetProviderServiceWithInfo(srv, &Comment_ServiceInfo)
}

var Comment_ServiceInfo = server.ServiceInfo{
	InterfaceName: "org.apache.dubbogo.samples.shop.comment.api.Comment",
	ServiceType:   (*CommentHandler)(nil),
	Methods: []server.MethodInfo{
		{
			Name: "GetComment",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(CommentReq)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*CommentReq)
				res, err := handler.(CommentHandler).GetComment(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
	},
}
