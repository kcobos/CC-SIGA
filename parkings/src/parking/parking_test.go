package parking

import (
	"testing"
)

func TestValidateInit(t *testing.T) {
	var parks Parkings
	_, err := parks.Len()
	if err == nil {
		t.Errorf("map is initialized")
	}
	if err.Error() != "Parkings not initialized" {
		t.Errorf("error message not valid")
	}
	parks.Init()
	if num, err := parks.Len(); err != nil || num != 0 {
		t.Errorf("map is not initialized")
	}
}

var parks Parkings
var mockPlaceList map[int]mockPlace

func setup() {
	parks.Init()
}
func setupMock() {
	mockPlaceList = make(map[int]mockPlace)
}
func tearDown()  {
	parks.parkingList = nil
	parks.lastID = 0
	mockPlaceList = nil
}

type mockPlace struct {
	freeParkings int
}

func mockCallPlaceAPI(idPlace int, i int) int {
	mp, ok := mockPlaceList[idPlace]
	ret := 404
	if ok {
		switch i {
		case -1:
			mp.freeParkings--
			mockPlaceList[idPlace] = mp
			ret = 200
		case 1:
			mp.freeParkings++
			mockPlaceList[idPlace] = mp
			ret = 200
		default:
			ret = 400
		}
	}
	return ret
}
func TestValidateAdd(t *testing.T) {
	setup()
	defer tearDown()
	p := parks.Add(21)
	if num, _ := parks.Len(); num != 1 {
		t.Errorf("map size is not 1")
	}
	if p.ID() != 1 || p.PlaceID() != 21 || p.Status() != -1 {
		t.Errorf("parking not well created")
	}
	pFL, err := parks.Get(1)
	if err != nil {
		t.Errorf("map not contains 1")
	}
	if p != pFL {
		t.Errorf("parkings not equal")
	}

	p1 := parks.Add(23)
	if num, _ := parks.Len(); num != 2 {
		t.Errorf("map size is not 2")
	}
	if p1.ID() != 2 || p1.PlaceID() != 23 || p1.Status() != -1 {
		t.Errorf("parking not well created")
	}
	p1FL, err := parks.Get(2)
	if err != nil {
		t.Errorf("map not contains 2")
	}
	if p1 != p1FL {
		t.Errorf("parkings not equal")
	}
}
func TestValidateExists(t *testing.T) {
	setup()
	defer tearDown()
	parks.Add(21)
	parks.Add(21)
	if !parks.Exists(1) {
		t.Errorf("parking 1 not exists")
	}
	if !parks.Exists(2) {
		t.Errorf("parking 2 not exists")
	}
	if parks.Exists(5) {
		t.Errorf("parking 5 exists")
	}
}

func TestValidateGet(t *testing.T) {
	setup()
	defer tearDown()
	p, err := parks.Get(1)
	if p.ID() != 0 || p.Status() != 0 || p.PlaceID() != 0 {
		t.Errorf("parking not default")
	}
	if err == nil {
		t.Errorf("no error on getting 1")
	}
	if err.Error() != "Parking 1 does not exist" {
		t.Errorf("error message not valid")
	}
	parks.Add(21)
	p, err = parks.Get(1)
	if p.ID() != 1 || p.Status() != -1 || p.PlaceID() != 21 {
		t.Errorf("parking not well created")
	}
	if err != nil {
		t.Errorf("error on getting 1")
	}
}
func TestValidateDelete(t *testing.T) {
	setup()
	defer tearDown()
	parks.Add(21)
	if err := parks.Delete(1); err != nil {
		t.Errorf("deleting 1. Error %s", err.Error())
	}
	err := parks.Delete(1)
	if err == nil {
		t.Errorf("deleting 1 previously deleted")
	}
	if err.Error() != "Parking 1 does not exist" {
		t.Errorf("error message not valid")
	}
	if parks.Exists(1) {
		t.Errorf("parking 1 exists")
	}
	parks.Add(21)
	if err := parks.Delete(2); err != nil {
		t.Errorf("deleting 2. Error %s", err.Error())
	}
	if err := parks.Delete(2); err == nil {
		t.Errorf("deleting 2 previously deleted")
	}
	if parks.Exists(2) {
		t.Errorf("parking 2 exists")
	}
	if num, _ := parks.Len(); num != 0 {
		t.Errorf("map size is not 0")
	}
}

func TestValidateUpdateStatus(t *testing.T) {
	setup()
	defer tearDown()
	parks.Add(21)
	setupMock()
	mockPlaceList[21] = mockPlace{0}
	if mockPlaceList[21].freeParkings != 0 {
		t.Errorf("free parking is not 0")
	}
	tf, err := parks.UpdateStatus(1, -2, mockCallPlaceAPI)
	if tf != false || err == nil {
		t.Errorf("status -2 passes")
	}
	if err.Error() != "Parking status -2 does not exist" {
		t.Errorf("error message not valid")
	}
	tf, err = parks.UpdateStatus(1, 5, mockCallPlaceAPI)
	if tf != false || err == nil {
		t.Errorf("status 5 passes")
	}
	if err.Error() != "Parking status 5 does not exist" {
		t.Errorf("error message not valid")
	}
	if tf, err := parks.UpdateStatus(1, 0, mockCallPlaceAPI); tf != true || err != nil {
		t.Errorf("status 0 not passes")
	}
	if mockPlaceList[21].freeParkings != 1 {
		t.Errorf("no free parking")
	}
	if tf, err := parks.UpdateStatus(1, 1, mockCallPlaceAPI); tf != true || err != nil {
		t.Errorf("status 1 not passes")
	}
	if mockPlaceList[21].freeParkings != 0 {
		t.Errorf("free parking")
	}
	for st := -1; st < 5; st++ {
		if tf, err := parks.UpdateStatus(1, st, mockCallPlaceAPI); tf != true || err != nil {
			t.Errorf("status %d not passes", st)
		}
	}

	tf, err = parks.UpdateStatus(2, 0, mockCallPlaceAPI)
	if tf != false || err == nil {
		t.Errorf("no parking 2 passes")
	}

	parks.Add(3)
	tf, err = parks.UpdateStatus(2, 0, mockCallPlaceAPI)
	if tf != false || err == nil {
		t.Errorf("parking 2 no place passes")
	}
	if err.Error() != "Parking status not updated. Error 404 on Places (Place ID 3)" {
		t.Errorf("error message not valid %s", err.Error())
	}
}
