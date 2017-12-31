/*
 * Unit tests for library package
 */
package library

import (
	"testing"
)

func TestLoadLibrary0(t *testing.T) {
	emptyLibLoader := EmptyLibraryLoader{"Empty"}
	emptyLibLoader.LoadLibrary()
}

func TestLoadLibrary1(t *testing.T) {
	mockLoader := MockLibraryLoader{"Mock"}
	mockLoader.LoadLibrary()
}