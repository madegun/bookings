package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//AppConfig hold config in application
type AppConfig struct {
	UseChace      bool
	TemplateChace map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
