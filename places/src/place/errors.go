package place

import "fmt"

// https://golangbot.com/custom-errors/

// NotInitialized error
type NotInitialized struct {
}

func (e *NotInitialized) Error() string {
	return fmt.Sprintf("Places not initialized")
}

// AlreadyExists error
type AlreadyExists struct {
	Lat, Long float64
}

// ParkingAPIError error
type ParkingAPIError struct {
	Status int
	Fun    string
}

func (e *ParkingAPIError) Error() string {
	return fmt.Sprintf("Error %d on Parking while calling %s", e.Status, e.Fun)
}

func (e *AlreadyExists) Error() string {
	return fmt.Sprintf("Place in %f,%f already exists", e.Lat, e.Long)
}

// NotExists error
type NotExists struct {
	ID int
}

func (e *NotExists) Error() string {
	return fmt.Sprintf("Place %d does not exist", e.ID)
}

// DeleteParkingNotPossible error
type DeleteParkingNotPossible struct {
	ID int
}

func (e *DeleteParkingNotPossible) Error() string {
	return fmt.Sprintf("Place %d has only one parking and it cannot be deleted \nTry to delete the place", e.ID)
}

// DeleteParkingNotInPlace error
type DeleteParkingNotInPlace struct {
	placeID, parkingID int
}

func (e *DeleteParkingNotInPlace) Error() string {
	return fmt.Sprintf("Place %d does not have %d parking", e.placeID, e.parkingID)
}
