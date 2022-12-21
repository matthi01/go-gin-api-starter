package util

import (
	"math/rand"
	"time"
)

// TODO: set this up with UUIDs
func GenerateId() int {
	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 99999
	return rand.Intn(max-min+1) + min
}
