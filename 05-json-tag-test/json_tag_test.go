package jsontagtest

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	// ImageURL string `json:"image_url,omitempty"` // 🚀 Hilang dari JSON jika string-nya kosong ""
	// NOTE: Gunakan json:"-" untuk menyembunyikan field sensitif
	// seperti password, flag internal backend, atau object database (*sql.DB).
	// SecretCode string `json:"-"` // 🔒 Tidak akan pernah ikut di-Marshal maupun di-Unmarshal
	// Password string `json:"-"` // 🔒 Disembunyikan!
	// IsProcessed bool   `json:"-"` // 🔒 Flag internal backend
	// DB      *sql.DB `json:"-"` // 🔒 Objek koneksi database tidak bisa/boleh di-JSON-kan
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "0001",
		Name:     "Apple Mac Book Pro",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"0001","name":"Apple Mac Book Pro","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
