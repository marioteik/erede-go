package eredego

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const sandbox = "https://sandbox-erede.useredecloud.com.br"
const production = "https://api.userede.com.br/erede"

// AuthorizationRequest is the struct that contains the request body
// more documentation at https://developer.userede.com.br/e-rede#endpoint
type AuthorizationRequest struct {
	Capture                bool   `json:"capture"`
	Kind                   string `json:"kind" validate:"oneof=credit debit"`
	Reference              string `json:"reference" validate:"required,max=16"`
	Amount                 int64  `json:"amount" validate:"required,max=9999999999,numeric"`
	Installments           int8   `json:"installments" validate:"min=2,max=12,numeric"`
	CardholderName         string `json:"kicardholderNamend" validate:"max=30"`
	CardNumber             string `json:"cardNumber" validate:"required,max=19"`
	ExpirationMonth        int8   `json:"expirationMonth" validate:"required,max=2,numeric"`
	ExpirationYear         int16  `json:"expirationYear" validate:"required,max=4,numeric"`
	SecurityCode           string `json:"securityCode" validate:"max=4"`
	SoftDescriptor         string `json:"softDescriptor" validate:"max=18"`
	Subscription           bool   `json:"subscription"`
	Origin                 int8   `json:"origin" validate:"max=2,numeric"`
	DistributorAffiliation int32  `json:"distributorAffiliation"  validate:"max=9,numeric"`
	BrandTid               string `json:"brandTid"  validate:"max=16"`
	StorageCard            string `json:"storageCard"  validate:"max=2,numeric"`
	TransactionCredentials TransactionCredentials
}

// TransactionCredentials is the struct that contains the credentials for the transaction
type TransactionCredentials struct {
	CredentialId string `json:"credentialId" validate:"required,max=2"`
}

// AuthorizationResponse is the struct that contains the response body
type AuthorizationResponse struct {
	ReturnCode        string `json:"returnCode"`
	ReturnMessage     string `json:"returnMessage"`
	Reference         string `json:"reference"`
	Tid               string `json:"tid"`
	Nsu               string `json:"nsu"`
	AuthorizationCode string `json:"authorizationCode"`
	DateTime          string `json:"dateTime"`
	Amount            int64  `json:"amount"`
	CardBin           string `json:"cardBin"`
	Last4             string `json:"last4"`
	Brand             Brand  `json:"brand"`
}

// Brand is the struct that contains the brand information
type Brand struct {
	Name               string `json:"name"`
	ReturnCode         string `json:"returnCode"`
	ReturnMessage      string `json:"returnMessage"`
	MerchantAdviceCode string `json:"merchantAdviceCode"`
	AuthorizationCode  string `json:"authorizationCode"`
	BrandTid           string `json:"brandTid"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// GetAuthorization makes the request to the API to get the authorization for a transaction
func GetAuthorization(authorization AuthorizationRequest) (AuthorizationResponse, error) {
	const endpoint = "/v1/transactions"

	url := sandbox + endpoint

	authorizationJosn, err := json.Marshal(authorization)

	if err != nil {
		return AuthorizationResponse{}, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(authorizationJosn))

	if err != nil {
		return AuthorizationResponse{}, err
	}

	req.Header.Add("Authorization", "Basic "+basicAuth("72548122", "737df67d33374596b7f199e8329d8a39"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return AuthorizationResponse{}, err
	}

	resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return AuthorizationResponse{}, err
	}

	var result AuthorizationResponse
	err = json.Unmarshal([]byte(body), &result)

	if err != nil {
		return AuthorizationResponse{}, err
	}

	if result.ReturnCode != "00" {
		return AuthorizationResponse{}, errors.New("ReturnCode: " + result.ReturnCode + "; " + result.ReturnMessage)
	}

	return result, nil
}
