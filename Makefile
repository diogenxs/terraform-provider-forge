default: build

.PHONY=build
build:
	go build -o terraform-provider-forge

test: build
	terraform init
	terraform plan -out plan
