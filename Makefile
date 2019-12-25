parkingsCoverage:
	go test ./parkings/src/... -cover -coverprofile .coverageParkings.out
	go tool cover -func=.coverageParkings.out
parkingsTest:
	go test ./parkings/src/...

placesCoverage:
	go test ./places/src/... -cover -coverprofile .coveragePlaces.out
	go tool cover -func=.coveragePlaces.out
placesTest:
	go test ./places/src/...

goTest: parkingsTest placesTest
goCoverage: parkingsCoverage placesCoverage
	cat .coverageParkings.out > coverage.txt
	cat .coveragePlaces.out >> coverage.txt
	rm .coverageParkings.out
	rm .coveragePlaces.out

usersTest:
	pytest --doctest-modules ./users/users/test
usersCoverage:
	pytest --cov-report=xml --cov=users ./users/users/test
usersDependencies:
	pip install -r ./users/requirements.txt 

pythonTest: usersTest
pythonCoverage: usersCoverage