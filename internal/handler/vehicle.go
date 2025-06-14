package handler

import (
	"app/internal"
	"encoding/json"
	"fmt"
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

		err = h.sv.Create(newVehicleAttributes)
		if err != nil {
			w.Write([]byte(`{message: 409 Conflict: Identificador do veículo já existente.}`))
			response.JSON(w, http.StatusBadRequest, 400)
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
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
			w.Write([]byte(`{"message": "404 Not Found: Nenhum veículo encontrado com esses critérios." }`))
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
			w.Write([]byte(`{"message": "404 Not Found: Nenhum veículo encontrado com esses critérios." }`))
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
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    averageSpeed})

	}
}

func (h *VehicleDefault) CreateSome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input []internal.VehicleAttributes
		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			w.Write([]byte(`"message": "400 Bad Request: Dados de algum veículo malformados ou incompletos."`))
			response.JSON(w, http.StatusBadRequest, 400)
			return
		}

		err = h.sv.CreateSome(input)

		if err != nil {
			w.Write([]byte(`"message": "409 Conflict: Algum veículo possui um identificador já existente."`))
			response.JSON(w, http.StatusConflict, 409)
			return
		}

		response.JSON(w, http.StatusCreated, 201)
	}
}

func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s struct {
			Speed float64 `json:"speed"`
		}

		err := json.NewDecoder(r.Body).Decode(&s)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, `{"message": "400 Bad Request: Velocidade malformada ou fora de alcance."}`)
			return
		}

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			w.Write([]byte(`400 Bad Request: Velocidade malformada ou fora de alcance.`))
			response.JSON(w, http.StatusBadRequest, 400)
			return
		}

		var u internal.UpdateSpeed
		u = internal.UpdateSpeed{
			Id:    int(id),
			Speed: s.Speed,
		}

		fmt.Println(u, "struct on handler")

		err = h.sv.UpdateSpeed(u)

		if err != nil {
			w.Write([]byte(`404 Not Found: Veículo não encontrado.`))
			response.JSON(w, http.StatusNotFound, 404)
			return
		}

		response.JSON(w, http.StatusOK, 200)
	}
}

func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuelType := chi.URLParam(r, "type")

		data, err := h.sv.GetByFuelType(fuelType)

		if err != nil {
			w.Write([]byte(`404 Not Found: Não foram encontrados veículos com esse tipo de combustível.`))
			response.JSON(w, http.StatusNotFound, 404)
			return
		}

		if len(data) == 0 {
			w.WriteHeader(http.StatusNotFound)
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": "404 Not Found: Não foram encontrados veículos com esse tipo de combustível.",
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
		return
	}
}

func (h *VehicleDefault) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
			response.JSON(w, http.StatusBadRequest, 400)
			return
		}

		err = h.sv.DeleteById(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "204 No Content: Veículo removido com sucesso.",
		})
	}
}

func (h *VehicleDefault) GetByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := chi.URLParam(r, "type")

		data, err := h.sv.GetByTransmissionType(t)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			w.Write([]byte(`{message: 404 Not Found: Não foram encontrados veículos com esse tipo de transmissão.}`))
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
