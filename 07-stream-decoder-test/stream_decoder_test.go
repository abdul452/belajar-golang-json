package streamdecodertest

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestStreamDecoder(t *testing.T) {
	reader, err := os.Open("customer.json")
	if err != nil {
		panic(err)
	}
	// 🟢 WAJIB ditaruh di sini untuk close file setelah dibaca
	defer reader.Close()
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)
	fmt.Println(customer)
}

func TestStreamDecoderMap(t *testing.T) {
	reader, err := os.Open("customer.json")
	if err != nil {
		panic(err)
	}
	// 🟢 WAJIB ditaruh di sini untuk close file setelah dibaca
	defer reader.Close()
	decoder := json.NewDecoder(reader)

	var customer = map[string]any{}
	err = decoder.Decode(&customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
