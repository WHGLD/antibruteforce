package server

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/WHGLD/antibruteforce/internal/app"
	"github.com/golang/protobuf/ptypes/empty"
)

type Server struct {
	UnimplementedABruteforceServer
	app        app.App
	addr       string
	grpcServer *grpc.Server
}

func NewServer(a app.App, addr string) *Server {
	return &Server{app: a, addr: addr}
}

func (s *Server) Start(ctx context.Context) error {
	s.grpcServer = grpc.NewServer()
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	RegisterABruteforceServer(s.grpcServer, s)

	err = s.grpcServer.Serve(lis)

	<-ctx.Done()

	return err
}

func (s *Server) Stop(ctx context.Context) error { // todo если timeout работает убрать от сюда ctx
	s.grpcServer.Stop()
	return nil
}

func (s *Server) Auth(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {
	ok, err := s.app.CheckRateLimits(ctx, req.GetLogin(), req.GetPassword(), req.GetIp())
	if err != nil {
		return nil, err
	}

	// todo check return

	return &AuthResponse{Ok: ok}, nil
}

func (s *Server) Reset(ctx context.Context, req *ResetRequest) (*empty.Empty, error) {
	err := s.app.Reset(ctx, req.GetLogin(), req.GetPassword(), req.GetIp())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *Server) AddToWhiteList(ctx context.Context, req *AddNetMaskRequest) (*AddNetMaskResponse, error) {
	err := s.app.AddToWhiteList(ctx, req.GetIp(), req.GetMask())
	if err != nil {
		return nil, err
	}

	return &AddNetMaskResponse{Ip: req.GetIp(), Mask: req.GetMask()}, nil
}

func (s *Server) AddToBlackList(ctx context.Context, req *AddNetMaskRequest) (*AddNetMaskResponse, error) {
	err := s.app.AddToBlackList(ctx, req.GetIp(), req.GetMask())
	if err != nil {
		return nil, err
	}

	return &AddNetMaskResponse{Ip: req.GetIp(), Mask: req.GetMask()}, nil
}

func (s *Server) RemoveFromWhiteList(ctx context.Context, req *RemoveNetMaskRequest) (*empty.Empty, error) {
	err := s.app.RemoveFromWhiteList(ctx, req.GetIp(), req.GetMask())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *Server) RemoveFromBlackList(ctx context.Context, req *RemoveNetMaskRequest) (*empty.Empty, error) {
	err := s.app.RemoveFromBlackList(ctx, req.GetIp(), req.GetMask())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
