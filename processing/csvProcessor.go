package processing

import (
	"bufio"
	"os"
)

//CSVProcessor is a struct for processing csv data.
//It implemets fileProcessor interface
type CSVProcessor struct {
	f *os.File
}

//Process is a func that processes incoming csv file and returns products with max price
//and max rating. It scans products line by line so that program does not run out of memory
func (c CSVProcessor) Process() ([]string, []string, error) {
	scanner := bufio.NewScanner(c.f)
	scanner.Scan()
	var maxRatingProducts products
	var maxPriceProducts products
	for scanner.Scan() {
		text := scanner.Text()
		p, err := newProductFromString(text)
		if err != nil {
			return []string{}, []string{}, err
		}
		maxPriceProducts.UpdatePrice(p)
		maxRatingProducts.UpdateRating(p)
	}
	return maxPriceProducts.ProductList(), maxRatingProducts.ProductList(), nil
}
