package fastid_test

import "github.com/cristalhq/fastid"

func ExampleGenerator() {
	g, err := fastid.NewGenerator(fastid.DefaultEpoch, 11)
	if err != nil {
		panic(err)
	}

	id := g.Next()
	println(id)

	// Output:
}
