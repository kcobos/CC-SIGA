package parking

import "fmt"

// https://golangbot.com/custom-errors/

// NotInitialized error
type NotInitialized struct {
}

func (e *NotInitialized) Error() string {
	return fmt.Sprintf("Parkings not initialized")
}

// NotExists error
type NotExists struct {
	ID int
}

func (e *NotExists) Error() string {
	return fmt.Sprintf("Parking %d does not exist", e.ID)
}

// StatusNotExists error
type StatusNotExists struct {
	Status int
}

func (e *StatusNotExists) Error() string {
	return fmt.Sprintf("Parking status %d does not exist", e.Status)
}

// StatusNotUpdated error
type StatusNotUpdated struct {
	Status, IDplace int
}

func (e *StatusNotUpdated) Error() string {
	return fmt.Sprintf("Parking status not updated. Error %d on Places (Place ID %d)", e.Status, e.IDplace)
}
