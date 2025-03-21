package exercise5

import "sync"

type Driver struct {
	ID        int  `json:"id"`
	Available bool `json:"available"`
}

type Rider struct {
	ID     int  `json:"id"`
	InRide bool `json:"inRide"`
}

type Ride struct {
	DriverID int `json:"driverId"`
	RiderID  int `json:"riderId"`
	Status   string `json:"status"`
}

type CentralManagement struct {
	Drivers []Driver
	Riders  []Rider
	Rides   []Ride
	mu      sync.Mutex
}

type RideUpdate struct {
	Ride Ride `json:"ride"`
	Status string `json:"status"`
}

// Simulate a database
func (cm *CentralManagement) StartCentral(drivers int, riders int) {
	cm.StartDrivers(drivers)
	cm.StartRiders(riders)
}

func (cm *CentralManagement) StartDrivers(drivers int) {
	for i := 0; i < drivers; i++ {
		newDriver := Driver{
			ID:        i + 1,
			Available: false,
		}

		cm.Drivers = append(cm.Drivers, newDriver)
	}
}

func (cm *CentralManagement) StartRiders(riders int) {
	for i := 0; i < riders; i++ {
		newRider := Rider{
			ID:        i + 1,
			InRide: false,
		}

		cm.Riders = append(cm.Riders, newRider)
	}
}

func (cm *CentralManagement) ChangeDriverAvailability(driver *Driver) {
	cm.mu.Lock()   // Lock the mutex
    defer cm.mu.Unlock() // Unlock the mutex when the function completes

	for _, d := range cm.Drivers {
		if d.ID == driver.ID {
			d.Available = !d.Available
			break
		}
	}
}

func (cm *CentralManagement) CreateRide(rider *Rider) {
	cm.mu.Lock()   // Lock the mutex
    defer cm.mu.Unlock() // Unlock the mutex when the function completes

	for _, r := range cm.Riders {
		if r.ID == rider.ID {
			r.InRide = true
		}
	}

	newRide := &Ride{
		RiderID: rider.ID,
		Status: "Matching",
	}
	cm.Rides = append(cm.Rides, *newRide)
}

func (cm *CentralManagement) UpdateRideState(ride *Ride, status string) {
	cm.mu.Lock()   // Lock the mutex
    defer cm.mu.Unlock() // Unlock the mutex when the function completes

	for _, r := range cm.Rides {
		if r.RiderID == ride.RiderID {
			r.Status = status
		}
	}
}

func (cm *CentralManagement) MatchRide(riderID int) (*Ride) {
	cm.mu.Lock()   // Lock the mutex
    defer cm.mu.Unlock() // Unlock the mutex when the function completes

	var availableDriver *Driver
	for _, driver := range cm.Drivers {
		if driver.Available {
			availableDriver = &driver
			break
		}
	}

	var currentRider *Rider
	for _, rider := range cm.Riders {
		if rider.ID == riderID {
			currentRider = &rider
			break
		}
	}

	var finalRide *Ride
	for _, ride := range cm.Rides {
		if ride.RiderID == riderID {
			ride.DriverID = availableDriver.ID
			availableDriver.Available = false
			currentRider.InRide = true
			finalRide = &ride
			break
		}
	}
	return finalRide
}

var Central = CentralManagement{
	Drivers: []Driver{},
	Riders:  []Rider{},
	Rides:   []Ride{},
}
