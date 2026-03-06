package listfunction

import (
	"fmt"
	"strconv"

	"github.com/virgilIw/koda-b6-golang/input"
	"github.com/virgilIw/koda-b6-golang/model"
)

func ShowPizza(menu []model.Menu, keranjang *[]model.ItemKeranjang) { // Tambah parameter keranjang

	for {
		fmt.Println("\n== DAFTAR MENU PIZZA ==")

		for _, m := range menu {
			fmt.Printf("%d. %s - Rp%d\n", m.Id, m.Name, m.Harga)
		}

		fmt.Println("0. Kembali ke menu utama")

		pilih := input.TanyaUser("\nPilih nomor pizza: ")

		if pilih == "0" {
			// Tampilkan total harga keranjang sebelum kembali
			tampilkanTotalKeranjang(keranjang)
			return
		}

		// Konversi pilihan ke integer
		idPilih, err := strconv.Atoi(pilih)
		if err != nil {
			fmt.Println("Input tidak valid!")
			continue
		}

		// Cari menu yang dipilih
		var menuDipilih *model.Menu
		for _, m := range menu {
			if m.Id == idPilih {
				menuDipilih = &m
				break
			}
		}

		if menuDipilih == nil {
			fmt.Println("Menu tidak ditemukan!")
			continue
		}

		// Tanya jumlah yang dipesan
		jumlahStr := input.TanyaUser("Masukkan jumlah: ")
		jumlah, err := strconv.Atoi(jumlahStr)
		if err != nil || jumlah <= 0 {
			fmt.Println("Jumlah tidak valid!")
			continue
		}

		// Cek apakah item sudah ada di keranjang
		itemExist := false
		for i, item := range *keranjang {
			if item.IdMenu == menuDipilih.Id {
				// Update jumlah jika item sudah ada
				(*keranjang)[i].Jumlah += jumlah
				itemExist = true
				fmt.Printf("✓ %s ditambahkan (%d porsi) - Total sekarang: %d porsi\n",
					menuDipilih.Name, jumlah, (*keranjang)[i].Jumlah)
				break
			}
		}

		// Jika item belum ada, tambah baru
		if !itemExist {
			itemBaru := model.ItemKeranjang{
				IdMenu: menuDipilih.Id,
				Nama:   menuDipilih.Name,
				Harga:  menuDipilih.Harga,
				Jumlah: jumlah,
			}
			*keranjang = append(*keranjang, itemBaru)
			fmt.Printf("✓ %s (%d porsi) ditambahkan ke keranjang\n", menuDipilih.Name, jumlah)
		}

		// Tampilkan total sementara
		tampilkanTotalKeranjang(keranjang)

		// Tanya apakah mau pesan lagi
		tanyaLagi := input.TanyaUser("\nMau pesan lagi? (y/n): ")
		if tanyaLagi != "y" && tanyaLagi != "Y" {
			return
		}
	}
}

// Fungsi helper untuk menampilkan total keranjang
func tampilkanTotalKeranjang(keranjang *[]model.ItemKeranjang) {
	if len(*keranjang) == 0 {
		fmt.Println("\nKeranjang belanja kosong")
		return
	}

	fmt.Println("\n== RINGKASAN KERANJANG ==")
	var total int
	for i, item := range *keranjang {
		subtotal := item.Harga * item.Jumlah
		total += subtotal
		fmt.Printf("%d. %s - %d porsi x Rp%d = Rp%d\n",
			i+1, item.Nama, item.Jumlah, item.Harga, subtotal)
	}
	fmt.Println("------------------------")
	fmt.Printf("TOTAL: Rp%d\n", total)
}
