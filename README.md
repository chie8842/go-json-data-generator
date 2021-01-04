# go-json-data-generator

A tool to create json dummy data using [gofakeit](https://github.com/brianvoe/gofakeit)
## How To Use
1. Create schema file like [this](schemas/schema.json)
2. Execute the following command
```
go run cli/cli.go -s schema.json -d destination.json

```
