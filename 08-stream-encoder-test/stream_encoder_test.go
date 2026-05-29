package streamencodertest

import (
	"encoding/json"
	"os"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	// 🟢 Wajib tambahkan defer untuk memastikan file otomatis ditutup saat fungsi selesai
	defer writer.Close()
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
	}

	encoder.Encode(customer)
}

func TestEncoderMap(t *testing.T) {
	writer, _ := os.Create("CustomerOutMap.json")
	// 🟢 Wajib tambahkan defer untuk memastikan file otomatis ditutup saat fungsi selesai
	defer writer.Close()
	encoder := json.NewEncoder(writer)

	customer := map[string]any{
		"FirstName":  "Eko",
		"MiddleName": "Kurniawan",
		"LastName":   "Khannedy",
	}

	encoder.Encode(customer)
}
