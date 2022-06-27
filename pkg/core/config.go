package core

type Config struct {
	Tags map[string]string `json:"tags,omitempty" yaml:"tags"`
	Title string `json:"title,omitempty" yaml:"title"`
	Description string `json:"description,omitempty" yaml:"description"`
	Version string `json:"version,omitempty" yaml:"version"`
	ContactName  string  `json:"contactName,omitempty" yaml:"contactName"`
	ContactURL   string  `json:"contactUrl,omitempty" yaml:"contactUrl"`
	ContactEmail string `json:"contactEmail,omitempty" yaml:"contactEmail"`
}
