package place

// Place represents a parking location which has several parking lot.
type Place struct {
	id                        int
	lat, long                 float64
	address                   string
	numParkings, freeParkings int
	parkings                  []int
}

// ID returns place's id
func (p *Place) ID() int {
	return p.id
}

// Coor returns place's coordinates (lat,long)
func (p *Place) Coor() (float64, float64) {
	return p.lat, p.long
}

// Address returns place's address
func (p *Place) Address() string {
	return p.address
}

// NumParkings returns place's num parking
func (p *Place) NumParkings() int {
	return p.numParkings
}

// FreeParkings returns place's num free parkings
func (p *Place) FreeParkings() int {
	return p.freeParkings
}

// OneFreed liberates one parking
func (p *Place) OneFreed() {
	p.freeParkings++
}

// OneOccupied occupies one parking
func (p *Place) OneOccupied() {
	p.freeParkings--
}
