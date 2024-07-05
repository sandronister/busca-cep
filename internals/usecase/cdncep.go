package usecase

import (
	"fmt"
	"log"

	"github.com/sandronister/busca-cep/internals/pkg/interfaces"
)

type CdnDTO struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	Code       string `json:"code"`
	District   string `json:"district"`
	Ok         bool   `json:"ok"`
	State      string `json:"state"`
	Status     int64  `json:"status"`
	StatusText string `json:"statusText"`
}

type CDNCep struct {
	path string
	http interfaces.HttpService
}

func NewCDNCep(path string, http interfaces.HttpService) *CDNCep {
	return &CDNCep{path: path, http: http}
}

func (c *CDNCep) Get(data chan<- []byte, cep string) {
	url := fmt.Sprintf(c.path, cep)
	res, err := c.http.Do(url)
	if err != nil {
		log.Fatal(err)
	}
	data <- res
}
