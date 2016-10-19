package = github.com/fcingolani/chronobrake

.PHONY: release

release:
	go get
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/chronobrake-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o release/chronobrake-linux-386 $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/chronobrake-darwin-amd64 $(package)
	GOOS=darwin GOARCH=386 go build -o release/chronobrake-darwin-386 $(package)
	GOOS=windows GOARCH=amd64 go build -o release/chronobrake-windows-amd64.exe $(package)
	GOOS=windows GOARCH=386 go build -o release/chronobrake-windows-386.exe $(package)
