package domain

type Decoder interface {
	Decode() (interface{}, error)
}
