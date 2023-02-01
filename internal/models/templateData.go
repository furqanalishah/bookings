package models

import "github.com/furqanalishah/bookings/internal/forms"

// TemplateData hold data sent to the templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
