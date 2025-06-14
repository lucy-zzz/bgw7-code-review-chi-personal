package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	FindByColorAndYear(vehicle VehicleAttributes) (v map[int]Vehicle, err error)
	FindByBrandAndYearInterval(r BrandYearRangeSearchType) (v map[int]Vehicle, err error)
	GetAverageSpeedByBrand(b string) (v float64, err error)
	Create(v VehicleAttributes) (err error)
	CreateSome(vs []VehicleAttributes) (err error)
	UpdateSpeed(v UpdateSpeed) (err error)
	GetByFuelType(t string) (v map[int]Vehicle, err error)
	DeleteById(id int) (err error)
	GetByTransmissionType(t string) (v map[int]Vehicle, err error)
	UpdateFuelType(u UpdateFuel) (err error)
	GetAverageCapacityByBrand(b string) (v float64, err error)
	GetByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]Vehicle, err error)
}
