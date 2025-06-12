package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	FindByColorAndYear(vehicle VehicleAttributes) (v map[int]Vehicle, err error)
	FindByBrandAndYearInterval(r BrandYearRangeSearchType) (v map[int]Vehicle, err error)
	GetAverageSpeedByBrand(b string) (v float64, err error)
}
