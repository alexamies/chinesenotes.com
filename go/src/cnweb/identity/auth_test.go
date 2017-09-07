// Unit tests for the config package
package identity

import (
	"log"
	"testing"
)

// Test package initialization, which requires a database connection
func TestInit(t *testing.T) {
	log.Printf("TestCheckLogin: Begin unit tests\n")
}

// Test check login method
func TestCheckLogin1(t *testing.T) {
	user := CheckLogin("guest", "guest")
	if len(user) != 1 {
		t.Error("TestCheckLogin1: len(user) != 1, ", len(user))
	}
}

// Test check login method
func TestCheckLogin2(t *testing.T) {
	user := CheckLogin("admin", "changeme")
	if len(user) != 0 {
		t.Error("TestCheckLogin2: len(user) != 0, ", len(user))
	}
}

// Test check login method
func TestNewSessionId(t *testing.T) {
	sessionid := NewSessionId()
	if sessionid == "invalid" {
		t.Error("TestNewSessionId: ", sessionid)
	}
}

// Test check login method
func TestGetSiteDomain(t *testing.T) {
	domain := GetSiteDomain()
	if domain != "localhost" {
		t.Error("TestGetSiteDomain: domain = ", domain)
	}
}

// Test Logout method
func TestLogout(t *testing.T) {
	sessionid := NewSessionId()
	Logout(sessionid)
}

// Test check login method
func TestSaveSession(t *testing.T) {
	sessionid := NewSessionId()
	SaveSession(sessionid, "testuser")
}
