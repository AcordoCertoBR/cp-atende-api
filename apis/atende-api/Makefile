install-dependencies:
	npm ci
	npm install -g serverless@3.27

build:
	export GO111MODULE=on
	export CGO_ENABLED=1

	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/get-customer/bootstrap  get-customer/handler/main.go
	chmod +x bin/get-customer/bootstrap
	zip -j bin/get-customer.zip bin/get-customer/bootstrap

clean:
	rm -rf ./bin

deploy:
	serverless deploy --stage $(STAGE) --verbose