.PHONY: build docker clean ck

all: build docker clean

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o catdogs-user
	@echo "------ build go success ------"

docker:
	docker build -t catdogs-user .
	docker tag catdogs-user yokooll/catdogs-user
	docker push yokooll/catdogs-user
	@echo "------ docker push success ------"

clean:
	@rm -rf catdogs-user
	@go clean -i .
	@echo "------ clean done ------"
	@docker images

ck:
	@docker images|grep none|awk '{print $3}'|xargs docker rmi -f