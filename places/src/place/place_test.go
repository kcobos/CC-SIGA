package place

import (
	"fmt"
	"testing"
)

func TestValidateInit(t *testing.T) {
	var placs Places
	_, err := placs.Len()
	if err == nil {
		t.Errorf("map is initialized")
	}
	if err.Error() != "Places not initialized" {
		t.Errorf("error message not valid")
	}
	placs.Init()
	if num, err := placs.Len(); err != nil || num != 0 {
		t.Errorf("map is not initialized")
	}
}
func TestValidatePlace(t *testing.T) {
	var p Place
	lat, long := p.Coor()
	if p.ID() != 0 || p.Address() != "" || lat != 0 || long != 0 ||
		p.NumParkings() != 0 || p.FreeParkings() != 0 {
		t.Errorf("place not well created")
	}
	p = Place{id: 1, lat: 37.19742395414327, long: -3.624779980964916,
		address:     "Calle Periodista Daniel Saucedo Aranda, 18014 Granada",
		numParkings: 2, freeParkings: 1}
	lat, long = p.Coor()
	if p.ID() != 1 || lat != 37.19742395414327 || long != -3.624779980964916 ||
		p.Address() != "Calle Periodista Daniel Saucedo Aranda, 18014 Granada" ||
		p.NumParkings() != 2 || p.FreeParkings() != 1 {
		t.Errorf("place not well created")
	}
}

var placs Places

func setup() {
	placs.Init()
}
func setupMock() {
	mockParkingList.lastID = 0
	mockParkingList.list = make(map[int]int)
}
func tearDown() {
	placs.lastID = 0
	placs.placeList = nil
	mockParkingList.lastID = 0
	mockParkingList.list = nil
}

type mockParking struct {
	lastID int
	list   map[int]int //[ParkingID]PlaceID
}

var mockParkingList mockParking

func mockCallParkingAddAPI(idPlace int) (int, int) {
	if mockParkingList.list == nil {
		return -1, 500
	}
	mockParkingList.lastID++
	mockParkingList.list[mockParkingList.lastID] = idPlace
	return mockParkingList.lastID, 200
}
func mockCallParkingDeleteAPI(idParking int) int {
	if mockParkingList.list == nil {
		return 500
	}
	if _, ok := mockParkingList.list[idParking]; !ok {
		return 404
	}
	delete(mockParkingList.list, idParking)
	return 200
}

