package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig is the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache TemplateCache
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}

type TemplateCache map[string]*template.Template
