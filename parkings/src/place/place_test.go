package place

import (
	"fmt"
	"testing"

	"../parking"
)

func TestValidateInit(t *testing.T) {
	var placs Places
	var parks parking.Parkings
	parks.Init()
	_, err := placs.Len()
	if err == nil {
		t.Errorf("map is initialized")
	}
	if err.Error() != "Places not initialized" {
		t.Errorf("error message not valid")
	}
	placs.Init(&parks)
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

var parks parking.Parkings
var placs Places

func setup() {
	parks.Init()
	placs.Init(&parks)
}

func TestValidateCreatePlace(t *testing.T) {
	setup()
	p, err := placs.CreatePlace(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", &parks)
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

	p, err = placs.CreatePlace(37.19742395414327, -3.624779980964916,
		"Calle Periodista Eugenio Selles, 18014 Granada", &parks)
	if err == nil {
		t.Errorf("error creating duplicate coordinates place")
	}
	if err.Error() != fmt.Sprintf("Place in %f,%f already exists", 37.19742395414327, -3.624779980964916) {
		t.Errorf("error message not valid")
	}
}

func TestValidatePlaceExists(t *testing.T) {
	setup()
	placs.CreatePlace(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", &parks)
	placs.CreatePlace(37.195458, -3.626633,
		"Calle Periodista Eugenio Selles, 18014 Granada", &parks)
	if !placs.PlaceExists(1) {
		t.Errorf("place 1 not exists")
	}
	if !placs.PlaceExists(2) {
		t.Errorf("place 2 not exists")
	}
	if placs.PlaceExists(5) {
		t.Errorf("place 5 exists")
	}
	if num, _ := placs.Len(); num != 2 {
		t.Errorf("map size is not 2")
	}
	if num, _ := parks.Len(); num != 2 {
		t.Errorf("map size is not 2")
	}
}

func TestValidateAddParking(t *testing.T) {
	setup()
	place, _ := placs.CreatePlace(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", &parks)
	place.AddParking(&parks)
	if place.NumParkings() != 2 {
		t.Errorf("no parking added")
	}
	if num, _ := parks.Len(); num != 2 {
		t.Errorf("map size is not 2")
	}
	place.AddParking(&parks)
	if place.NumParkings() != 3 {
		t.Errorf("no parking added")
	}
	if num, _ := parks.Len(); num != 3 {
		t.Errorf("map size is not 3")
	}
	if num, _ := placs.Len(); num != 1 {
		t.Errorf("map size is not 1")
	}
}

func TestValidateDeleteParking(t *testing.T) {
	setup()
	place, _ := placs.CreatePlace(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", &parks)
	place.AddParking(&parks)
	place.AddParking(&parks)
	if place.NumParkings() != 3 {
		t.Errorf("no parking added")
	}
	if num, _ := parks.Len(); num != 3 {
		t.Errorf("map size is not 3")
	}
	if num, _ := placs.Len(); num != 1 {
		t.Errorf("map size is not 1")
	}

	err := place.DeleteParking(1, &parks)
	if err != nil {
		t.Errorf("error on deleting parking")
	}
	if place.NumParkings() != 2 {
		t.Errorf("no parking deleted")
	}
	if num, _ := parks.Len(); num != 2 {
		t.Errorf("map size is not 2")
	}

	err = place.DeleteParking(1, &parks)
	if err == nil {
		t.Errorf("error on deleting deleted parking")
	}
	if err.Error() != "Place 1 does not have 1 parking" {
		t.Errorf("error message not valid")
	}
	if place.NumParkings() != 2 {
		t.Errorf("no parking deleted")
	}
	if num, _ := parks.Len(); num != 2 {
		t.Errorf("map size is not 2")
	}

	err = place.DeleteParking(2, &parks)
	if err != nil {
		t.Errorf("error on deleting parking")
	}
	if place.NumParkings() != 1 {
		t.Errorf("no parking deleted")
	}
	if num, _ := parks.Len(); num != 1 {
		t.Errorf("map size is not 1")
	}

	err = place.DeleteParking(3, &parks)
	if err == nil {
		t.Errorf("error on deleting last parking")
	}
	if err.Error() != "Place 1 has only one parking and it cannot be deleted \nTry to delete the place" {
		t.Errorf("error message not valid")
	}
}

func TestValidateGet(t *testing.T) {
	setup()
	p, _ := placs.CreatePlace(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", &parks)
	pFL, err := placs.Get(1)
	if err != nil {
		t.Errorf("map not contains 1")
	}
	if num, _ := parks.Len(); num != 1 {
		t.Errorf("map size is not 1")
	}
	lat, long := p.Coor()
	latFL, longFL := pFL.Coor()
	if p.ID() != pFL.ID() || p.Address() != pFL.Address() || lat != latFL || long != longFL ||
		p.NumParkings() != pFL.NumParkings() || p.FreeParkings() != pFL.FreeParkings() {
		t.Errorf("places not equal")
	}
	notP, err := placs.Get(12)
	latNotP, longNotP := notP.Coor()
	if err == nil {
		t.Errorf("map contains 12")
	}
	if err.Error() != "Place 12 does not exist" {
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

func TestValidateDeletePlace(t *testing.T) {
	setup()
	place, _ := placs.CreatePlace(37.19742395414327, -3.624779980964916,
		"Calle Periodista Daniel Saucedo Aranda, 18014 Granada", &parks)
	place.AddParking(&parks)
	place.AddParking(&parks)
	if place.NumParkings() != 3 {
		t.Errorf("no parking added")
	}
	if num, _ := parks.Len(); num != 3 {
		t.Errorf("map size is not 3")
	}
	if num, _ := placs.Len(); num != 1 {
		t.Errorf("map size is not 1")
	}
	bo, err := placs.DeletePlace(1, &parks)
	if !bo || err != nil {
		t.Errorf("error deleting place 1")
	}

	bo, err = placs.DeletePlace(1, &parks)
	if bo || err == nil {
		t.Errorf("error deleting place 1 not exists")
	}
	if err.Error() != "Place 1 does not exist" {
		t.Errorf("error message not valid")
	}
}

func TestValidate_(t *testing.T) {
	setup()
}
