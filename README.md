# go-json-data-generator

A tool to create json dummy data using [gofakeit](https://github.com/brianvoe/gofakeit)
## How To Use
1. Create schema file like [this](schemas/schema.json)
2. Execute the following command
```
go run cli/cli.go -s schema.json -d destination.json

```

## schema file example
There are two patterns of schema file for `go-json-data-generator`.
The first is a simple json file as below.

```schema.json
{
    "Id": "{{UUID()}}",
    "Name": "A green door",
    "Age": "{{Number(1,50)}}",
    "Price": 12.50
}
```
When using this style, `go-json-data-generator` creates one json record.

```destination.json
{"Age":10,"Id":"52fdfc07-2182-454f-963f-5f0f9a621d72","Name":"A green door","Price":12.5}

```


And the another pattern is to write the number of records at the top of list as below.

```schema.json
[
  "repeat(2)",
  {
      "Id": "{{UUID()}}",
      "Name": "A green door",
      "Age": "{{Number(1,50)}}",
      "Price": 12.50,
      "tags": ["home", "green"]
  }
]

```destination.json
{"Age":10,"Id":"52fdfc07-2182-454f-963f-5f0f9a621d72","Name":"A green door","Price":12.5,"tags":["home","green"]}
{"Age":26,"Id":"9566c74d-10d1-42c6-8981-855ad8681d0d","Name":"A green door","Price":12.5,"tags":["home","green"]}
```

When using this, `go-json-data-generator` creates a new line delimited json with the number specified by `repeat(number)`. 


Please see [gofakeit document](https://godoc.org/github.com/brianvoe/gofakeit#pkg-index) to confirm the functions which is able to use in the schema
