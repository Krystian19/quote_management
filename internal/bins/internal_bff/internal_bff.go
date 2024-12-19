package internal_bff

import (
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
