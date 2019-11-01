.PHONY: default dep build kill run

default: build

dep:
	@go mod download

build: dep
	@go build

kill:
	@sudo pkill -9 gst || true

run: build kill
	@sudo $(PWD)/gst -f $(PWD)/gst.toml