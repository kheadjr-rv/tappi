build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./src
	go build -o tappi ./src

static:
	GOARCH=wasm GOOS=js go build -o docs/web/app.wasm ./src
	go build -o tappi ./src
	./tappi github -o docs
	rm ./tappi

run: build
	./tappi local