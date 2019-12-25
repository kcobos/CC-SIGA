package place

// Places represents a list of Places.
// It has a map an a variable to know the last ID
type Places struct {
	placeList map[int]Place
	lastID    int
}

// Init initialize Places and Parkings struct.
// It sets the map and the last ID
func (p *Places) Init() {
	p.placeList = make(map[int]Place)
	p.lastID = 0
}

// Len returns length of place list if it is initialized
func (p *Places) Len() (int, error) {
	if p.placeList == nil {
		return 0, &NotInitialized{}
	}
	return len(p.placeList), nil
}

// AddParking adds a parking lot in a place
func (p *Place) AddParking(parkingAddAPI func(int) (int, int)) error {
	par, responseStatus := parkingAddAPI(p.id)
	if responseStatus != 200 {
		return &ParkingAPIError{responseStatus, "AddParking"}
	}
	p.numParkings++
	p.parkings = append(p.parkings, par)
	return nil
}

// DeleteParking deletes a parking lot in a place
func (p *Place) DeleteParking(id int, parkingDeleteAPI func(int) int) error {
	if p.numParkings == 1 {
		return &DeleteParkingNotPossible{p.id}
	}
	parkIndex := -1
	for i := range p.parkings {
		if p.parkings[i] == id {
			parkIndex = i
			break
		}
	}

	if parkIndex != -1 && p.parkings[parkIndex] == id {
		responseStatus := parkingDeleteAPI(id)
		if responseStatus == 200 {
			p.numParkings--
			p.parkings[parkIndex] = p.parkings[len(p.parkings)-1]
			p.parkings = p.parkings[:len(p.parkings)-1]
			return nil
		}
		return &ParkingAPIError{responseStatus, "DeleteParking"}
	}
	return &DeleteParkingNotInPlace{p.id, id}
}

// Add create a parking place and adds a parking
func (p *Places) Add(lat float64, long float64, address string, parkingAddAPI func(int) (int, int)) (Place, error) {
	for _, pla := range p.placeList {
		if pla.lat == lat && pla.long == long {
			return Place{}, &AlreadyExists{lat, long}
		}
	}
	p.lastID++
	pla := Place{id: p.lastID, lat: lat, long: long, address: address}
	err := pla.AddParking(parkingAddAPI)
	if err == nil {
		p.placeList[p.lastID] = pla
		return p.placeList[p.lastID], nil
	}
	delete(p.placeList, p.lastID)
	p.lastID--
	return Place{}, err
}

// Exists checks if a place is in the Places struct
func (p *Places) Exists(id int) bool {
	_, ok := p.placeList[id]
	if ok {
		return true
	}
	return false
}

// Get returns a place from the map by id
func (p *Places) Get(id int) (Place, error) {
	if p.Exists(id) {
		return p.placeList[id], nil
	}
	return Place{}, &NotExists{id}
}

// Delete deletes a place and all its parkings
func (p *Places) Delete(id int, parkingDeleteAPI func(int) int) error {
	pla, ok := p.placeList[id]
	if !ok {
		return &NotExists{id}
	}
	for len(pla.parkings) != 1 {
		err := pla.DeleteParking(pla.parkings[0], parkingDeleteAPI)
		if err != nil {
			return err
		}
	}
	responseStatus := parkingDeleteAPI(id)
	if responseStatus == 200 {
		pla.numParkings--
		pla.parkings = pla.parkings[:0]
	} else {
		return &ParkingAPIError{responseStatus, "AddParking"}
	}

	// if len(pla.parkings) != 0 || pla.numParkings != 0 {
	// 	return &DeleteParkingNotPossible{-1}
	// }
	delete(p.placeList, id)
	return nil
}

// GetFrees return all places which have, at least, one free parking
// func (p *Places) GetFrees() map[int]Place {
// 	free := make(map[int]Place)
// 	for id, pla := range p.placeList {
// 		if pla.freeParkings != 0 {
// 			free[id] = pla
// 		}
// 	}
// 	return free
// }

// Get Window
// Get Nearest frees
