build:
	go1.8rc2 build -buildmode=plugin -o plugin.so plugin.go
	go1.8rc2 build main.go

clean:
	go clean
