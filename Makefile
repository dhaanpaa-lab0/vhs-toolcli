all: build
clean:
	rm -rfv ./bin
build: clean
	GOOS=linux GOARCH=amd64 go build -o bin/vhs-toolcli.linux-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/vhs-toolcli.osx-arm64 main.go

test: build
	scp bin/vhs-toolcli.linux-amd64 denadmin@ec2-43146.nexus-svcs.net:vhs-toolcli

deploy: build
	scp bin/vhs-toolcli.linux-amd64 denadmin@system-core:/opt/nexus/bin/vhs-toolcli