package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LavaJover/storage-api-gateway/internal/config"
	storagepb "github.com/LavaJover/storage-master/storage-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateStorage(w http.ResponseWriter, r *http.Request){
	
}

func CreateCell(w http.ResponseWriter, r *http.Request){

}

func CreateBox(w http.ResponseWriter, r *http.Request){

}

func main(){

	cfg := config.MustLoad()
	fmt.Println(cfg)

	// Connect to storage-service
	storageServiceConn, err := grpc.Dial(cfg.GRPCStorageService.Host+":"+cfg.GRPCStorageService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("failed to connect storage service: %v", err)
	}

	defer storageServiceConn.Close()
	storageServiceClient := storagepb.NewStorageServiceClient(storageServiceConn)

	http.HandleFunc("/api/v1/storages", )
	http.HandleFunc("/api/v1/cells", nil)
	http.HandleFunc("/api/v1/boxes". )

}