package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sandronister/busca-cep/configs"
	"github.com/sandronister/busca-cep/internals/di"
	"github.com/sandronister/busca-cep/internals/usecase"
)

func main() {

	cdnChannel := make(chan []byte)
	viaChannel := make(chan []byte)

	config, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		w.Header().Set("Content-Type", "application/json")
		cdnService := di.NewCDNCep(ctx, config.CDNCepPath)
		viaService := di.NewVIACep(ctx, config.VIACepPath)

		go cdnService.Get(cdnChannel, "18050-605")
		go viaService.Get(viaChannel, "18050-605")

		select {
		case msg := <-cdnChannel:
			var cepCDN usecase.CdnDTO
			json.Unmarshal(msg, &cepCDN)
			json.NewEncoder(w).Encode(cepCDN)
		case msg := <-viaChannel:
			var cepVIA usecase.VIACepDTO
			json.Unmarshal(msg, &cepVIA)
			json.NewEncoder(w).Encode(cepVIA)
		case <-time.After(time.Second):
			w.Write([]byte("Timeout"))
		}
	})

	http.ListenAndServe(":8080", mux)

}
