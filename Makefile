build:
	go build ./...

test:
	go test ./...

run:
	go run main.go

.PHONY: deploy
deploy:
	sudo systemctl stop web-service
	GOOS=linux GOARCH=amd64 go build
	mkdir -p ~/prod
	mv web-service ~/prod/web-service
	cp ./deploy/web-service.service /etc/systemd/system/web-service.service
	sudo systemctl daemon-reload
	sudo systemctl start web-service
	sudo systemctl enable web-service