func TestValidateAdd(t *testing.T) {
	setup()
	defer tearDown()
	p, err := placs.Add(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", mockCallParkingAddAPI)
	if err == nil {
		t.Errorf("fail to create a place (parking API)")
	}
	if err.Error() != "Error 500 on Parking while calling AddParking" {
		t.Errorf("error message not valid")
	}
	setupMock()
	p, err = placs.Add(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", mockCallParkingAddAPI)
	if err != nil {
		t.Errorf("fail to create a place")
	}
	if num, _ := placs.Len(); num != 1 {
		t.Errorf("map size is not 1")
	}
	lat, long := p.Coor()
	if p.ID() != 1 || lat != 37.19742395414327 || long != -3.624779980964916 ||
		p.Address() != "Calle Periodista Daniel Saucedo Aranda, 18014 Granada" ||
		p.NumParkings() != 1 || p.FreeParkings() != 0 {
		t.Errorf("place not well created or saved")
	}

	p, err = placs.Add(37.19742395414327, -3.624779980964916,
		"Calle Periodista Eugenio Selles, 18014 Granada", mockCallParkingAddAPI)
	if err == nil {
		t.Errorf("error creating duplicate coordinates place")
	}
	if err.Error() != fmt.Sprintf("Place in %f,%f already exists", 37.19742395414327, -3.624779980964916) {
		t.Errorf("error message not valid")
	}
}

func TestValidateExists(t *testing.T) {
	setup()
	setupMock()
	defer tearDown()
	placs.Add(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", mockCallParkingAddAPI)
	placs.Add(37.195458, -3.626633,
		"Calle Periodista Eugenio Selles, 18014 Granada", mockCallParkingAddAPI)
	if !placs.Exists(1) {
		t.Errorf("place 1 not exists")
	}
	if !placs.Exists(2) {
		t.Errorf("place 2 not exists")
	}
	if placs.Exists(5) {
		t.Errorf("place 5 exists")
	}
	if num, _ := placs.Len(); num != 2 {
		t.Errorf("map size is not 2")
	}
	if len(mockParkingList.list) != 2 {
		t.Errorf("parkingList is not 2")
	}
}

func TestValidateAddParking(t *testing.T) {
	setup()
	setupMock()
	defer tearDown()
	place, _ := placs.Add(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", mockCallParkingAddAPI)
	place.AddParking(mockCallParkingAddAPI)
	if place.NumParkings() != 2 {
		t.Errorf("no parking added")
	}
	if len(mockParkingList.list) != 2 {
		t.Errorf("map size is not 2")
	}
	place.AddParking(mockCallParkingAddAPI)
	if place.NumParkings() != 3 {
		t.Errorf("no parking added")
	}
	if len(mockParkingList.list) != 3 {
		t.Errorf("map size is not 3")
	}
	if num, _ := placs.Len(); num != 1 {
		t.Errorf("map size is not 1")
	}
}

func TestValidateDeleteParking(t *testing.T) {
	setup()
	setupMock()
	defer tearDown()
	place, _ := placs.Add(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", mockCallParkingAddAPI)
	place.AddParking(mockCallParkingAddAPI)
	place.AddParking(mockCallParkingAddAPI)
	if place.NumParkings() != 3 {
		t.Errorf("no parking added")
	}

	err := place.DeleteParking(1, mockCallParkingDeleteAPI)
	if err != nil {
		t.Errorf("error on deleting parking")
	}
	if place.NumParkings() != 2 {
		t.Errorf("no parking deleted")
	}
	if len(mockParkingList.list) != 2 {
		t.Errorf("map size is not 2")
	}

	err = place.DeleteParking(1, mockCallParkingDeleteAPI)
	if err == nil {
		t.Errorf("error on deleting deleted parking")
	}
	if err.Error() != "Place 1 does not have 1 parking" {
		t.Errorf("error message not valid")
	}
	if place.NumParkings() != 2 {
		t.Errorf("no parking deleted")
	}
	if len(mockParkingList.list) != 2 {
		t.Errorf("map size is not 2")
	}

	err = place.DeleteParking(2, mockCallParkingDeleteAPI)
	if err != nil {
		t.Errorf("error on deleting parking")
	}
	if place.NumParkings() != 1 {
		t.Errorf("no parking deleted")
	}
	if len(mockParkingList.list) != 1 {
		t.Errorf("map size is not 1")
	}

	err = place.DeleteParking(3, mockCallParkingDeleteAPI)
	if err == nil {
		t.Errorf("error on deleting last parking")
	}
	if err.Error() != "Place 1 has only one parking and it cannot be deleted \nTry to delete the place" {
		t.Errorf("error message not valid")
	}
}

func TestValidateGet(t *testing.T) {
	setup()
	defer tearDown()
	_, err := placs.Add(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", mockCallParkingAddAPI)
	if err == nil {
		t.Errorf("fail to create a place (parking API)")
	}
	_, err1 := placs.Get(1)
	if err1 == nil {
		t.Errorf("map contains 1")
	}
	setupMock()
	p, _ := placs.Add(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", mockCallParkingAddAPI)
	pFL, err2 := placs.Get(1)
	if err2 != nil {
		t.Errorf("map not contains 1")
	}
	lat, long := p.Coor()
	latFL, longFL := pFL.Coor()
	if p.ID() != pFL.ID() || p.Address() != pFL.Address() || lat != latFL || long != longFL ||
		p.NumParkings() != pFL.NumParkings() || p.FreeParkings() != pFL.FreeParkings() {
		t.Errorf("places not equal")
	}

	notP, err3 := placs.Get(12)
	latNotP, longNotP := notP.Coor()
	if err3 == nil {
		t.Errorf("map contains 12")
	}
	if err3.Error() != "Place 12 does not exist" {
		t.Errorf("error message not valid")
	}
	pNil := Place{}
	latNil, longNil := pNil.Coor()
	if notP.ID() != pNil.ID() || latNotP != latNil || longNotP != longNil ||
		notP.Address() != pNil.Address() ||
		notP.NumParkings() != pNil.NumParkings() || notP.FreeParkings() != pNil.NumParkings() {
		t.Errorf("place not well created")
	}
}

func TestValidateDelete(t *testing.T) {
	setup()
	setupMock()
	defer tearDown()
	place, _ := placs.Add(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", mockCallParkingAddAPI)
	place.AddParking(mockCallParkingAddAPI)
	place.AddParking(mockCallParkingAddAPI)
	if place.NumParkings() != 3 {
		t.Errorf("no parking added")
	}
	if len(mockParkingList.list) != 3 {
		t.Errorf("map size is not 3")
	}
	if num, _ := placs.Len(); num != 1 {
		t.Errorf("map size is not 1")
	}
	err := placs.Delete(1, mockCallParkingDeleteAPI)
	if err != nil {
		t.Errorf("error deleting place 1")
	}

	err = placs.Delete(1, mockCallParkingDeleteAPI)
	if err == nil {
		t.Errorf("error deleting place 1 not exists")
	}
	if err.Error() != "Place 1 does not exist" {
		t.Errorf("error message not valid")
	}
}

// func TestValidate_(t *testing.T) {
// 	setup()
// }
