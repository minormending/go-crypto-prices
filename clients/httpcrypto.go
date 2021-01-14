package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTPCrypto talks to a sever an returns the response
type HTTPCrypto struct{}

// Get retrieves data at the url and converts it to the json obj
func (h *HTTPCrypto) Get(url string, j interface{}) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to get price from %s : %v", url, err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&j); err != nil {
		return nil, fmt.Errorf("unable to parse response from %s : %s", url, err)
	}
	return j, nil
}
