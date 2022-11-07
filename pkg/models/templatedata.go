package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string //security token 'Cross Site Request Forgery Token' for forms
	Flash     string //'flash message' posted to used (such as "success!")
	Warning   string
	Error     string
}
