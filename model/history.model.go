package model

// Struct untuk Histori Transaksi
type HistoriTransaksi struct {
	IDTransaksi string
	Tanggal     string
	Items       []ItemKeranjang
	TotalHarga  int
}
