package main

import (
	"database/sql"
	"github.com/firdavsich/surl/pkg/storage"
	"log"
	"net"
	"context"

	_ "github.com/lib/pq"

	"github.com/firdavsich/surl/pkg/api"
	"google.golang.org/grpc"
)


type GPRCServer struct{
	api.SurlServer

}

var db, _= sql.Open("postgres", "postgres://surl:surl@localhost/surl_db?sslmode=disable")


func (s *GPRCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	result, err:=storage.Add(db, req.GetUrl())
	log.Println(result)
	if err != nil{
		log.Print(err)
		return &api.CreateResponse{Shorturl: ""}, err
	}
	return &api.CreateResponse{Shorturl: result}, nil
}

func (s *GPRCServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	result, err:=storage.Get(db, req.GetShorturl())
	log.Println(result)
	if err != nil{
		log.Print(err)
		return &api.GetResponse{Url: ""}, err
	}

	return &api.GetResponse{Url: result}, nil
}


func main() {
	// open connect to postgresql

	s := grpc.NewServer()
	srv := &GPRCServer{}
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
