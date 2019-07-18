package gobool_test

import (
	"github.com/hypnoglow/gobool"
	"testing"
)

func TestIsTrue(t *testing.T) {
	if !gobool.IsTrue(true) {
		t.Fatalf("Expected true")
	}

	if gobool.IsTrue(false) {
		t.Fatalf("Expected false")
	}
}

func TestIsFalse(t *testing.T) {
	if gobool.IsFalse(true) {
		t.Fatalf("Expected false")
	}

	if !gobool.IsFalse(false) {
		t.Fatalf("Expected true")
	}
}
