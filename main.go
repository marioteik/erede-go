package eredego

import "math/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";

// ERedeGo is the main struct of the package
type ERedeGo struct {}

// GenerateRandomString generates a random string with the length passed as parameter
func (e ERedeGo) GenerateRandomString(length int) string {
	var result string

	for i := 0; i < length; i++ {
		result += string(randomStringSource[rand.Intn(len(randomStringSource))])
	}
	
	return result
}