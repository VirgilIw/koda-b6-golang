package listfunction

import (
	"fmt"
	"strconv"

	"github.com/virgilIw/koda-b6-golang/input"
	"github.com/virgilIw/koda-b6-golang/model"
)

func AddCart(menu []model.Menu, keranjang *[]model.ItemKeranjang) {
	fmt.Println("\n== TAMBAH KE KERANJANG ==")

	// Tampilkan menu
	for _, m := range menu {
		fmt.Printf("%d. %s - Rp%d\n", m.Id, m.Name, m.Harga)
	}

	// Minta user memilih menu
	pilih := input.TanyaUser("\nPilih nomor menu: ")
	idPilih, _ := strconv.Atoi(pilih)

	// Cari menu
	var menuDipilih *model.Menu
	for _, m := range menu {
		if m.Id == idPilih {
			menuDipilih = &m
			break
		}
	}

	if menuDipilih == nil {
		fmt.Println("Menu tidak ditemukan!")
		return
	}

	// Minta jumlah
	jumlahStr := input.TanyaUser("Masukkan jumlah: ")
	jumlah, _ := strconv.Atoi(jumlahStr)

	if jumlah <= 0 {
		fmt.Println("Jumlah tidak valid!")
		return
	}

	// Tambah ke keranjang (logika sama seperti di ShowPizza)
	itemBaru := model.ItemKeranjang{
		IdMenu: menuDipilih.Id,
		Nama:   menuDipilih.Name,
		Harga:  menuDipilih.Harga,
		Jumlah: jumlah,
	}
	*keranjang = append(*keranjang, itemBaru)
	fmt.Printf("✓ %s (%d porsi) ditambahkan ke keranjang\n", menuDipilih.Name, jumlah)
}
