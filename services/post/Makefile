.PHONY: build docker clean ck

all: build docker clean

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o catdogs-post
	@echo "------ build go success ------"

docker:
	docker build -t catdogs-post .
	docker tag catdogs-post yokooll/catdogs-post
	docker push yokooll/catdogs-post
	@echo "------ docker push success ------"

clean:
	@rm -rf catdogs-post
	@go clean -i .
	@echo "------ clean done ------"
	@docker images

ck:
	@docker images|grep none|awk '{print $3}'|xargs docker rmi -f