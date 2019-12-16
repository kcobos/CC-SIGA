coverageParkings:
	go test ./parkings/src/... -cover -coverprofile .coverageParkings.out
	go tool cover -func=.coverageParkings.out
testParkings:
	go test ./parkings/src/...

coveragePlaces:
	go test ./places/src/... -cover -coverprofile .coveragePlaces.out
	go tool cover -func=.coveragePlaces.out
testPlaces:
	go test ./places/src/...

testGo: testParkings testPlaces
coverageGo: coverageParkings coveragePlaces
	cat .coverageParkings.out > coverage.txt
	cat .coveragePlaces.out >> coverage.txt
	rm .coverageParkings.out
	rm .coveragePlaces.out