DB_PORT := 3320

run:
	air -c .air.conf

lint:
	@#hash golangci-lint 2>/dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1
	golangci-lint run

build:
	go build -ldflags "-s -w" -o ./tmp/server ./cmd/main.go

gci:
	@GO111MODULE=off go get github.com/daixiang0/gci
	gci -w --local github.com/namnguyen/backend .

gen-mock:
	@mockery --inpackage --name=Repository --dir=./repository/option
	@mockery --inpackage --name=Repository --dir=./repository/poll
	@mockery --inpackage --name=Repository --dir=./repository/user
	@mockery --inpackage --name=Repository --dir=./repository/userpoll

unit-test:
	@mkdir coverage || true
	@go test -race -v -coverprofile=coverage/coverage.txt.tmp -count=1  ./...
	@cat coverage/coverage.txt.tmp | grep -v "mock_" > coverage/coverage.txt
	@go tool cover -func=coverage/coverage.txt
	@go tool cover -html=coverage/coverage.txt -o coverage/index.html
other_file:
	echo "This will run first"
	touch other_file
some_file: other_file
	echo "This will run second, because it depends on other_file"
	touch some_file
test: go/vendor
	@FAKE_DB_HOST=localhost:$(DB_PORT) FAKE_DB_USERNAME=root FAKE_DB_PASSWORD=root FAKE_DB_DATABASE=pang_test go test -p 1 -cover $(shell go list ./... |grep -v 'vendor') -coverprofile=coverage.out




