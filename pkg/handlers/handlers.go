package handlers

import (
	"fmt"

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

//About hanler page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	StringMap := make(map[string]string)
	StringMap["test"] = "halo dari data map"

	remoteIP := m.app.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	fmt.Println("get ipnya :", remoteIP)

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: StringMap,
	})
}

//Home handler page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.app.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}
