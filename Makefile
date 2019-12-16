coverageParkings:
	go test ./parkings/src/... -cover -coverprofile .coverage.out
	go tool cover -func=.coverage.out
testParkings:
	go test ./parkings/src/...

coveragePlaces:
	go test ./places/src/... -cover -coverprofile .coverage.out
	go tool cover -func=.coverage.out
testPlaces:
	go test ./places/src/...

test: testParkings testPlaces