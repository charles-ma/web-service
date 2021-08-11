build:
	go build ./...
	docker build \
         --build-arg USER_ID=$(id -u) \
         --build-arg GROUP_ID=$(id -g) \
         -t web-service .

test:
	go test ./...

run:
	go run main.go

.PHONY: deploy-box
deploy-box:
	sudo systemctl stop web-service
	GOOS=linux GOARCH=amd64 go build
	mkdir -p ~/prod
	mv web-service ~/prod/web-service
	cp ./deploy/web-service.service /etc/systemd/system/web-service.service
	sudo systemctl daemon-reload
	sudo systemctl start web-service
	sudo systemctl enable web-service

.PHONY: deploy-container
deploy-container:
	$(MAKE) build
	# docker container stop web-service
	docker run -it -d --rm -p 8010:8080  --name 'web-service' -v /root/go/src/web-service:/go/src/web-service web-service

.PHONY: deploy
deploy:
	@echo 'deploy container'
	$(MAKE) deploy-container
	@echo 'deploy box'
	$(MAKE) deploy-box
