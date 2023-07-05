package eredego

import "testing"

func TestERedeGo_GenerateRandomString(t *testing.T) {
	var eredego ERedeGo

	result := eredego.GetAuthorization(16)

	if len(result) != 16 {
		t.Errorf("Expected string length of 16, got %d", len(result))
	}
}
