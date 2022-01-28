build:
	go build

buildlinux:
	GOOS=linux GOARCH=amd64 go build -o ./release/linux-amd64-ngcli

buildmac:
	GOOS=darwin GOARCH=amd64 go build -o ./release/darwin-amd64-ngcli

release: buildlinux buildmac
	cd ./release; \
	shasum -a 256 linux-amd64-ngcli darwin-amd64-ngcli > checksumfile; \
	shasum -a 256 -c checksumfile; \
	cat checksumfile; \
	cd -

clear:
	rm -rf *ngcli 

.PHONY: release
