package processing

import (
	"encoding/json"
	"os"
)

//JSONProcessor is a struct for processing json data.
//It implemets fileProcessor interface
type JSONProcessor struct {
	f *os.File
}

//Process is a func that processes incoming json file and returns products with max price
//and max rating. It processes products object by object in a json array so that program
//does not run out of memory
func (js JSONProcessor) Process() ([]string, []string, error) {
	dec := json.NewDecoder(js.f)
	_, err := dec.Token()
	if err != nil {
		return []string{}, []string{}, err
	}
	var maxPriceProducts products
	var maxRatingProducts products
	for dec.More() {
		var p product
		err := dec.Decode(&p)
		if err != nil {
			return []string{}, []string{}, err
		}
		maxPriceProducts.UpdatePrice(p)
		maxRatingProducts.UpdateRating(p)
	}
	return maxPriceProducts.ProductList(), maxRatingProducts.ProductList(), nil
}
