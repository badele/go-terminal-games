all: deps build

deps:
	go get -u github.com/JoelOtter/termloop

build:
	go build

test:
	go test -v

win:
    GOOS=windows GOARCH=386 go build -o test.exe test.go