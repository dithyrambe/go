package util

import "testing"

type EmailTest struct {
	value string
	err   error
}

func TestValidEmail(t *testing.T) {
	emails := []EmailTest{
		{"m.dupond@gmail.com", nil},
		{"m.dupond@asldfjlasdjfs.com", ErrEmailDomainNotValid},
		{"m.dupond@asl@gmail.com", ErrEmailNotValid},
		{"m.dupond@yopmail.fr", ErrEmailBlackListed},
	}
	for _, e := range emails {
		err := ValidEmail(e.value)
		if err != e.err {
			t.Errorf("waiting for %v got %v", e.err, err)
		}
	}
}
