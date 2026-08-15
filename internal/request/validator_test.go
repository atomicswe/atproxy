package request

import (
	"net/http"
	"net/url"
	"testing"
)

func TestNewValidatorConfigReturnsEmptyAllowedDomains(t *testing.T) {
	sut := NewValidatorConfig()
	if len(sut.AllowedDomains) > 0 {
		t.Error("NewValidatorConfig should create an empty list for AllowedDomains")
	}
}

func TestAllowedAllowsRequestWithAllowedDomain(t *testing.T) {
	config := NewValidatorConfig()
	config.AllowedDomains = append(config.AllowedDomains, "validurl.test")
	sut := NewValidator(config)

	r := &http.Request{
		URL: &url.URL{
			Host: "validurl.test:9999",
		},
	}

	if !sut.Allowed(r) {
		t.Errorf("a request from an allowed domain should be allowed")
	}
}

func TestAllowedDoesNotAllowRequestWithoutAllowedDomain(t *testing.T) {
	config := NewValidatorConfig()
	sut := NewValidator(config)

	r := &http.Request{
		URL: &url.URL{
			Host: "validurl.test:9999",
		},
	}

	if sut.Allowed(r) {
		t.Errorf("a request from a disallowed domain shouldn't be allowed")
	}
}
