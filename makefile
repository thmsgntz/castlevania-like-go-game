SOURCEDIR = castlevania-like
MAIN = main.go
SOURCES = $(wildcard $(SOURCEDIR)/*.go)

build:
	go build -o bin/main ${SOURCES}

run: build
	bin/main

gorun:
	go run ${SOURCES}
	
clean:
	rm -f bin/*