# This is how we want to name the binary output
BINARY=upm

# These are the values we want to pass for Version and BuildTime
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

.DEFAULT_GOAL: $(BINARY)

$(BINARY):
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=386 go build -o bin/${BINARY}-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/${BINARY}-windows-386 main.go