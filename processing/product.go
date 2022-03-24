package processing

import (
	"fmt"
	"strconv"
	"strings"
)

//product is a struct that holds individual product information
type product struct {
	Name   string `json:"product"`
	Price  int    `json:"price"`
	Rating int    `json:"rating"`
}

//products is a struct on top of product to keep track of max values for price and rating
type products struct {
	maxPrice  int
	maxRating int
	P         []product
}

//addProduct is a helper func to add new values to P. It is needed to avoid duplicates
func (p *products) addProduct(pr product) {
	for _, v := range p.P {
		if v.Name == pr.Name {
			fmt.Printf(`Product with name: "%v" already exists, skipping duplicates\n`, pr.Name)
			return
		}
	}
	p.P = append(p.P, pr)
}

//UpdatePrice adds a new max price item to P or flushes P and adds a new item with new max price value
func (p *products) UpdatePrice(pr product) {
	if pr.Price > p.maxPrice {
		p.maxPrice = pr.Price
		p.P = nil
		p.P = append(p.P, pr)
		return
	}
	if pr.Price == p.maxPrice {
		p.addProduct(pr)
		return
	}
}

//UpdateRating has the same functionality as UpdatePrice but for rating
func (p *products) UpdateRating(pr product) {
	if pr.Rating > p.maxRating {
		p.maxRating = pr.Rating
		p.P = nil
		p.P = append(p.P, pr)
		return
	}
	if pr.Rating == p.maxRating {
		p.addProduct(pr)
		return
	}
}

//ProductList returns names of products that P holds
func (p products) ProductList() []string {
	var list []string
	for _, v := range p.P {
		list = append(list, v.Name)
	}
	return list
}

//newProductFromString parses csv string into product struct
func newProductFromString(s string) (product, error) {
	var p product
	split := strings.Split(s, ",")
	if len(split) != 3 {
		return p, fmt.Errorf("wrong csv format")
	}
	price, err := strconv.Atoi(split[1])
	if err != nil {
		return p, err
	}
	rating, err := strconv.Atoi(split[2])
	if err != nil {
		return p, err
	}
	p.Name = split[0]
	p.Price = price
	p.Rating = rating
	return p, nil
}
