package model

type TemplateFile struct {
	Content string `json:"content"`
}

type Template struct {
	ID            string                   `json:"id"`
	Name          string                   `json:"name"`
	DefaultName   string                   `json:"default_name"`
	Category      string                   `json:"category"`
	TemplateFiles map[string]*TemplateFile `json:"template_files"`
	StateSettings *StateSettings           `json:"state_settings"`
}
