package main

import (
	"math/rand"
	"time"
)

type password struct{
	length int
	types []charsets
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func (p password) generate() string {
	indexTypes := len(p.types)
	result := ""
	for i := 0; i <= p.length; i++ {

		number := r.Intn(indexTypes)
		typ := p.types[number]
		result += typ.getRandomChar()

	}
	return result

}