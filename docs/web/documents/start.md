# Getting started

Getting started with TAPPI.

```go
package main

func main() {
    //...
}
```

```bash
# Build the wasm program that contains the user interface.
GOARCH=wasm GOOS=js go build -o web/app.wasm
```