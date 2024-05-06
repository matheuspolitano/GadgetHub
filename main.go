package main

import (
	"context"
	"log"
	"net"

	"fmt"

	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matheuspolitano/GadgetHub/gapi"
	db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"
	"github.com/matheuspolitano/GadgetHub/pkg/pb"
	"github.com/matheuspolitano/GadgetHub/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Print the webhook payload
	fmt.Println("Webhook Payload:", string(body))

	// Respond with success status
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Register webhook handler function
	http.HandleFunc("/webhook", webhookHandler)

	// Start the server
	fmt.Println("Server listening on port 9000...")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal("Server error:", err)
	}
}

func main1() {
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
