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

	_ "github.com/LavaJover/storage-api-gateway/cmd/api-gateway/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	storageServiceClient storagepb.StorageServiceClient
)

// @title storage-master API
// @version 1.0
// @description API for storage-api-gateway
// @BasePath /api/v1

// Endpoint to create new storage
type createStorageOkResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// @Summary Create new storage
// @Description Create new named storage to store cells
// @Tags storages
// @Accept json
// @Produce json
// @Success 201 {object} createStorageOkResponse
// @Failure 400 {string} string "Bad request"
// @Failure 405 {string} string "Method is not supported"
// @Failure 500 {string} string "Storage service failed"
// @Router /api/v1/storages [post]
func CreateStorageHandler(w http.ResponseWriter, r *http.Request) {

	// Add basic headers to response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Ensure request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Creating request to storage-service
	var newStorage storagepb.CreateStorageRequest
	err := json.NewDecoder(r.Body).Decode(&newStorage)
	if err != nil {
		http.Error(w, "Failed to parse JSON: "+err.Error(), http.StatusBadRequest)
	}

	// Process response from storage-service
	response, err := storageServiceClient.CreateStorage(context.Background(), &newStorage)

	if err != nil {
		http.Error(w, "Error from storage service", http.StatusInternalServerError)
	}

	// Process HTTP response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

// Endpoint to create new cell
type createCellOkResponse struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	StorageID uint64 `json:"storage_id"`
}

// @Summary Create new cell
// @Description Create new named cell connected to storage to store boxes
// @Tags cells
// @Accept json
// @Produce json
// @Success 201 {object} createCellOkResponse
// @Failure 400 {string} string "Bad request"
// @Failure 405 {string} string "Method is not supported"
// @Failure 500 {string} string "Storage service failed"
// @Router /api/v1/cells [post]
func CreateCellHandler(w http.ResponseWriter, r *http.Request) {
	// Add basic headers to response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Ensure request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Creating request to storage-service
	var newCell storagepb.AddCellRequest
	err := json.NewDecoder(r.Body).Decode(&newCell)
	if err != nil {
		http.Error(w, "Failed to parse JSON: "+err.Error(), http.StatusBadRequest)
	}

	// Process response from storage-service
	response, err := storageServiceClient.AddCell(context.Background(), &newCell)

	if err != nil {
		http.Error(w, "Error from storage service", http.StatusInternalServerError)
	}

	// Process HTTP response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Endpoint to create new box
type createBoxOkResponse struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	CellID uint64 `json:"cell_id"`
}

// @Summary Create new box
// @Description Create new named box connected to cell
// @Tags boxes
// @Accept json
// @Produce json
// @Success 201 {object} createBoxOkResponse
// @Failure 400 {string} string "Bad request"
// @Failure 405 {string} string "Method is not supported"
// @Failure 500 {string} string "Storage service failed"
// @Router /api/v1/boxes [post]
func CreateBoxHandler(w http.ResponseWriter, r *http.Request) {
	// Add basic headers to response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Ensure request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Creating request to storage-service
	var newBox storagepb.AddBoxRequest
	err := json.NewDecoder(r.Body).Decode(&newBox)
	if err != nil {
		http.Error(w, "Failed to parse JSON: "+err.Error(), http.StatusBadRequest)
	}

	// Process response from storage-service
	response, err := storageServiceClient.AddBox(context.Background(), &newBox)

	if err != nil {
		http.Error(w, "Error from storage service", http.StatusInternalServerError)
	}

	// Process HTTP response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {

	cfg := config.MustLoad()
	fmt.Println(cfg)

	// Connect to storage-service
	storageServiceConn, err := grpc.Dial(":"+cfg.GRPCStorageService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect storage service: %v", err)
	}

	defer storageServiceConn.Close()
	storageServiceClient = storagepb.NewStorageServiceClient(storageServiceConn)

	http.HandleFunc("/api/v1/storages", CreateStorageHandler)
	http.HandleFunc("/api/v1/cells", CreateCellHandler)
	http.HandleFunc("/api/v1/boxes", CreateBoxHandler)
	http.HandleFunc("/api/v1/swagger/", httpSwagger.WrapHandler)

	if err := http.ListenAndServe(cfg.HTTPServer.Host+":"+cfg.HTTPServer.Port, nil); err != nil {
		log.Fatalf("failed to start HTTP server")
	}

	slog.Info("HTTP server successfully serving", "address", cfg.HTTPServer.Host+":"+cfg.HTTPServer.Port)
}
