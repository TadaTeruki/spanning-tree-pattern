all:

asset:
	(cd src; GOOS=js GOARCH=wasm go build -o ../main.wasm pattern.go main.go)
