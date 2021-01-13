package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTPCrypto talks to a sever an returns the response
type HTTPCrypto struct{}

// Get retrieves data at the url and converts it to the json obj
func (h *HTTPCrypto) Get(url string, j interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to get price from %s : %v", url, err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&j); err != nil {
		return fmt.Errorf("unable to parse response from %s : %s", url, err)
	}
	return nil
}
