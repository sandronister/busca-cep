package di

import (
	"context"

	"github.com/sandronister/busca-cep/internals/pkg/http"
	"github.com/sandronister/busca-cep/internals/pkg/interfaces"
	"github.com/sandronister/busca-cep/internals/pkg/request"
	"github.com/sandronister/busca-cep/internals/usecase"
)

func NewVIACep(ctx context.Context, path string) interfaces.CepServices {
	request := request.New(ctx)
	httpService := http.New(request)
	return usecase.NewVIACep(path, httpService)
}
