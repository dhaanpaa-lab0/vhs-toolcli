all: build
clean:
	rm -rfv ./bin
build: clean
	GOOS=linux GOARCH=amd64 go build -o bin/vhs-toolcli.linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/vhs-toolcli.linux-arm64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/vhs-toolcli.osx-arm64 main.go

deploy: build
	scp bin/vhs-toolcli.linux-amd64 denadmin@system-core:/opt/nexus/bin/vhs-toolcli.x86_64
	scp bin/vhs-toolcli.linux-arm64 denadmin@system-core:/opt/nexus/bin/vhs-toolcli.aarch64