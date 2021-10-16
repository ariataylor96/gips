package records

type Record struct {
	Offset  uint32
	Size    uint16
	RLESize uint16
	IsRLE   bool

	rawValue []byte
}

func (r *Record) Value() []byte {
	if r.IsRLE {
		return r.RLEValue()
	}

	return r.rawValue
}

func (r *Record) RLEValue() []byte {
	res := make([]byte, 0)

	for i := 0; i < int(r.RLESize); i++ {
		res = append(res, r.rawValue[0])
	}

	return res
}

func (r *Record) Apply(data *[]byte) {
	for idx, val := range r.Value() {
		(*data)[r.Offset+uint32(idx)] = val
	}
}
