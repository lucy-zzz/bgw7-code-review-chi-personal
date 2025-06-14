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

func (s *VehicleDefault) Create(new internal.VehicleAttributes) (err error) {
	err = s.rp.Create(new)

	if err != nil {
		return err
	}
	return nil
}

func (s *VehicleDefault) FindByColorAndYear(vehicle internal.VehicleAttributes) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByColorAndYear(vehicle)

	if err != nil {
		return v, err
	}

	return v, nil
}

func (s *VehicleDefault) FindByBrandAndYearInterval(r internal.BrandYearRangeSearchType) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByBrandAndYearInterval(r)

	if err != nil {
		return nil, err
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

func (s *VehicleDefault) CreateSome(vs []internal.VehicleAttributes) (err error) {
	err = s.rp.CreateSome(vs)

	if err != nil {
		return err
	}

	return nil
}

func (s *VehicleDefault) UpdateSpeed(v internal.UpdateSpeed) (err error) {
	err = s.rp.UpdateSpeed(v)

	if err != nil {
		return err
	}

	return nil
}

func (s *VehicleDefault) GetByFuelType(t string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByFuelType(t)

	if err != nil {
		return v, err
	}

	return v, nil
}

func (s *VehicleDefault) DeleteById(id int) (err error) {
	err = s.rp.DeleteById(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *VehicleDefault) GetByTransmissionType(t string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByTransmissionType(t)

	return v, err
}
