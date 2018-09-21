# go-jsonlogic
Golang implementation of JsonLogic (jsonlogic.com)

## Example

```golang
var rule, data interface{}

json.Unmarshal([]byte(`
{"if": [
	{"<": [{"var":"temp"}, 0] }, "freezing",
	{"<": [{"var":"temp"}, 100] }, "liquid",
	"gas"
]}
`), &rule)

json.Unmarshal([]byte(`{"temp":55}`), &data)

got, err := jsonlogic.Apply(rule, data)
if err != nil {
	// handle error
}

fmt.Println(got)
// Output: liquid

```