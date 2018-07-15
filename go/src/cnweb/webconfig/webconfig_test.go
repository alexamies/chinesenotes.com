// Unit tests for the identity package
package webconfig

import (
	"log"
	"testing"
)

// Test package initialization, which requires a database connection
func TestInit(t *testing.T) {
	log.Printf("TestInit: Begin unit tests\n")
}

// Test default serving port
func TestGetPort(t *testing.T) {
	port := GetPort()
	if port != 8080 {
		t.Error("TestGetPort: port = ", port)
	}
}

// Test site domain
func TestGetSiteDomain(t *testing.T) {
	domain := GetSiteDomain()
	if domain != "localhost" {
		t.Error("TestGetSiteDomain: domain = ", domain)
	}
}
