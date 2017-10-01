// Unit tests for the identity package
package identity

import (
	"log"
	"testing"
)

// Test package initialization, which requires a database connection
func TestInit(t *testing.T) {
	log.Printf("TestInit: Begin unit tests\n")
}

// Test check login method
func TestCheckLogin1(t *testing.T) {
	user, err := CheckLogin("guest", "fgs0722")
	if err != nil {
		t.Error("TestCheckLogin1: error, ", err)
	}
	if len(user) != 1 {
		t.Error("TestCheckLogin1: len(user) != 1, ", len(user))
	}
}

// Test check login method
func TestCheckLogin2(t *testing.T) {
	user, err := CheckLogin("admin", "changeme")
	if err != nil {
		t.Error("TestCheckLogin2: error, ", err)
	}
	if len(user) != 0 {
		t.Error("TestCheckLogin2: len(user) != 0, ", len(user))
	}
}

// Test CheckSession function with expected result that session does not exist
func TestCheckSession1(t *testing.T) {
	sessionid := NewSessionId()
	session := CheckSession(sessionid)
	if session.Valid {
		t.Error("TestCheckSession1: session.Valid, sessionid: ",
			sessionid)
	}
}

// Test CheckSession function with session that does exist
func TestCheckSession2(t *testing.T) {
	sessionid := NewSessionId()
	userInfo := UserInfo{
		UserID: 1,
		UserName: "unittest",
		Email: "",
		FullName: "",
		Role: "",
	}
	SaveSession(sessionid, userInfo, 1)
	session := CheckSession(sessionid)
	if (session.Authenticated != 1) {
		t.Error("TestCheckSession2: session.Authenticated != 1, SessionID: ",
			sessionid)
	}
}

// Test CheckSession function with session that does exist
func TestCheckSession3(t *testing.T) {
	sessionid := NewSessionId()
	userInfo := UserInfo{
		UserID: 1,
		UserName: "guest",
		Email: "",
		FullName: "",
		Role: "",
	}
	SaveSession(sessionid, userInfo, 1)
	session := CheckSession(sessionid)
	if session.Authenticated != 1 {
		t.Error("TestCheckSession3: session.Authenticated != 1, SessionID: ",
			sessionid)
	}
}

// Test check login method
func TestNewSessionId(t *testing.T) {
	sessionid := NewSessionId()
	if sessionid == "invalid" {
		t.Error("TestNewSessionId: ", sessionid)
	}
}

// Test Logout method
func TestLogout(t *testing.T) {
	sessionid := NewSessionId()
	Logout(sessionid)
}

func TestSaveSession(t *testing.T) {
	sessionid := NewSessionId()
	userInfo := UserInfo{
		UserID: 1,
		UserName: "testuser",
		Email: "",
		FullName: "",
		Role: "",
	}
	SaveSession(sessionid, userInfo, 1)
}
