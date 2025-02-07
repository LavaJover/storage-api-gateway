package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/LavaJover/storage-api-gateway/internal/config"
	// models "github.com/LavaJover/storage-master/storage-service/pkg/models"
	storagepb "github.com/LavaJover/storage-master/storage-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	storageServiceClient storagepb.StorageServiceClient
)

func CreateStorageHandler(w http.ResponseWriter, r *http.Request){

	// Add basic headers to response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions{
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Ensure request method is POST
	if r.Method != http.MethodPost{
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Creating request to storage-service
	var newStorage storagepb.CreateStorageRequest
	err := json.NewDecoder(r.Body).Decode(&newStorage)
	if err != nil{
		http.Error(w, "Failed to parse JSON: " + err.Error(), http.StatusBadRequest)
	}

	// Process response from storage-service
	response, err := storageServiceClient.CreateStorage(context.Background(), &newStorage)

	if err != nil{
		http.Error(w, "Error from storage service", http.StatusInternalServerError)
	}

	// Process HTTP response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func CreateCellHandler(w http.ResponseWriter, r *http.Request){

}

func CreateBoxHandler(w http.ResponseWriter, r *http.Request){

}

func main(){

	cfg := config.MustLoad()
	fmt.Println(cfg)

	// Connect to storage-service
	storageServiceConn, err := grpc.Dial(":"+cfg.GRPCStorageService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("failed to connect storage service: %v", err)
	}

	defer storageServiceConn.Close()
	storageServiceClient = storagepb.NewStorageServiceClient(storageServiceConn)

	http.HandleFunc("/api/v1/storages", CreateStorageHandler)
	http.HandleFunc("/api/v1/cells", CreateCellHandler)
	http.HandleFunc("/api/v1/boxes", CreateBoxHandler)

	if err := http.ListenAndServe(cfg.HTTPServer.Host+":"+cfg.HTTPServer.Port, nil); err != nil{
		log.Fatalf("failed to start HTTP server")
	}

	slog.Info("HTTP server successfully serving", "address", cfg.HTTPServer.Host+":"+cfg.HTTPServer.Port)
}