package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "welcome my dear James" {
		t.Error("got ", s, " expected, welcome my dear James")
	}
}

func ExampleGreet(t *testing.T) {
	fmt.Println(Greet("James"))
	// Output:
	// welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
