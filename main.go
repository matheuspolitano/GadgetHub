package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matheuspolitano/GadgetHub/gapi"
	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"github.com/matheuspolitano/GadgetHub/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
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
	//runGatewayServer(context.Background(), store, conf)
}

func runGatewayServer(
	ctx context.Context,
	store db.Store,
	config utils.Config,

) {
	// create gapi server
	server, err := gapi.NewServer(store, config)
	if err != nil {
		log.Fatal(err)
	}
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	// runtime server mux
	gserverMux := runtime.NewServeMux(jsonOption)

	err = pb.RegisterGadgetHubHandlerServer(ctx, gserverMux, server)
	if err != nil {
		log.Fatal(err)
	}

	serverMux := http.NewServeMux()
	serverMux.Handle("/", gserverMux)

	listener, err := net.Listen("tcp", config.HTTPCServerAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening %s", config.HTTPCServerAddress)

	err = http.Serve(listener, serverMux)
	if err != nil {
		log.Fatal(err)
	}
}

func runGrpcServer(ctx context.Context, store db.Store, conf utils.Config) {
	server, err := gapi.NewServer(store, conf)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGadgetHubServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", conf.GRPCServerAddress)
	if err != nil {
		log.Fatalf("Cannot create listener %s", conf.GRPCServerAddress)
	}
	log.Printf("server started %s", conf.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start sRPC server")
	}

}
