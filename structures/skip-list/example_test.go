package skip_list

import "fmt"

func ExampleString() {
	d := NewDictionary(0.75)
	var input = map[int]string{
		1: "aa",
		2: "ab",
		3: "ac",
		4: "ba",
		5: "bb",
		6: "bc",
		7: "ca",
		8: "cb",
		9: "cc",
	}

	for k, v := range input {
		d.Insert(k, v)
	}

	fmt.Printf("%v", d)

}
