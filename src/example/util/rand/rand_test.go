package rand

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	i := rand.Int()
	fmt.Println("rand int ", i)

	for i := 0; i < 10; i++ {
		intn := rand.Intn(10)
		fmt.Println("rand intn ", intn)
	}

	source := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(source)
	fmt.Println(r1.Intn(10))
}
