package rick_morty

import (
	"testing"
)

func TestGetCaracters(t *testing.T) {
	caracters, err := GetCaracters()
	if err != nil {
		t.Error(err)
	}
	if caracters.Results[0].Name != "Toxic Rick"{
		t.Error("First caractere should be Toxic Rick")
	}
}
