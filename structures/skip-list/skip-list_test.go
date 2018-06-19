package skip_list

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func PrintFatal(t *testing.T, expected fmt.Stringer, got fmt.Stringer) {
	t.Fatalf("Expected: \n%v, but got: \n%v\n", expected, got)
}

func RandStringBytesMaskImprSrc(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func TestInitializer(t *testing.T) {

	d := NewDictionary(0.5)
	node := d.Lookup(5)
	if node != nil {
		t.Fatalf("Searching for non-present key, should return nil, returned: %v\n", node)
	}
}

func TestInsert(t *testing.T) {

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

	var no_input = map[int]string{
		11: "aa",
		12: "ab",
		13: "ac",
		14: "ba",
		15: "bb",
		16: "bc",
		17: "ca",
		18: "cb",
		19: "cc",
	}

	d := NewDictionary(0.5)

	for k, v := range input {
		d.Insert(k, v)

	}

	for k, v := range input {
		node := d.Lookup(k)
		if node.value != v || node.key != k {
			t.Fatalf("Expected key-value %v-%v, Got %v-%v\n", k, v, node.key, node.value)
		}
	}

	for k, _ := range no_input {
		node := d.Lookup(k)
		if node != nil {
			t.Fatalf("Found a key that was not inserted: %v\n", node)
		}
	}

	t.Logf("%v", d)

}

func TestInsertDelete(t *testing.T) {

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

	var no_input = map[int]string{
		11: "aa",
		12: "ab",
		13: "ac",
		14: "ba",
		15: "bb",
		16: "bc",
		17: "ca",
		18: "cb",
		19: "cc",
	}

	d := NewDictionary(0.5)

	for k, v := range input {
		d.Insert(k, v)
	}

	for k, v := range input {
		removed := d.Delete(k)

		if removed != true {
			t.Fatalf("Expected that element %v-%v was removed (instead it was not found)n", k, v)
		}
	}

	for k, v := range no_input {
		removed := d.Delete(k)

		if removed != false {
			t.Fatalf("Element %v-%v was not present and delete reports that it was correctly removed \n", k, v)
		}
	}

	for k, v := range input {
		d.Insert(k, v)

	}

	for k, v := range input {
		node := d.Lookup(k)
		if node.value != v || node.key != k {
			t.Fatalf("Expected key-value %v-%v, Got %v-%v\n", k, v, node.key, node.value)
		}
	}

	for k, _ := range no_input {
		node := d.Lookup(k)
		if node != nil {
			t.Fatalf("Found a key that was not inserted: %v\n", node)
		}
	}

}

func TestRepeatedInsert(t *testing.T) {

	var input = []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"}

	d := NewDictionary(0.5)

	for _, v := range input {
		d.Insert(1, v)
		node := d.Lookup(1)
		if node.value != v || node.key != 1 {
			t.Fatalf("Expected key-value %v-%v, Got %v-%v\n", 1, v, node.key, node.value)
		}

	}

	node := d.Lookup(1)
	if node.value != "cc" || node.key != 1 {
		t.Fatalf("Expected key-value %v-%v, Got %v-%v\n", 1, "cc", node.key, node.value)
	}
}

func TestRandomInsert(t *testing.T) {

	d := NewDictionary(0.5)

	for i := 0; i < 1000; i++ {
		k, v := rand.Int(), RandStringBytesMaskImprSrc(5)
		d.Insert(k, v)
		node := d.Lookup(k)
		if node.value != v || node.key != k {
			t.Fatalf("Expected key-value %v-%v, Got %v-%v\n", k, v, node.key, node.value)
		}

	}

}
