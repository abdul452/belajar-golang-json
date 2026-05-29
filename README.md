# Belajar Golang JSON

Repositori ini berisi kumpulan materi, catatan, dan unit test praktis saat mempelajari pengelolaan data format JSON menggunakan bahasa pemrograman Go (Golang). Seluruh modul dirancang menggunakan pendekatan *test-driven* untuk memahami mekanisme *marshalling*, *unmarshalling*, kustomisasi tag, hingga manajemen alokasi memori stream I/O secara mendalam.

## 🚀 Daftar Materi & Cakupan Modul

1. **01-encode-json-test** & **02-json-object-test**
   - Memahami konversi dari objek Go (*Struct*) menjadi bentuk teks biner JSON menggunakan fungsi `json.Marshal`.
   - Aturan ekspor field struct (wajib diawali huruf kapital agar bisa diakses oleh package external).

2. **03-decode-json-test**
   - Memproses balik teks JSON menjadi bentuk objek Go menggunakan `json.Unmarshal`.
   - Memahami urgensi penggunaan pointer (`&`) saat operasi decode agar dapat memutasi nilai memori asli objek tujuan.

3. **04-json-array-test**
   - Mengelola tipe data koleksi (*Array/Slice*).
   - Penanganan *JSON Array of Primitives* (`[]string`), *JSON Array of Objects* (`[]Struct`), serta struktur *Pure JSON Array* murni yang diawali kurung siku `[...]`.

4. **05-json-tag-test**
   - Melakukan mapping nama properti PascalCase milik Go menjadi format *lowercase* atau *snake_case* standar API menggunakan metadata tag bawaan (`` `json:"nama_field"` ``).
   - Memanfaatkan opsi `,omitempty` untuk menyembunyikan properti yang bernilai kosong dari teks keluaran JSON.
   - Menggunakan aba-aba khusus `json:"-"` untuk mengisolasi data internal/sensitif (seperti password, penanda state, atau objek koneksi basis data `*sql.DB`) agar tidak ikut bocor ke luar.

5. **06-json-map-test**
   - Penggunaan tipe data dinamis modern `map[string]any` (sebagai alias dari versi lawas `interface{}`) untuk menangani struktur data JSON fleksibel yang bentuk propertinya tidak pasti atau berubah-ubah.

6. **07-stream-decoder-test** & **08-stream-encoder-test**
   - Implementasi efisiensi manajemen RAM menggunakan `json.NewDecoder` dan `json.NewEncoder` untuk membaca/menulis aliran data langsung (*I/O Stream*) dari file fisik (`os.File`) atau request jaringan internet (`r.Body` / `w`).

---

## 💡 Catatan Penting & Praktik Terbaik (Best Practices)

### 🔑 Perbedaan Utama Penampung Map (`var` vs `:=`)
* **`var customer map[string]any` (Nil Map)**
  Sangat direkomendasikan jika tujuan utama map tersebut murni sebagai wadah target operasi decode stream (`decoder.Decode(&customer)`). Fungsi internal Go otomatis akan mengalokasikan ruang memori saat proses penguraian sukses.
* **`customer := map[string]any{}` (Initialized Map)**
  Wajib digunakan jika Anda ingin mengisi nilai properti map tersebut secara manual satu-per-satu di dalam logika kode aplikasi (misal: `customer["name"] = "Eko"`). Jika Anda mencoba mengisi properti secara manual pada kondisi *Nil Map*, program akan langsung mengalami kegagalan fatal (`panic: assignment to entry in nil map`).

### 🛠️ Aturan Emas Pengelolaan File Stream
Setiap kali membuka atau membuat file stream baru menggunakan fungsi `os.Open()` ataupun `os.Create()`, pastikan untuk selalu mengunci pelepasan sisa resource tersebut tepat setelah penanganan error selesai menggunakan perintah **`defer`**:

```go
file, err := os.Open("data.json")
if err != nil {
    panic(err)
}
// Selalu pasang penutup otomatis agar terhindar dari error "too many open files"
defer file.Close()
```

## 📊 Panduan Cepat: Kapan Menggunakan Unmarshal vs Decoder?
Gunakan `json.Marshal` / `json.Unmarshal` jika sumber data JSON sudah nangkring langsung di memori sebagai variabel string atau berupa potongan data mentah (`[]byte`).

Gunakan `json.NewEncoder` / `json.NewDecoder` jika data JSON mengalir secara dinamis melalui jaringan komputer, HTTP Request Body, atau bersumber dari berkas file fisik berukuran besar untuk menghemat konsumsi RAM server.

## 🛠️ Cara Menjalankan Pengujian
Pastikan Anda sudah menginstal Go (versi minimum 1.18 untuk mendukung penggunaan keyword `any`). Masuk ke salah satu direktori materi, kemudian eksekusi perintah pengujian berikut di dalam terminal:
```Bash
# Menjalankan seluruh unit test pada modul tertentu
go test -v ./...

# Menjalankan unit test secara spesifik berdasarkan nama fungsi
go test -v -run=TestStreamDecoderMap
```
