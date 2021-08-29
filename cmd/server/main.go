package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"strings"

	"github.com/firdavsich/surl/pkg/storage"

	_ "github.com/lib/pq"

	"github.com/firdavsich/surl/pkg/api"
	"google.golang.org/grpc"
)

type GPRCServer struct {
	api.SurlServer
	db         *sql.DB
	listenPort string
	baseURL    string
}

func (s *GPRCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	result, err := storage.Add(s.db, req.GetUrl())
	log.Println(result)
	if err != nil {
		log.Print(err)
		return &api.CreateResponse{Shorturl: ""}, err
	}
	return &api.CreateResponse{Shorturl: s.baseURL + result}, nil
}

func (s *GPRCServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	hash := strings.TrimPrefix(req.GetShorturl(), s.baseURL)
	result, err := storage.Get(s.db, hash)
	log.Println(result)
	if err != nil {
		log.Print(err)
		return &api.GetResponse{Url: ""}, err
	}

	return &api.GetResponse{Url: result}, nil
}

func main() {

	var err error // ! FIXME: remove
	s := grpc.NewServer()
	srv := &GPRCServer{}

	// get application port from environment variable
	if len(os.Getenv("PORT")) == 0 {
		srv.listenPort = "8080"
	} else {
		srv.listenPort = os.Getenv("PORT")
	}
	if len(os.Getenv("BASE_URL")) == 0 {
		srv.baseURL = "http://localhost:8080/"
	} else {
		srv.baseURL = os.Getenv("BASE_URL")
	}

	// open connect to postgresql
	srv.db, err = sql.Open("postgres", "host=localhost port=5432 user=surl dbname=surl_db sslmode=disable password=surl")
	if err != nil {
		log.Fatal(err)
	}
	defer srv.db.Close()

	api.RegisterSurlServer(s, srv)
	l, err := net.Listen("tcp", net.JoinHostPort("", srv.listenPort))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on", l.Addr())
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
