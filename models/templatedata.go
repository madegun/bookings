package models

//TemplateData untuk parsing type data
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CRSFToken string
	Flash     string
	Warning   string
	Error     string
}
