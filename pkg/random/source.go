package random

import (
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

//func (d *Dictionary) SetSource(source rand.Source) *Dictionary {
//	d.source = source
//	return d
//}