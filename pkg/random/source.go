package random

import (
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())