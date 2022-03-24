package test_processing

import (
	"os"
	"testing"

	"github.com/kebmo91/omp-test-task/processing"
)

func TestNewTypeProcessor(t *testing.T) {
	var tests = []struct {
		t   string
		exp bool
	}{
		{"./testdata/newtypeprocessor/test1.csv", true},
		{"./testdata/newtypeprocessor/test2.csv", false},
		{"./testdata/newtypeprocessor/test3.csv", false},
		{"./testdata/newtypeprocessor/test4.csv", false},
		{"./testdata/newtypeprocessor/test5.json", true},
		{"./testdata/newtypeprocessor/test6.json", false},
		{"./testdata/newtypeprocessor/test7.json", false},
		{"./testdata/newtypeprocessor/test8.json", false},
	}
	for _, v := range tests {
		var got bool
		f, err := os.Open(v.t)
		defer f.Close()
		if err != nil {
			t.Error(err)
		}
		_, err = processing.NewTypeProcessor(f)
		if err != nil {
			got = false
		} else {
			got = true
		}
		if got != v.exp {
			t.Errorf(`Something went wrong with test %v`, v.t)
		}
	}
}

func TestProcess(t *testing.T) {
	var tests = []struct {
		t         string
		expPrice  []string
		expRating []string
	}{
		{"./testdata/process/test1.csv", []string{"Cookies", "Onion"}, []string{"Cookies", "Onion"}},
		{"./testdata/process/test2.csv", []string{"Cheese", "Garlic", "Carrots"}, []string{"Garlic", "Carrots"}},
		{"./testdata/process/test3.csv", []string{"Beef"}, []string{"Beef", "Cookies", "Onion", "Pork"}},
		{"./testdata/process/test4.csv", []string{"Cookies", "Onion", "Fish"}, []string{"Cookies", "Onion", "Fish"}},
		{"./testdata/process/test5.json", []string{"Cheese", "Fish"}, []string{"Cheese", "Fish", "Jam"}},
		{"./testdata/process/test6.json", []string{"Whiskey"}, []string{"Cognac"}},
		{"./testdata/process/test7.json", []string{"Milk"}, []string{"Milk"}},
		{"./testdata/process/test8.json", []string{"Jam", "Cheese", "Fish"}, []string{"Cheese", "Fish"}},
	}
	for _, v := range tests {
		f, err := os.Open(v.t)
		defer f.Close()
		if err != nil {
			t.Error(err)
		}
		tp, err := processing.NewTypeProcessor(f)
		if err != nil {
			t.Error(err)
		}
		maxPrice, maxRating, err := tp.Process()
		if err != nil {
			t.Error(err)
		}
		if !sameStringSlice(maxPrice, v.expPrice) {
			t.Errorf(`unexpected result for price in %v`, v.t)
		}
		if !sameStringSlice(maxRating, v.expRating) {
			t.Errorf(`unexpected result for rating in %v`, v.t)
		}
	}
}

func sameStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	amap := make(map[string]int)
	bmap := make(map[string]int)
	for _, s := range a {
		amap[s]++
	}
	for _, s := range b {
		bmap[s]++
	}
	for aMapKey, aMapV := range amap {
		if bmap[aMapKey] != aMapV {
			return false
		}
	}
	return true
}
