package processors

type Processor interface {
	Process(message interface{}) error
	GetID() string
}
