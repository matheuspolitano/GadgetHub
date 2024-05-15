package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matheuspolitano/GadgetHub/gapi"
	"github.com/matheuspolitano/GadgetHub/listener"
	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"github.com/matheuspolitano/GadgetHub/utils"
	"github.com/rs/zerolog/log"
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
	runDBMigration(conf.MigrationURL, conf.DBSource)

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
	webhookHandler := listener.NewWebhookHandler(config)

	err = pb.RegisterGadgetHubHandlerServer(ctx, gserverMux, server)
	if err != nil {
		log.Fatal().Err(err)
	}

	serverMux := mux.NewRouter()

	serverMux.PathPrefix("/api/").Handler(gserverMux)
	serverMux.HandleFunc("/webhook", webhookHandler.VerifyWebhook).Methods("GET")
	serverMux.HandleFunc("/webhook", webhookHandler.HandleWebhook).Methods("POST")
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
func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}
