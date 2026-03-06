package menu

import (
	"fmt"

	"github.com/virgilIw/koda-b6-golang/input"
	listfunction "github.com/virgilIw/koda-b6-golang/internal/listFunction"
	"github.com/virgilIw/koda-b6-golang/model"
)

func ShowMenu(menu []model.Menu, keranjang *[]model.ItemKeranjang) {
	for {

		fmt.Println("\n== MENU UTAMA ==")
		fmt.Println("1. Lihat menu pizza")
		fmt.Println("2. Tambah ke keranjang")
		fmt.Println("3. Lihat keranjang")
		fmt.Println("4. Hapus item keranjang")
		fmt.Println("5. Checkout")
		fmt.Println("6. Lihat histori")
		fmt.Println("7. Keluar")

		pilihan := input.TanyaUser("Pilih nomer = ")

		switch pilihan {

		case "1":
			// Perbaikan: Tambah parameter keranjang
			listfunction.ShowPizza(menu, keranjang)

		case "2":
			listfunction.AddCart(menu, keranjang)

		case "3":
			listfunction.ShowCart(*keranjang)

		case "4":
			listfunction.DeleteCart(keranjang)

		case "5":
			listfunction.Checkout(keranjang)

		case "6":
			listfunction.ShowHistory()

		case "7":
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
