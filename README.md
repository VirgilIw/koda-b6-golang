Aplikasi Pemesanan Pizza - CLI
Aplikasi sederhana berbasis terminal untuk pemesanan pizza menggunakan Golang.

Fitur
1. Lihat menu pizza

2. Tambah ke keranjang

3. Lihat keranjang

4. Hapus item keranjang

5. Checkout

6. Lihat histori transaksi

- Cara Menjalankan
git clone https://github.com/virgilIw/koda-b6-golang.git
cd koda-b6-golang
go run main.go
Struktur Proyek
text
├── main.go              # Entry point
├── menu/                # Menu utama
├── internal/listFunction # Fungsi-fungsi utama
├── model/               # Struct data
├── input/               # Utility input
└── data/                # Data dummy
Cara Penggunaan
Jalankan aplikasi: go run main.go

Pilih menu dengan memasukkan angka:

1 - Lihat menu pizza

2 - Tambah ke keranjang

3 - Lihat keranjang

4 - Hapus item

5 - Checkout

6 - Lihat histori

7 - Keluar

CLI (Command Line Interface)
