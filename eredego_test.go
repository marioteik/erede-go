package eredego

import (
	"testing"
)

func TestERedeGo_GenerateRandomString(t *testing.T) {
	result, err := GetAuthorization(AuthorizationRequest{
		Capture:                true,
		Kind:                   "credit",
		Reference:              "1234567890123458",
		Amount:                 100,
		CardholderName:         "John Doe",
		CardNumber:             "5448280000000007",
		ExpirationMonth:        12,
		ExpirationYear:         2026,
		SecurityCode:           "123",
		SoftDescriptor:         "SoftDescriptor",
		Subscription:           false,
		Origin:                 1,
		DistributorAffiliation: 0,
		BrandTid:               "1234567890123456",
		StorageCard:            "0",
		TransactionCredentials: TransactionCredentials{
			CredentialId: "01",
		},
	})

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if result.ReturnCode != "00" {
		t.Errorf("Expected result, got and trasaction error instead: %s", result.ReturnMessage)
	}
}
