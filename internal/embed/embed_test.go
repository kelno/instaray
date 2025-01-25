package embed

import (
	"testing"
)

func TestCheck(t *testing.T) {
	embed := New("platform", "newplatform.com")

	tests := []struct {
		input    string
		expected bool
	}{
		// Valid cases
		{"platform.com", true},
		{"www.platform.com", true},
		{"http://platform.com", true},
		{"https://www.platform.com", true},
		{"https://vm.platform.com", true},
		{"https://www.platform.com/ZMYnv3JGV/", true},
		{"http://platform.com/user123", true},
		{"https://platform.com/someone_else", true},
		{"https://www.platform.com/", true},
		{"platform.com/test", true},
		// Invalid cases
		{"https://notplatform.com/meh", false},
		{"ftp://platform.com/test", false},
		{"platform.co/user", false},
		{"this is not an url", false},
	}

	for _, test := range tests {
		got := embed.Check(test.input)
		if got != test.expected {
			t.Errorf("For input %q to Check(), expected %v, but got %v", test.input, test.expected, got)
		}
	}
}

func TestReplace(t *testing.T) {
	embed := New("platform", "newplatform.com")

	tests := []struct {
		input    string
		expected string
	}{
		// Valid cases
		{"platform.com", "newplatform.com"},
		{"www.platform.com", "www.newplatform.com"},
		{"http://platform.com", "http://newplatform.com"},
		{"https://www.platform.com", "https://www.newplatform.com"},
		{"https://vm.platform.com", "https://vm.newplatform.com"},
		{"https://www.platform.com/ZMYnv3JGV/", "https://www.newplatform.com/ZMYnv3JGV/"},
		{"http://platform.com/user123", "http://newplatform.com/user123"},
		{"https://platform.com/someone_else", "https://newplatform.com/someone_else"},
		{"https://www.platform.com/", "https://www.newplatform.com/"},
		{"platform.com/test", "newplatform.com/test"},
		// Invalid cases
		{"https://notplatform.com/meh", "https://notplatform.com/meh"},
		{"ftp://platform.com/test", "ftp://platform.com/test"},
		{"platform.co/user", "platform.co/user"},
		{"this is not an url", "this is not an url"},
	}

	for _, test := range tests {
		got := embed.Replace(test.input)
		if got != test.expected {
			t.Errorf("For input %q to Replace(), expected %q, but got %v", test.input, test.expected, got)
		}
	}
}
