install-api:
	cd ./go/api && go install

install-etl:
	cd ./go/etl && go install
	
test-go:
	go vet ./go/... || true
	go test -v -cover -p 1 ./go/...
