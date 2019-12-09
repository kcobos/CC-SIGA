coverageGO:
	go test ./parkings/src/... -cover -coverprofile .coverage.out
	go tool cover -func=.coverage.out

testGO:
	go test ./parkings/src/...