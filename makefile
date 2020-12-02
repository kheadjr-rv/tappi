build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./src
	go build -o tappi ./src

static: build
	./tappi github -o docs
	cp -r ./web ./docs
	rm ./tappi

run: build
	./tappi local