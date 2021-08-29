package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/firdavsich/surl/pkg/storage"

	_ "github.com/lib/pq"

	"github.com/firdavsich/surl/pkg/api"
	"google.golang.org/grpc"
)

type dbConfig struct {
	User     string
	Password string
	DB       string
	Host     string
	Port     string
	SSL      string
}

type GPRCServer struct {
	api.SurlServer
	db         *sql.DB
	listenPort string
	baseURL    string
}

func (s *GPRCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	log.Println(req.GetUrl())
	result, err := storage.Add(s.db, req.GetUrl())
	if err != nil {
		log.Print(err)
		return &api.CreateResponse{Shorturl: ""}, err
	}
	return &api.CreateResponse{Shorturl: s.baseURL + result}, nil
}

func (s *GPRCServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	log.Println(req.GetShorturl())
	hash := strings.TrimPrefix(req.GetShorturl(), s.baseURL)
	result, err := storage.Get(s.db, hash)
	if err != nil {
		log.Print(err)
		return &api.GetResponse{Url: ""}, err
	}

	return &api.GetResponse{Url: result}, nil
}

// get env or default value
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	dbConfig := dbConfig{
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DB:       getEnv("DB_NAME", "postgres"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		SSL:      getEnv("DB_SSLMODE", "disable"),
	}
	dbConection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DB,
		dbConfig.SSL,
		dbConfig.Password)

	var err error // ! FIXME: remove
	s := grpc.NewServer()
	srv := &GPRCServer{}
	srv.baseURL = getEnv("BASE_URL", "http://localhost:8080/")
	srv.listenPort = getEnv("PORT", "8080")

	// open connect to postgresql
	srv.db, err = sql.Open("postgres", dbConection)
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
