package main

import (
	"fmt"
	"github.com/wiryawan46/go-grpc-auth-svc/pkg/config"
	"github.com/wiryawan46/go-grpc-auth-svc/pkg/db"
	"github.com/wiryawan46/go-grpc-auth-svc/pkg/pb"
	"github.com/wiryawan46/go-grpc-auth-svc/pkg/services"
	"github.com/wiryawan46/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 3,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcsServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcsServer, &s)

	if err := grpcsServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
