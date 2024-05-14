package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/rs/zerolog/log"

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
		log.Fatal().Err(err)
	}

	conn, err := pgxpool.New(context.Background(), conf.DBSource)
	if err != nil {
		log.Fatal().Err(err)
	}

	store := db.NewStore(conn)
	//runGrpcServer(context.Background(), store, conf)
	runGatewayServer(context.Background(), store, conf)
}

func runGatewayServer(
	ctx context.Context,
	store db.Store,
	config utils.Config,

) {
	// create gapi server
	server, err := gapi.NewServer(store, config)
	if err != nil {
		log.Fatal().Err(err)
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
		log.Fatal().Err(err)
	}

	serverMux := http.NewServeMux()
	serverMux.Handle("/", gserverMux)

	httpServer := &http.Server{
		Handler: gapi.HttpLogger(serverMux),
		Addr:    config.HTTPCServerAddress,
	}
	log.Info().Msgf("start HTTP gateway server at %s", httpServer.Addr)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err)
	}
}

func runGrpcServer(ctx context.Context, store db.Store, conf utils.Config) {
	server, err := gapi.NewServer(store, conf)
	if err != nil {
		log.Fatal().Err(err)
	}
	grpcLog := grpc.UnaryInterceptor(gapi.GrpcLogger)

	grpcServer := grpc.NewServer(grpcLog)
	pb.RegisterGadgetHubServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", conf.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("Cannot create listener %s", conf.GRPCServerAddress))
	}
	log.Printf("server started %s", conf.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot start sRPC server"))
	}

}
