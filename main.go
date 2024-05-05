package main

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matheuspolitano/GadgetHub/gapi"
	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"github.com/matheuspolitano/GadgetHub/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	conf, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgxpool.New(context.Background(), conf.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	runGrpcServer(context.Background(), store, conf)

}

func runGrpcServer(ctx context.Context, store db.Store, conf utils.Config) {
	grpcServer := grpc.NewServer()
	server, err := gapi.NewServer(store, conf)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterGadgetHubServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", conf.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener")
	}
	log.Printf("server started")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start sRPC server")
	}

}
