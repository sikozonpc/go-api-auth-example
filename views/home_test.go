package views

import (
	"strings"
	"testing"

	"github.com/markbates/goth"
)

func TestHomePage(t *testing.T) {
	comp, err := componentToString(Home(goth.User{Name: "John Doe"}))
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(comp, "Car Show Example App") {
		t.Errorf("Expected Car Show Example App', got '%s'", comp)
	}

	if !strings.Contains(comp, "John Doe") {
		t.Errorf("Expected 'John Doe', got '%s'", comp)
	}
}
