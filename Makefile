include ./.env
export

start:
	go run ./cmd/digital-library/main.go

local:
	go run ./cmd/digital-library/main.go -config ./config/local.yaml
