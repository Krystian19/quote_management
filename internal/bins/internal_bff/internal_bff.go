package internal_bff

import (
	"fmt"
	"net"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/Krystian19/quote_management/proto"
	"google.golang.org/grpc"
)

type InternalBff struct {
	srv *grpc.Server
	lis net.Listener
}

type Opts struct {
	Port int
	DB   *db.DB
}

func New(opts Opts) (*InternalBff, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.Port))
	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer()
	proto.RegisterInternalBffAPIServer(srv, newServer(serverOpts{
		DB: opts.DB,
	}))

	return &InternalBff{
		srv: srv,
		lis: lis,
	}, nil
}

func (s *InternalBff) Run() error {
	return s.srv.Serve(s.lis)
}

type server struct {
	db *db.DB

	proto.UnimplementedInternalBffAPIServer
}

type serverOpts struct {
	DB *db.DB
}

func newServer(opts serverOpts) server {
	return server{
		db: opts.DB,
	}
}
