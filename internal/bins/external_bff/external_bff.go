package external_bff

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

const (
	// endpoint used to receive graphql requests
	GQLEndpoint = "/q"
)

type Server struct {
	port string

	resolver resolver
}

type Opts struct {
	Port string

	// should be disabled by default
	AllowPlayground bool

	DB *db.DB
}

func New(opts Opts) Server {
	return Server{
		port: opts.Port,

		resolver: newResolver(resolverOpts{
			DB: opts.DB,
		}),
	}
}

func (s *Server) Run() error {
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},
	}).Handler)

	gqlSrv := handler.NewDefaultServer(NewExecutableSchema(Config{
		Resolvers: s.resolver,
	}))

	gqlSrv.Use(extension.Introspection{})
	gqlSrv.Use(apollotracing.Tracer{})

	router.Handle("/", playground.Handler("Playground", GQLEndpoint))
	router.Handle(GQLEndpoint, gqlSrv)

	return http.ListenAndServe(":"+s.port, router)
}

type resolver struct {
	db *db.DB
}

type resolverOpts struct {
	DB *db.DB
}

func newResolver(opts resolverOpts) resolver {
	return resolver{
		db: opts.DB,
	}
}

type queryResolver struct{ resolver }
type mutationResolver struct{ resolver }

func (r resolver) Query() QueryResolver       { return queryResolver{r} }
func (r resolver) Mutation() MutationResolver { return mutationResolver{r} }
