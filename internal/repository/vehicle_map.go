package repository

import "app/internal"

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) FindByColorAndYear(vehicle internal.VehicleAttributes) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if vehicle.Color == value.Color && vehicle.FabricationYear == value.FabricationYear {
			v[key] = value
		}
	}

	if len(v) == 0 {
		return v, err
	}

	return v, nil
}

func (r *VehicleMap) FindByBrandAndYearInterval(req internal.BrandYearRangeSearchType) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Brand == req.Brand {
			if value.FabricationYear >= req.StartYear && value.FabricationYear <= req.EndYear {
				v[key] = value
			}
		}
	}

	if len(v) == 0 {
		return v, err
	}

	return v, nil
}
