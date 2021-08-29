package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/firdavsich/surl/pkg/storage"

	_ "github.com/lib/pq"

	"github.com/firdavsich/surl/pkg/api"
	"google.golang.org/grpc"
)

type GPRCServer struct {
	api.SurlServer
	db *sql.DB
}

func (s *GPRCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	result, err := storage.Add(s.db, req.GetUrl())
	log.Println(result)
	if err != nil {
		log.Print(err)
		return &api.CreateResponse{Shorturl: ""}, err
	}
	return &api.CreateResponse{Shorturl: result}, nil
}

func (s *GPRCServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	result, err := storage.Get(s.db, req.GetShorturl())
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

	// open connect to postgresql
	srv.db, err = sql.Open("postgres", "postgres://surl:surl@localhost:5432/surl_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer srv.db.Close()

	api.RegisterSurlServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on", l.Addr())
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
