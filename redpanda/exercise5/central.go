package exercise5

type Driver struct {
	ID string
	Available bool
}

type Rider struct {
	ID string
	InRide bool
}

type Ride struct {
	DriverID string
	RiderID string
	Status string
}

type CentralManagement struct {
	Drivers []Driver
	Riders []Rider
	Rides []Ride
}