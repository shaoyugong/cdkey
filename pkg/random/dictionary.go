package random

type Dictionary struct {
	bytes string
	index Index
}

type Index struct {
	bits int
	mask int64
	max  int
}

func (d *Dictionary) SetBytes(str string) *Dictionary {
	d.bytes = str
	return d
}

func (d *Dictionary) SetIndex(bits int) *Dictionary {
	d.index.bits = bits 		 // n bits to represent a letter index
	d.index.mask = 1 << bits - 1 // All 1-bits, as many as letterIdxBits
	d.index.max = 63 / bits		 // n of letter indices fitting in 63 bits
	return d
}

// create a dictionary by custom string
func New(bytes string, bits int) *Dictionary {
	dict := &Dictionary{}
	dict.SetBytes(bytes)
	dict.SetIndex(bits)
	return dict
}