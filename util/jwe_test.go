package util

import (
	"reflect"
	"testing"

	jwe "github.com/square/go-jose"
)

func TestCreateGetJWE(t *testing.T) {
	originalClaim := NewClaim("4b020797-25a8-4163-a693-3d6f9657d54d", "Henri", 0)
	jweValue, err := CreateJWE(originalClaim)
	if err != nil {
		t.Errorf("create jwe: %v", err)
		return
	}
	t.Logf("success jwe: %v\n", jweValue)

	retrivedClaim, err := GetJWE(jweValue)
	if err != nil {
		t.Errorf("create jwe: %v", err)
		return
	}

	if !reflect.DeepEqual(originalClaim, retrivedClaim) {
		t.Errorf("create jwe: %v", err)
		return
	}
}

func TestGetJWE(t *testing.T) {
	_, err := GetJWE("eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkExMjhHQ00ifQ.NW4iGe6ZMC_lHJb-CkbKko0NrLM84cMIdxoF15nIw6iwKTxgCLygF2KKCUuzw4S_JAz0rDtJkBBZwZ0g8QiGTfxFNivhyMy7fBO9yXO4QgbJZAxht7C7oOPmbVMAiPMB0xcah7lBZF9hbWsxasoqkbmUJ3rAbmg79kpFdNQbnH_QzdUvOWBI6NLLPnThUH8kjmIHpsxLpnt-i8tc0DBnQ4Vd9ixy05sFqLy7mrGKFYZVeo8H0tSkgTJ84PmPm7bHBgLgOYzw3wajLHFJKkc1qA10efwHKRSK_XaLOr_2dFb_6p7_VEmRqI-x_dqDQfwFh8AdyQvXspMrp5NMdNmf4A.NBkg1l3p_8uUI2gT.u9xKs6HSvuO8M_plVGYnDc4QrVkCMOXNnxqMBI6Y5egAj610e58AQuMFA96pdG7nprbQZvk9NspKrMBtoteoKUWlaouYbFoNb1EtMdQAaICRDx_Jdc4uM6l8IIBD3Z_6QhUDTvUpWGA.CMogyFVrSlZiRHYlmr2pGw")
	if err != jwe.ErrCryptoFailure {
		t.Errorf("error: %v", err)
		return
	}
	t.Logf("success error %v",err)
}
