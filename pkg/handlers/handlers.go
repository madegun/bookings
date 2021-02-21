package handlers

import (
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

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}

//About hanler page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})

}

//Reservation hanler page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "make-reservation.page.html", &models.TemplateData{})
}

//Availability hanler page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "search-availability.page.html", &models.TemplateData{})
}

//PostAvailability hanler page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "search-availability.page.html", &models.TemplateData{})
}

//General hanler page
func (m *Repository) General(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "general.page.html", &models.TemplateData{})
}

//Major hanler page
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "major.page.html", &models.TemplateData{})
}

//Contact hanler page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}
