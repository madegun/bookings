package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/madegun/bookings/models"
	"github.com/madegun/bookings/pkg/config"
	"github.com/madegun/bookings/pkg/render"

	"net/http"
)

//Repo for repository
var Repo *Repository

//Repository struct
type Repository struct {
	app *config.AppConfig
}

//NewRepo create the new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

//NewHandlers set the new handler for repository
func NewHandlers(r *Repository) {
	Repo = r
}

//Home handler page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})

}

//About hanler page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})

}

//Reservation hanler page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}

//Availability hanler page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

//PostAvailability hanler page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s, end date is %s", start, end)))

	// render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

type jsonRespon struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//PostAvailabilityJSON func
func (m *Repository) PostAvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	resp := jsonRespon{
		OK:      true,
		Message: "available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

//General hanler page
func (m *Repository) General(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "general.page.html", &models.TemplateData{})
}

//Major hanler page
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "major.page.html", &models.TemplateData{})
}

//Contact hanler page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}
