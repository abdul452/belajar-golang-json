package decodejsontest

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestDecodeJSON(t *testing.T) {
	// Membuat mentahan string JSON.
	jsonString := `{"FirstName":"Eko","MiddleName":"Kurniawan","LastName":"Khannedy","Age":30,"Married":true}`
	// Melakukan konversi string tersebut menjadi []byte
	jsonBytes := []byte(jsonString)

	// Menyiapkan variabel penampung berupa pointer struct
	customer := &Customer{} // 🎯 Ini adalah pointer ke struct Customer

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
