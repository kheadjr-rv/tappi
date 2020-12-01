build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./src
	go build -o tappi ./src

run: build
	./tappi local