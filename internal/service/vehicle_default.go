package service

import (
	"app/internal"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) Create(new internal.VehicleAttributes) (v internal.Vehicle, err error) {
	all, err := s.rp.FindAll()

	if err != nil {
		return internal.Vehicle{}, err
	}

	v = internal.Vehicle{
		Id:                getNextID(all),
		VehicleAttributes: new,
	}
	return v, nil
}

func getNextID(vehicles map[int]internal.Vehicle) int {
	maxID := 0
	for _, v := range vehicles {
		if v.Id > maxID {
			maxID = v.Id
		}
	}
	return maxID + 1
}

func (s *VehicleDefault) FindByColorAndYear(vehicle internal.VehicleAttributes) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByColorAndYear(vehicle)

	if err != nil {
		return map[int]internal.Vehicle{}, err
	}

	return v, nil
}

func (s *VehicleDefault) FindByBrandAndYearInterval(r internal.BrandYearRangeSearchType) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByBrandAndYearInterval(r)

	if err != nil {
		return map[int]internal.Vehicle{}, err
	}

	return v, nil
}

func (s *VehicleDefault) GetAverageSpeedByBrand(b string) (v float64, err error) {
	v, err = s.rp.GetAverageSpeedByBrand(b)

	if err != nil {
		return 0, err
	}

	return v, nil
}
