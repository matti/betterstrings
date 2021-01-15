package betterstrings

import "strings"

// SliceFunc ...
type SliceFunc struct {
	elements map[int]string
}

// One ...
func (a *SliceFunc) One() string {
	return a.elements[0]
}

// Two ...
func (a *SliceFunc) Two() (string, string) {
	return a.elements[0], a.elements[1]
}

// Three ...
func (a *SliceFunc) Three() (string, string, string) {
	return a.elements[0], a.elements[1], a.elements[2]
}

// Four ...
func (a *SliceFunc) Four() (string, string, string, string) {
	return a.elements[0], a.elements[1], a.elements[2], a.elements[3]
}

// Five ...
func (a *SliceFunc) Five() (string, string, string, string, string) {
	return a.elements[0], a.elements[1], a.elements[2], a.elements[3], a.elements[4]
}

// Split ...
func Split(s string, token string) *SliceFunc {
	sf := &SliceFunc{
		elements: make(map[int]string),
	}

	for i, part := range strings.Split(s, token) {
		sf.elements[i] = part
	}
	return sf
}
