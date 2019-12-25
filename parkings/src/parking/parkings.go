package parking

// Parkings represents a list of Parkings.
type Parkings struct {
	parkingList map[int]Parking
	lastID      int
}

// Init initialize Parkings struct.
// It sets the map and the last ID
func (p *Parkings) Init() {
	p.parkingList = make(map[int]Parking)
	p.lastID = 0
}

// Len returns length of parking list if it is initialized
func (p *Parkings) Len() (int, error) {
	if p.parkingList == nil {
		return 0, &NotInitialized{}
	}
	return len(p.parkingList), nil
}

// Add adds a new parking in the Parkings struct.
// placeID is the ID of the place where parking will be in
func (p *Parkings) Add(placeID int) Parking {
	p.lastID++
	p.parkingList[p.lastID] = Parking{id: p.lastID, status: -1, placeID: placeID}
	return p.parkingList[p.lastID]
}

// Exists checks if a parking is in the Parkings struct
func (p *Parkings) Exists(id int) bool {
	_, ok := p.parkingList[id]
	if ok {
		return true
	}
	return false
}

// Get returns a Parking from the map by id
func (p *Parkings) Get(id int) (Parking, error) {
	if p.Exists(id) {
		return p.parkingList[id], nil
	}
	return Parking{}, &NotExists{id}
}

// Delete delete a parking from the Parkings struct if exists
func (p *Parkings) Delete(id int) error {
	exists := p.Exists(id)
	if exists {
		delete(p.parkingList, id)
		return nil
	}
	return &NotExists{id}
}

// UpdateStatus update the status of a Parking in the Parkings struct if the parking exists and if the status is possible (-1..4)
func (p *Parkings) UpdateStatus(id int, status int, placeAPIUpdateFrees func(int, int) int) (bool, error) {
	_, ok := parkingPosStatus[status]
	if !ok {
		return false, &StatusNotExists{status}
	}
	if !p.Exists(id) {
		return false, &NotExists{id}
	}
	par := p.parkingList[id]
	preStatus := par.status

	// call to places to update free parkings
	responseStatus := 200
	if preStatus == 0 && status != 0 {
		responseStatus = placeAPIUpdateFrees(par.placeID, -1)
	} else if preStatus != 0 && status == 0 {
		responseStatus = placeAPIUpdateFrees(par.placeID, 1)
	}
	if responseStatus != 200 {
		return false, &StatusNotUpdated{responseStatus, par.placeID}
	}
	par.status = status
	p.parkingList[id] = par
	return true, nil
}
