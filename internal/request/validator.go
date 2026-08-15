package request

import (
	"net/http"
	"slices"
)

type ValidatorConfig struct {
	AllowedDomains []string `json:"allowed_domains"`
}

type Validator struct {
	config ValidatorConfig
}

func NewValidator(config ValidatorConfig) *Validator {
	return &Validator{
		config: config,
	}
}

func (v *Validator) Allowed(r *http.Request) bool {
	return slices.Contains(v.config.AllowedDomains, r.URL.Hostname())
}
