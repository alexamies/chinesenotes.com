// Unit tests for the config package
package config

import (
	"log"
	"testing"
)

// Test reading of files for HTML conversion
func TestAvoidSubDomains(t *testing.T) {
	log.Printf("AvoidSubDomains: Begin unit tests\n")
	avoidSub := AvoidSubDomains()
	expect := 0
	if len(avoidSub) != expect {
		t.Errorf("TestAvoidSubDomains: Got %d, expcted %d", len(avoidSub), expect)
	}
}

// Test reading of files for HTML conversion
func TestGenres(t *testing.T) {
	vars := readConfig()
	log.Printf("TestReadConfig: len(vars): %d\n", len(vars))
	if len(vars) == 0 {
		t.Error("TestReadConfig: No vars found")
	}
}

// Test reading of files for HTML conversion
func TestReadConfig(t *testing.T) {
	vars := readConfig()
	if len(vars) == 0 {
		t.Error("TestReadConfig: No vars found")
	}
}


// Test reading of lexical unit file names
func TestLUFileNames(t *testing.T) {
	luFiles := LUFileNames()
	log.Printf("TestLUFileNames: len(luFiles): %d\n", len(luFiles))
	if len(luFiles) == 0 {
		t.Error("TestLUFileNames: No lexical unit files found")
	}
}

// Test reading of files for HTML conversion
func TestGetHTMLConversions(t *testing.T) {
	log.Printf("TestGetHTMLConversions: Begin unit tests\n")
	conversions := GetHTMLConversions()
	log.Printf("TestGetHTMLConversions: # conversions: %d\n", len(conversions))
}
