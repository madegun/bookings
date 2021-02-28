package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/madegun/bookings/internal/config"
	"github.com/madegun/bookings/internal/forms"
	"github.com/madegun/bookings/internal/render"
	"github.com/madegun/bookings/models"

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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

//PostReservation handler to post a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{

		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("last_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.app.Session.Put(r.Context(), "reservation", reservation)

	//redirect to page /reservation-summary
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
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

//ReservationSummary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {

	reservation, ok := m.app.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("tidak ada reservation data session")
		m.app.Session.Put(r.Context(), "error", "tidak ada reservation session data!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.app.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	})
}
