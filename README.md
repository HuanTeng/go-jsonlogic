# go-jsonlogic

[![Travis CI](https://travis-ci.org/HuanTeng/go-jsonlogic.svg?branch=master)](https://travis-ci.org/HuanTeng/go-jsonlogic)
[![Go Report Card](https://goreportcard.com/badge/github.com/HuanTeng/go-jsonlogic)](https://goreportcard.com/report/github.com/HuanTeng/go-jsonlogic)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/3e9df51b227c47b6b903a2a78ae62072)](https://www.codacy.com/app/the729/go-jsonlogic?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=HuanTeng/go-jsonlogic&amp;utm_campaign=Badge_Grade)

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