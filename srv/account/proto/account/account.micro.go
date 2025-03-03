// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/account/proto/account/account.proto

package accountsrv

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	Exist(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserExistResponse, error)
	List(ctx context.Context, in *UserListQuery, opts ...client.CallOption) (*UserListResponse, error)
	Get(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	Create(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	Update(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	Delete(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "accountsrv"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Exist(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserExistResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Exist", in)
	out := new(UserExistResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) List(ctx context.Context, in *UserListQuery, opts ...client.CallOption) (*UserListResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.List", in)
	out := new(UserListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Get(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Get", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Create(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Create", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Update(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Update", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Exist(context.Context, *UserRequest, *UserExistResponse) error
	List(context.Context, *UserListQuery, *UserListResponse) error
	Get(context.Context, *UserRequest, *UserResponse) error
	Create(context.Context, *UserRequest, *UserResponse) error
	Update(context.Context, *UserRequest, *UserResponse) error
	Delete(context.Context, *UserRequest, *UserResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Exist(ctx context.Context, in *UserRequest, out *UserExistResponse) error
		List(ctx context.Context, in *UserListQuery, out *UserListResponse) error
		Get(ctx context.Context, in *UserRequest, out *UserResponse) error
		Create(ctx context.Context, in *UserRequest, out *UserResponse) error
		Update(ctx context.Context, in *UserRequest, out *UserResponse) error
		Delete(ctx context.Context, in *UserRequest, out *UserResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Exist(ctx context.Context, in *UserRequest, out *UserExistResponse) error {
	return h.UserServiceHandler.Exist(ctx, in, out)
}

func (h *userServiceHandler) List(ctx context.Context, in *UserListQuery, out *UserListResponse) error {
	return h.UserServiceHandler.List(ctx, in, out)
}

func (h *userServiceHandler) Get(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *userServiceHandler) Create(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *userServiceHandler) Update(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.Update(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}

// Client API for ProfileService service

type ProfileService interface {
	List(ctx context.Context, in *ProfileListQuery, opts ...client.CallOption) (*ProfileListResponse, error)
	Get(ctx context.Context, in *ProfileRequest, opts ...client.CallOption) (*ProfileResponse, error)
	Create(ctx context.Context, in *ProfileRequest, opts ...client.CallOption) (*ProfileResponse, error)
}

type profileService struct {
	c    client.Client
	name string
}

func NewProfileService(name string, c client.Client) ProfileService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "accountsrv"
	}
	return &profileService{
		c:    c,
		name: name,
	}
}

func (c *profileService) List(ctx context.Context, in *ProfileListQuery, opts ...client.CallOption) (*ProfileListResponse, error) {
	req := c.c.NewRequest(c.name, "ProfileService.List", in)
	out := new(ProfileListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileService) Get(ctx context.Context, in *ProfileRequest, opts ...client.CallOption) (*ProfileResponse, error) {
	req := c.c.NewRequest(c.name, "ProfileService.Get", in)
	out := new(ProfileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileService) Create(ctx context.Context, in *ProfileRequest, opts ...client.CallOption) (*ProfileResponse, error) {
	req := c.c.NewRequest(c.name, "ProfileService.Create", in)
	out := new(ProfileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProfileService service

type ProfileServiceHandler interface {
	List(context.Context, *ProfileListQuery, *ProfileListResponse) error
	Get(context.Context, *ProfileRequest, *ProfileResponse) error
	Create(context.Context, *ProfileRequest, *ProfileResponse) error
}

func RegisterProfileServiceHandler(s server.Server, hdlr ProfileServiceHandler, opts ...server.HandlerOption) error {
	type profileService interface {
		List(ctx context.Context, in *ProfileListQuery, out *ProfileListResponse) error
		Get(ctx context.Context, in *ProfileRequest, out *ProfileResponse) error
		Create(ctx context.Context, in *ProfileRequest, out *ProfileResponse) error
	}
	type ProfileService struct {
		profileService
	}
	h := &profileServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProfileService{h}, opts...))
}

type profileServiceHandler struct {
	ProfileServiceHandler
}

func (h *profileServiceHandler) List(ctx context.Context, in *ProfileListQuery, out *ProfileListResponse) error {
	return h.ProfileServiceHandler.List(ctx, in, out)
}

func (h *profileServiceHandler) Get(ctx context.Context, in *ProfileRequest, out *ProfileResponse) error {
	return h.ProfileServiceHandler.Get(ctx, in, out)
}

func (h *profileServiceHandler) Create(ctx context.Context, in *ProfileRequest, out *ProfileResponse) error {
	return h.ProfileServiceHandler.Create(ctx, in, out)
}
