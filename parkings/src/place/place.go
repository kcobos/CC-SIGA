package place

import "../parking"

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

// OneFree liberates one parking
func (p *Place) OneFree() {
	p.freeParkings++
}

// OneOccupied occupies one parking
func (p *Place) OneOccupied() {
	p.freeParkings--
}

// Places represents a list of Places.
// It has a map an a variable to know the last ID
type Places struct {
	placeList map[int]Place
	lastID    int
}

// Init initialize Places and Parkings struct.
// It sets the map and the last ID
func (p *Places) Init(ps *parking.Parkings) {
	p.placeList = make(map[int]Place)
	p.lastID = 0
	ps.Init()
}

// Len returns length of place list if it is initialized
func (p *Places) Len() (int, error) {
	if p.placeList == nil {
		return 0, &NotInitialized{}
	}
	return len(p.placeList), nil
}

// AddParking adds a parking lot in a place
func (p *Place) AddParking(ps *parking.Parkings) {
	p.numParkings++
	pa := ps.AddParking(p.id)
	p.parkings = append(p.parkings, pa.ID())
}

// DeleteParking deletes a parking lot in a place
func (p *Place) DeleteParking(id int, ps *parking.Parkings) error {
	if p.numParkings == 1 {
		return &DeleteParkingNotPossible{p.id}
	}
	parkIndex := -1
	// Search index https://stackoverflow.com/questions/38654383/how-to-search-for-an-element-in-a-golang-slice
	for i := range p.parkings {
		if p.parkings[i] == id {
			parkIndex = i
			break
		}
	}

	if parkIndex != -1 && p.parkings[parkIndex] == id {
		err := ps.DeleteParking(id)
		if err == nil {
			p.numParkings--
			// Delete id https://yourbasic.org/golang/delete-element-slice/
			p.parkings[parkIndex] = p.parkings[len(p.parkings)-1]
			p.parkings = p.parkings[:len(p.parkings)-1]
		}
		return err
	}
	return &DeleteParkingNotInPlace{p.id, id}
}

// CreatePlace create a parking place and adds a parking
func (p *Places) CreatePlace(lat float64, long float64, address string, ps *parking.Parkings) (Place, error) {
	for _, pla := range p.placeList {
		if pla.lat == lat && pla.long == long {
			return Place{}, &AlreadyExists{lat, long}
		}
	}
	p.lastID++
	pla := Place{id: p.lastID, lat: lat, long: long, address: address}

	pla.AddParking(ps)
	p.placeList[p.lastID] = pla
	return p.placeList[p.lastID], nil
}

// PlaceExists checks if a place is in the Places struct
func (p *Places) PlaceExists(id int) bool {
	_, ok := p.placeList[id]
	if ok {
		return true
	}
	return false
}

// Get returns a place from the map by id
func (p *Places) Get(id int) (Place, error) {
	if p.PlaceExists(id) {
		return p.placeList[id], nil
	}
	return Place{}, &NotExists{id}
}

// DeletePlace deletes a place and all its parkings
func (p *Places) DeletePlace(id int, ps *parking.Parkings) (bool, error) {
	pla, ok := p.placeList[id]
	if !ok {
		return false, &NotExists{id}
	}
	for len(pla.parkings) != 1 {
		err := pla.DeleteParking(pla.parkings[0], ps)
		if err != nil {
			return false, err
		}
	}
	err := ps.DeleteParking(id)
	if err == nil {
		pla.numParkings--
		pla.parkings = pla.parkings[:0]
	}

	delete(p.placeList, id)
	return true, nil
}

// GetFrees return all places which have, at least, one free parking
func (p *Places) GetFrees() map[int]Place {
	free := make(map[int]Place)
	for id, pla := range p.placeList {
		if pla.freeParkings != 0 {
			free[id] = pla
		}
	}
	return free
}

// Get Window
// Get Nearest frees
