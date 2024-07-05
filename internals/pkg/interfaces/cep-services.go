package interfaces

type CepServices interface {
	Get(data chan<- []byte, cep string)
}
