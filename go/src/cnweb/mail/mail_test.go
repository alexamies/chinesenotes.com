// Unit tests for the mail package
package mail

import (
	"log"
	"testing"
	"cnweb/identity"
)

// Test package initialization, which requires a database connection
func TestSendPasswordReset(t *testing.T) {
	log.Printf("TestSendPasswordReset: Begin unit tests\n")
	userInfo := identity.UserInfo{
		UserID: 100,
		UserName: "test",
		Email: "alex@chinesenotes.com",
		FullName: "Alex Test",
		Role: "tester",
	}
	err := SendPasswordReset(userInfo)
	if err != nil {
		log.Println("TestSendPasswordReset: Error, ", err)
	}
}
