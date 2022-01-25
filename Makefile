build:
	go build
buildlinux:
	GOOS=linux GOARCH=amd64 go build -o linux-amd64-ngcli
buildmac:
	GOOS=darwin GOARCH=amd64 go build -o darwin-amd64-ngcli
clear:
	rm -rf *ngcli 
