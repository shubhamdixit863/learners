package internal

type Cache interface {
	SetData(key interface{}, value interface{})
	GetData(key interface{}) (interface{}, error)
}
