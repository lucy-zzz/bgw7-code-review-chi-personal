package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the application
func (a *ServerChi) Run() (err error) {
	// dependencies
	// - loader
	ld := loader.NewVehicleJSONFile(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}
	// - repository
	rp := repository.NewVehicleMap(db)
	// - service
	sv := service.NewVehicleDefault(rp)
	// - handler
	hd := handler.NewVehicleDefault(sv)
	// router
	rt := chi.NewRouter()
	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
	// - endpoints
	rt.Route("/vehicles", func(rt chi.Router) {
		// - GET /vehicles
		rt.Get("/", hd.GetAll())
		rt.Post("/", hd.Create())
		rt.Get("/vehiclesc", hd.GetByColorAndYear())
		rt.Get("/brand/{brand}/between/{start_year}/{end_year}", hd.GetByBrandAndYearInterval())
		rt.Get("/average_speed/brand/{brand}", hd.GetAverageSpeedByBrand())
		rt.Post("/batch", hd.CreateSome())
		rt.Put("/{id}/update_speed", hd.UpdateSpeed())
		rt.Get("/fuel_type/{type}", hd.GetByFuelType())
		rt.Delete("/{id}", hd.DeleteById())
		rt.Get("/transmission/{type}", hd.GetByTransmissionType())
		rt.Put("/{id}/update_fuel", hd.UpdateFuel())
		rt.Get("/average_capacity/brand/{brand}", hd.GetAverageCapacityByBrand())
		rt.Get("/dimensions", hd.GetByDimensions())
		rt.Get("/weight", hd.GetByWeight())
	})

	rt.Route("/vehiclesc", func(rt chi.Router) {
		rt.Get("/", hd.GetByColorAndYear())
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
