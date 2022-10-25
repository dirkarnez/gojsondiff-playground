package main

import (
	"encoding/json"
	"fmt"

	"github.com/yudai/gojsondiff"
)

func main() {
	s := gojsondiff.New()
	a := []string{"1", "2"}
	b := []string{"1", "3", "4"}

	aa, _ := json.Marshal(&a)
	bb, _ := json.Marshal(&b)
	var aaa []interface{}
	_ = json.Unmarshal(aa, &aaa)

	var bbb []interface{}
	_ = json.Unmarshal(bb, &bbb)
	d := s.CompareArrays(aaa, bbb)

	fmt.Printf("%+v", d.Deltas()) //prints: [2, 1]
}