package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/LavaJover/storage-api-gateway/internal/config"
	"github.com/LavaJover/storage-api-gateway/pkg/middleware"

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

// @title Storage-Master API
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
// @Router /storages [post]
func CreateStorageHandler(w http.ResponseWriter, r *http.Request) {

	// Creating request to storage-service
	var newStorage storagepb.CreateStorageRequest
	err := json.NewDecoder(r.Body).Decode(&newStorage)
	if err != nil {
		http.Error(w, "Failed to parse JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Process response from storage-service
	response, err := storageServiceClient.CreateStorage(context.Background(), &newStorage)

	if err != nil {
		http.Error(w, "Error from storage service", http.StatusInternalServerError)
		return
	}

	// Process HTTP response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

// Endpoint to get storages by user_id
type getStoragesOkResponse struct{
	Storages []struct{
		ID uint64		`json:"id"` 
		Name string		`json:"name"`
		UserID uint64	`json:"user_id"`
	}
}

// @Summary Get storages by user_id
// @Description Get all storage instances related to the given user_id
// @Tags storages
// @Accept json
// @Produce json
// @Success 201 {object} getStoragesOkResponse
// @Failure 400 {string} string "Bad request"
// @Failure 403 {string} string "You dont't have enough permissions"
// @Failure 405 {string} string "Method is not supported"
// @Failure 500 {string} string "Storage service failed"
// @Router /storages [get]
func GetStoragesHandler(w http.ResponseWriter, r *http.Request){

	// Extract user_id from JWT
	// ...
	// Check if user have enough permissions to request this
	userID := 2

	// Create request to storage service
	getStoragesRequest := storagepb.GetStoragesRequest{
		UserId: uint64(userID),
	}

	// Process response from storage-service
	response, err := storageServiceClient.GetStorages(context.Background(), &getStoragesRequest)

	if err != nil{
		http.Error(w, "Error from storage service", http.StatusInternalServerError)
		return
	}

	// Process HTTP response
	w.WriteHeader(http.StatusOK)
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
// @Router /cells [post]
func CreateCellHandler(w http.ResponseWriter, r *http.Request) {

	// Creating request to storage-service
	var newCell storagepb.AddCellRequest
	err := json.NewDecoder(r.Body).Decode(&newCell)
	if err != nil {
		http.Error(w, "Failed to parse JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Process response from storage-service
	response, err := storageServiceClient.AddCell(context.Background(), &newCell)

	if err != nil {
		http.Error(w, "Error from storage service", http.StatusInternalServerError)
		return
	}

	// Process HTTP response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Endpoint to get cells by storage_id
type getCellsOkResponse struct{
	Cells []struct{
		ID uint64 `json:"id"`
		Name string `json:"name"`
		StorageID uint64 `json:"storage_id"`
	}
}

// @Summary Get cells by storage_id
// @Description Get all cells by given storage_id with permission checking
// @Tags cells
// @Accept json
// @Produce json
// @Param storage_id query uint true "Storage ID"
// @Success 201 {object} getCellsOkResponse
// @Failure 400 {string} string "Bad request"
// @Failure 403 {string} string "You don't have enough permissions"
// @Failure 405 {string} string "Method is not supported"
// @Failure 500 {string} string "Storage service failed"
// @Router /cells [get]
func GetCellsHandler(w http.ResponseWriter, r *http.Request){

	// Extract JWT and user_id
	// ....
	// Check if user has enough permissions to do this request

	// Extract storage_id from query
	storageIDStr := r.URL.Query().Get("storage_id")
	if storageIDStr == ""{
		http.Error(w, "Query param storage_id was not found", http.StatusBadRequest)
		return
	}

	storageID, err := strconv.ParseUint(storageIDStr, 10, 64)
	if err != nil{
		http.Error(w, "Incorrect param storage_id", http.StatusBadRequest)
		return
	}

	// Creating request to storage service
	getCellsRequest := storagepb.GetCellsRequest{
		StorageId: storageID,
	}

	response, err := storageServiceClient.GetCells(context.Background(), &getCellsRequest)
	if err != nil{
		http.Error(w, "Error from storage service", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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
// @Router /boxes [post]
func CreateBoxHandler(w http.ResponseWriter, r *http.Request) {

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

	// Handle endpoints
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/storages", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodPost:
			CreateStorageHandler(w, r)
		case http.MethodGet:
			GetStoragesHandler(w, r)
		default:
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/cells", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodPost:
			CreateCellHandler(w, r)
		case http.MethodGet:
			GetCellsHandler(w, r)
		default:
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/boxes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodPost:
			CreateBoxHandler(w, r)
		default:
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/swagger/", httpSwagger.WrapHandler)

	// Add middleware
	finalHandler := middleware.ChainMiddleware(mux, middleware.CorsMiddleware, middleware.RateLimitMiddleware, middleware.LoggingMiddleware)

	if err := http.ListenAndServe(cfg.HTTPServer.Host+":"+cfg.HTTPServer.Port, finalHandler); err != nil {
		log.Fatalf("failed to start HTTP server")
	}

	slog.Info("HTTP server successfully serving", "address", cfg.HTTPServer.Host+":"+cfg.HTTPServer.Port)
}