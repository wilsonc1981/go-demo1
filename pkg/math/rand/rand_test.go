package rand_test

import (
	"fmt"
	"go-demo1/pkg/math/rand"
	"log"
)

func ExampleGetRandomIntBetween() {
	for i := 0; i < 20; i++ {
		myInt := rand.GetRandomIntBetween(10, 200)
		if myInt < 10 || myInt >= 200 {
			log.Fatalf("myInt的數值應該在[10, 200)之間: %d", myInt)
		}
	}
	fmt.Println(rand.GetRandomIntBetween(10, 11))
	// Output:
	// 10
}
