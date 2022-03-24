package processing

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type fileProcessor interface {
	Process() ([]string, []string, error)
}

//TypeProcessor is a struct that is meant to determine what type of file is going to be
//processed: csv or json. It has a fileProcessor interface field where a csv or json
//processor is instantiated after NewTypeProcessor call
type TypeProcessor struct {
	fileProcessor
}

//NewTypeProcessor creates a TypeProcessor object and instantiates it with
//a suitable processor json or csv
func NewTypeProcessor(f *os.File) (TypeProcessor, error) {
	var tp TypeProcessor
	if tp.isValidCSV(f) {
		var csv CSVProcessor
		csv.f = f
		tp.fileProcessor = csv
		return tp, nil
	}
	if tp.isValidJSON(f) {
		var js JSONProcessor
		js.f = f
		tp.fileProcessor = js
		return tp, nil
	}
	return tp, fmt.Errorf("file format is neither a valid json nor csv")
}

//isValidJSON checks first 10 json tokens. If they are valid, returns true.
//means it's indeed a json file with at least one object inside
func (f TypeProcessor) isValidJSON(file *os.File) bool {
	js := json.NewDecoder(file)
	for i := 0; i < 10; i++ {
		js.More()
		_, err := js.Token()
		if err != nil {
			file.Seek(0, io.SeekStart)
			return false
		}
	}
	file.Seek(0, io.SeekStart)
	return true
}

//isValidCSV checks first two lines in a file. If those are two valid csv lines,
//returns true. if not returns false
func (f TypeProcessor) isValidCSV(file *os.File) bool {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	split := strings.Split(text, ",")
	if len(split) != 3 {
		file.Seek(0, io.SeekStart)
		return false
	}
	if split[0] != "Product" || split[1] != "Price" || split[2] != "Rating" {
		file.Seek(0, io.SeekStart)
		return false
	}
	scanner.Scan()
	text = scanner.Text()
	split = strings.Split(text, ",")
	if len(split) != 3 {
		file.Seek(0, io.SeekStart)
		return false
	}
	file.Seek(0, io.SeekStart)
	return true
}
