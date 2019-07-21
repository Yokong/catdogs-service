.PHONY: genpb

all: genpb

genpb:
	protoc -I=$$GOPATH/src/catdogs-proto --micro_out=./pb --go_out=./pb $$GOPATH/src/catdogs-proto/*proto