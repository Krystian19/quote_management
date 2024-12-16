package external_bff

import (
	"github.com/Krystian19/quote_management/internal/libs/db"
)

const (
	// endpoint used to receive graphql requests
	GQLEndpoint = "/q"
)

type Server struct {
	allowPlayground bool
	port            string
}

type Opts struct {
	Port string

	// should be disabled by default
	AllowPlayground bool

	DB *db.DB
}
