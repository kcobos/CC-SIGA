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
	status int
}

func (e *StatusNotExists) Error() string {
	return fmt.Sprintf("Parking status %d does not exist", e.status)
}
