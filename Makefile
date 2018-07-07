PKG := github.com/dchirikov/uniqrode
all:
	go build uniqrode.go

clean:
	[[ -x ./uniqrode ]] && rm -f ./uniqrode

install:
	go install $(PKG)

test:
	go test -v -cover ./...
