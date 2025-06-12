package handler

import (
	"app/internal"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input internal.VehicleAttributes

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
			response.JSON(w, http.StatusBadRequest, 400)
			return
		}

		newVehicleAttributes := internal.VehicleAttributes{
			Brand:           input.Brand,
			Model:           input.Model,
			Registration:    input.Registration,
			Color:           input.Color,
			FabricationYear: input.FabricationYear,
			Capacity:        input.Capacity,
			MaxSpeed:        input.MaxSpeed,
			FuelType:        input.FuelType,
			Transmission:    input.Transmission,
			Weight:          input.Weight,
			Dimensions:      input.Dimensions,
		}

		id, err := h.sv.Create(newVehicleAttributes)
		if err != nil {
			w.Write([]byte(`{message: 409 Conflict: Identificador do veículo já existente.}`))
			response.JSON(w, http.StatusBadRequest, 400)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    id,
		})

	}
}

func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := r.URL.Query().Get("color")
		yearStr := r.URL.Query().Get("year")
		year, err := strconv.Atoi(yearStr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"Parâmetro year inválido ou ausente."}`))
			return
		}

		input := internal.VehicleAttributes{
			Color:           color,
			FabricationYear: year,
		}

		vehicle := internal.VehicleAttributes{
			Color:           input.Color,
			FabricationYear: input.FabricationYear,
		}

		vehiclesList, err := h.sv.FindByColorAndYear(vehicle)

		if err != nil {
			w.Write([]byte(`{message: 404 Not Found: Nenhum veículo encontrado com esses critérios. }`))
			response.JSON(w, http.StatusNotFound, 404)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vehiclesList,
		})
	}
}

func (h *VehicleDefault) GetByBrandAndYearInterval() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		brand := chi.URLParam(r, "brand")
		startYearStr := chi.URLParam(r, "start_year")
		endYearStr := chi.URLParam(r, "end_year")

		startYear, err := strconv.Atoi(startYearStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"Parâmetro start_year inválido ou ausente."}`))
			return
		}

		endYear, err := strconv.Atoi(endYearStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"Parâmetro end_year inválido ou ausente."}`))
			return
		}

		req := internal.BrandYearRangeSearchType{
			Brand:     brand,
			StartYear: startYear,
			EndYear:   endYear,
		}

		vehiclesList, err := h.sv.FindByBrandAndYearInterval(req)

		if err != nil {
			w.Write([]byte(`{message: 404 Not Found: Nenhum veículo encontrado com esses critérios. }`))
			response.JSON(w, http.StatusNotFound, 404)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vehiclesList,
		})
	}
}

func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		averageSpeed, err := h.sv.GetAverageSpeedByBrand(brand)

		if err != nil {
			w.Write([]byte(`{message: 404 Not Found: Nenhum veículo encontrado dessa marca.}`))
			response.JSON(w, http.StatusNotFound, 404)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    averageSpeed})

	}
}
