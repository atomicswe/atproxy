package server

import "testing"

func TestNewServerConfigReturnsConfigWithDefaultPortAndAdress(t *testing.T) {
	sut := NewServerConfig()

	if sut.Port != defaultPort {
		t.Errorf("NewServerConfig returned a different port than the default: default = %d, got %d", defaultPort, sut.Port)
	}
	if sut.Address != defaultAdress {
		t.Errorf("NewServerConfig returned a different address than the default: default = '%s', got '%s'", defaultAdress, sut.Address)
	}
}
