package interfaces

type Record struct {
	Offset  uint32
	Size    uint16
	RLESize uint16
	IsRLE   bool

	BaseValue []byte
}

type recordMethods interface {
	Value() []byte
}
