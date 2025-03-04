package services

type Service interface {
	GetData() string
	SearchData() string
	AddData(data string) error
}
