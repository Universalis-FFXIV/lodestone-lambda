# lodestone-character

[aws-lambda-go](https://github.com/aws/aws-lambda-go)

```pwsh
$env:GOOS = "linux"
$env:GOARCH = "arm64"
$env:CGO_ENABLED = "0"
go build -o bootstrap main.go
build-lambda-zip -o lambda-handler.zip bootstrap
```
