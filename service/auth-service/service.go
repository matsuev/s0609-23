package main

import (
	"context"
	"encoding/json"
	"log"
	"s0609-23/internal/demodb"
	"s0609-23/internal/proxyproto"
	"strconv"

	"github.com/jackc/pgx/v5"
)

// AuthService ...
type AuthService struct {
	proxyproto.UnimplementedCentrifugoProxyServer
	conn  *pgx.Conn
	query demodb.Querier
}

const connStr = "postgres://admin:adminpass@127.0.0.1:5432/demodb"

// NewAuthService ...
func NewAuthService() *AuthService {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalln(err)
	}

	query := demodb.New(conn)

	return &AuthService{
		conn:  conn,
		query: query,
	}
}

// Connect ...
func (s *AuthService) Connect(ctx context.Context, req *proxyproto.ConnectRequest) (*proxyproto.ConnectResponse, error) {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	authRequest := &Request{}

	if err := json.Unmarshal(req.Data, authRequest); err != nil {
		return &proxyproto.ConnectResponse{
			Error: &proxyproto.Error{
				Code:    107,
				Message: "bad request",
			},
		}, nil
	}

	acc, err := s.query.UserLogin(context.Background(), demodb.UserLoginParams{
		Username: authRequest.Username,
		Password: authRequest.Password,
	})

	if err != nil {
		return &proxyproto.ConnectResponse{
			Error: &proxyproto.Error{
				Code:    101,
				Message: "unauthrized",
			},
		}, nil
	}

	return &proxyproto.ConnectResponse{
		Result: &proxyproto.ConnectResult{
			User: strconv.Itoa(int(acc.ID)),
		},
	}, nil

}
