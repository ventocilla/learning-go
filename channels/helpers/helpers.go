package helpers

import (
	"math/rand"
	//"time"
)

func RandomNumber(n int) int {
	//rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value

	// seednum := rand.NewSource(time.Now().Unix())
	// randNum := rand.New(seednum)
	// return randNum.Int()

}
