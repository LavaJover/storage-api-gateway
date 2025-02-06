package main

import (
	"fmt"

	"github.com/LavaJover/storage-api-gateway/internal/config"
)

func main(){

	cfg := config.MustLoad()
	fmt.Println(cfg)
}