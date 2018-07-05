PKG := github.com/dchirikov/uniqrode
all:
	mkdir -p bin
	go build -o bin/uniqrode main.go

clean:
	rm -rf bin

install:
	go install $(PKG)

test:
	go test -v -cover ./...
