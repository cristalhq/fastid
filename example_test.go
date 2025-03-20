package fastid_test

import (
	"fmt"

	"github.com/cristalhq/fastid"
)

func ExampleGenerator() {
	g, err := fastid.NewGenerator(fastid.DefaultEpoch, 11)
	if err != nil {
		panic(err)
	}

	id := g.Next()

	id2, err := fastid.Parse(id.String())
	if err != nil {
		panic(err)
	}

	if id != id2 {
		panic(fmt.Sprintf("not equal: %x vs %x", id, id2))
	}

	// Output:
}
