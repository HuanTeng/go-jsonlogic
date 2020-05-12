package jsonlogic_test

import (
	"encoding/json"
	"fmt"

	"github.com/ShankarParimi/go-jsonlogic"
)

func ExampleApply() {
	var rule, data interface{}

	json.Unmarshal([]byte(`
{"if" : [
	{"<": [{"var":"temp"}, 0] }, "freezing",
	{"<": [{"var":"temp"}, 100] }, "liquid",
	"gas"
]}`), &rule)

	json.Unmarshal([]byte(`{"temp":55}`), &data)

	got, err := jsonlogic.Apply(rule, data)
	if err != nil {
		// handle error
	}

	fmt.Println(got)
	// Output: liquid
}
