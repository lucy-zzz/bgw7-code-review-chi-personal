package repository

import (
	"app/internal"
	"fmt"
)

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

func (r *VehicleMap) GetAverageSpeedByBrand(b string) (v float64, err error) {
	var brandList []internal.Vehicle
	for _, i := range r.db {
		if b == i.Brand {
			brandList = append(brandList, i)
		}
	}

	if len(brandList) == 0 {
		return 0, err
	}

	var sumSpeed float64

	for _, bl := range brandList {
		sumSpeed += bl.MaxSpeed
	}

	return sumSpeed / float64(len(brandList)), nil
}

func (r *VehicleMap) Create(v internal.VehicleAttributes) (err error) {
	vehicleList := make(map[int]internal.Vehicle)

	maxKey := 0
	for _, v := range r.db {
		if v.Id > maxKey {
			maxKey = v.Id
		}
	}

	newID := maxKey + 1

	if _, exists := r.db[newID]; exists {
		return err
	}

	vehicleList[newID] = internal.Vehicle{
		Id:                newID,
		VehicleAttributes: v,
	}

	return nil
}

func (r *VehicleMap) CreateSome(vs []internal.VehicleAttributes) (err error) {
	vehicleList := make(map[int]internal.Vehicle)

	maxKey := 0
	for k := range r.db {
		if k > maxKey {
			maxKey = k
		}
	}

	for i, v := range vs {
		newID := maxKey + 1 + i

		if _, exists := r.db[newID]; exists {
			return err
		}

		vehicleList[i] = internal.Vehicle{
			Id:                newID,
			VehicleAttributes: v,
		}
	}

	return nil
}

func (r *VehicleMap) UpdateSpeed(v internal.UpdateSpeed) (err error) {
	var vehicle internal.Vehicle

	for i := 0; i <= len(r.db); i++ {
		if r.db[i].Id == v.Id {
			vehicle = r.db[i]
			vehicle.MaxSpeed = v.Speed
			break
		}
	}

	if vehicle.Id == 0 {
		return err
	}

	return nil
}

func (r *VehicleMap) GetByFuelType(t string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, i := range r.db {
		if i.FuelType == t {
			v[key] = i
		}
	}

	if len(v) == 0 {
		return v, err
	}

	fmt.Println("len", len(v))

	return v, nil
}

func (r *VehicleMap) DeleteById(id int) (err error) {
	found := false
	db := r.db
	for key := range r.db {
		if key == id {
			delete(db, key)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("404 Not Found: Veículo não encontrado.")
	}

	return nil
}

func (r *VehicleMap) GetByTransmissionType(t string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, vs := range r.db {
		if vs.Transmission == t {
			v[key] = vs
		}
	}

	if len(v) == 0 {
		return v, fmt.Errorf("404 Not Found: Não foram encontrados veículos com esse tipo de transmissão.")
	}

	return v, err
}

func (r *VehicleMap) UpdateFuelType(u internal.UpdateFuel) (err error) {
	found := false

	for _, vs := range r.db {
		if vs.Id == u.Id {
			found = true
			temp := vs
			temp.FuelType = u.FuelType
			fmt.Println("vehicle", temp)
		}
	}

	if !found {
		return fmt.Errorf("404 Not Found: Veículo não encontrado")
	}

	return nil
}

func (r *VehicleMap) GetAverageCapacityByBrand(b string) (v float64, err error) {
	var sum int
	var list []internal.Vehicle
	for _, i := range r.db {
		if i.Brand == b {
			sum += i.Capacity
			list = append(list, i)
		}
	}

	if len(list) == 0 {
		return 0, fmt.Errorf("404 Not Found: Não foram encontrados veículos dessa marca.")
	}

	v = float64(sum) / float64(len(list))

	return v, err
}

func (r *VehicleMap) GetByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, i := range r.db {
		if i.Dimensions.Length >= minLength && i.Dimensions.Length <= maxLength {
			if i.Dimensions.Width >= minWidth && i.Dimensions.Width <= maxWidth {
				v[key] = i
			}
		}
	}

	if len(v) == 0 {
		return v, fmt.Errorf("404 Not Found: Não foram encontrados veículos com essas dimensões.")
	}

	return v, nil
}

func (r *VehicleMap) GetByWeight(minW, maxW float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, i := range r.db {
		if i.Weight >= minW && i.Weight <= maxW {
			v[key] = i
		}
	}

	if len(v) == 0 {
		return v, fmt.Errorf("404 Not Found: Não foram encontrados veículos nessa faixa de peso.")
	}

	return v, nil
}
