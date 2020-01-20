
all: linux osx windows pi pizero

linux:
	GOOS=linux GOARCH=amd64 go build -o build/freshpow-linux-amd64

osx:
	GOOS=darwin GOARCH=amd64 go build -o build/freshpow-darwin-amd64

windows:
	GOOS=windows GOARCH=amd64 go build -o build/freshpow-windows-amd64

pi:
	GOOS=linux GOARCH=arm GOARM=7 go build -o build/freshpow-linux-arm7

pizero:
	GOOS=linux GOARCH=arm GOARM=6 go build -o build/freshpow-linux-arm6
