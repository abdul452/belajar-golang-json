package jsonmaptest

import (
	"encoding/json"
	"fmt"
	"testing"
)

// map digunakan untuk json yang isi data nya dinamis, karna kalau struct datanya harus di tentukan dulu
func TestMap(t *testing.T) {
	jsonString := `{"id":"p0001", "name":"Apple Mac Book Pro", "price":"20000000"}`
	jsonBytes := []byte(jsonString)

	var result map[string]any
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	produk := map[string]any{
		"id":    "p0001",
		"name":  "Apple Mac Book Pro",
		"price": "20000000",
	}

	bytes, _ := json.Marshal(produk)
	fmt.Println(string(bytes))
}
