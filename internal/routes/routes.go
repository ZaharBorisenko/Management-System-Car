package routes

import (
	carHandler "github.com/ZaharBorisenko/Management-System-Car/internal/handler/car"
	engineHandler "github.com/ZaharBorisenko/Management-System-Car/internal/handler/engine"
	"net/http"
)

func RegisterRoutes(car *carHandler.CarHandler, engine *engineHandler.EngineHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /cars", car.GetAllCar)
	mux.HandleFunc("GET /cars/{id}", car.GetCarById)
	mux.HandleFunc("GET /cars/vin/{vin}", car.GetCarByVinCode)
	mux.HandleFunc("GET /cars/brand/{brand}", car.GetCarByBrand)
	mux.HandleFunc("POST /cars", car.CreateCar)
	mux.HandleFunc("DELETE /cars/delete/{id}", car.DeleteCar)
	mux.HandleFunc("PATCH  /cars/update/{id}", car.UpdateCar)

	mux.HandleFunc("GET /engine/{id}", engine.GetEngineById)
	mux.HandleFunc("POST /engine", engine.CreateEngine)

	return mux
}
