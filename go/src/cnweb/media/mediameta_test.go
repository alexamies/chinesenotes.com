// Unit tests for query parsing functions
package media

import (
	"log"
	"testing"
)

// Test trivial query
func TestFindMedia(t *testing.T) {
	log.Printf("TestFindMedia: Begin unit tests\n")
	metadata, err := FindMedia("hello")
	if err != nil {
		t.Error("TestFindMedia: encountered error: ", err)
		return
	}
	log.Printf("TestFindMedia: metadata.ObjectId", metadata.ObjectId)
}
