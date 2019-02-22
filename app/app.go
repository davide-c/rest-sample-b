package app

import (
	"fmt"
	"github.com/davide-c/rest-sample-b/config"
	"github.com/davide-c/rest-sample-b/handler"
	"github.com/davide-c/rest-sample-b/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)

	db.LogMode(true)

	// fmt.Printf("%+v\n", db)

	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRoutes()
}

func (a *App) setRoutes() {
	a.Get("/properties/", a.GetProperties)
	a.Get("/property/{title}", a.GetProperty)
	a.Post("/property/", a.CreateProperty)
	a.Get("/assets/", a.GetAssets)
	a.Post("/asset/", a.GetAssets)
}

func (a *App) GetProperties(w http.ResponseWriter, r *http.Request) {
	handler.GetProperties(a.DB, w, r)
}

func (a *App) GetProperty(w http.ResponseWriter, r *http.Request) {
	handler.GetProperty(a.DB, w, r)
}

func (a *App) CreateProperty(w http.ResponseWriter, r *http.Request) {
	handler.CreateProperty(a.DB, w, r)
}

func (a *App) GetAssets(w http.ResponseWriter, r *http.Request) {
	handler.GetAssets(a.DB, w, r)
}

func (a *App) CreateAsset(w http.ResponseWriter, r *http.Request) {
	handler.CreateAsset(a.DB, w, r)
}

// wrapper for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// wrapper for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
