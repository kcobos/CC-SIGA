package parking

var parkingPosStatus map[int]string

func init() {
	parkingPosStatus = make(map[int]string)
	parkingPosStatus[-1] = "undefined"
	parkingPosStatus[0] = "vacant"
	parkingPosStatus[1] = "taken"
	parkingPosStatus[2] = "wrong"
	parkingPosStatus[3] = "notified"
	parkingPosStatus[4] = "inProcess"
}

// Parking represents a parking lot.
type Parking struct {
	id, status, placeID int
}

// ID returns parking's id
func (p *Parking) ID() int {
	return p.id
}

// Status returns parking's status
func (p *Parking) Status() int {
	return p.status
}

// PlaceID returns parking's place id
func (p *Parking) PlaceID() int {
	return p.placeID
}

