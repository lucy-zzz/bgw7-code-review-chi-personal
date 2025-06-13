package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Create(newVehicle VehicleAttributes) (err error)
	FindByColorAndYear(vehicle VehicleAttributes) (v map[int]Vehicle, err error)
	FindByBrandAndYearInterval(r BrandYearRangeSearchType) (v map[int]Vehicle, err error)
	GetAverageSpeedByBrand(b string) (v float64, err error)
	CreateSome(vs []VehicleAttributes) (err error)
}
