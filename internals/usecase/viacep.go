package usecase

import (
	"fmt"
	"log"

	"github.com/sandronister/busca-cep/internals/pkg/interfaces"
)

type VIACepDTO struct {
	Bairro      string `json:"bairro"`
	Cep         string `json:"cep"`
	Complemento string `json:"complemento"`
	Ddd         string `json:"ddd"`
	Gia         string `json:"gia"`
	Ibge        string `json:"ibge"`
	Localidade  string `json:"localidade"`
	Logradouro  string `json:"logradouro"`
	Siafi       string `json:"siafi"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
}

type VIACep struct {
	path string
	http interfaces.HttpService
}

func NewVIACep(path string, http interfaces.HttpService) *VIACep {
	return &VIACep{path: path, http: http}
}

func (c *VIACep) Get(data chan<- []byte, cep string) {
	url := fmt.Sprintf(c.path, cep)
	res, err := c.http.Do(url)
	if err != nil {
		log.Fatal(err)
	}

	data <- res
}
