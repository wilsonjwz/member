// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type UserServiceHandler interface {
	ClientManualAuthorize(context.Context, *ManualAuthorizeRequest) (*ManualAuthorizeReply, error)
}

func NewUserServiceHandler(srv UserServiceHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/v1/user", func(w http.ResponseWriter, r *http.Request) {
		var in ManualAuthorizeRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ClientManualAuthorize(ctx, req.(*ManualAuthorizeRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ManualAuthorizeReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	return r
}

type UserServiceHTTPClient interface {
	ClientManualAuthorize(ctx context.Context, req *ManualAuthorizeRequest, opts ...http1.CallOption) (rsp *ManualAuthorizeReply, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http1.Client
}

func NewUserServiceHTTPClient(client *http1.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) ClientManualAuthorize(ctx context.Context, in *ManualAuthorizeRequest, opts ...http1.CallOption) (out *ManualAuthorizeReply, err error) {
	path := binding.EncodePath("POST", "/v1/user", in)
	out = &ManualAuthorizeReply{}

	err = c.cc.Invoke(ctx, path, in, &out, http1.Method("POST"), http1.PathPattern("/v1/user"))

	return
}