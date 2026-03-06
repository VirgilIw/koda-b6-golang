package listfunction

import (
	"fmt"

	"github.com/virgilIw/koda-b6-golang/model"
)

func DeleteCart(keranjang *[]model.ItemKeranjang) {

	if len(*keranjang) == 0 {
		fmt.Println("Keranjang kosong")
		return
	}

	var index int

	fmt.Print("Masukkan nomor item yang akan dihapus: ")
	fmt.Scanln(&index)

	if index <= 0 || index > len(*keranjang) {
		fmt.Println("Index tidak valid")
		return
	}

	*keranjang = append((*keranjang)[:index-1], (*keranjang)[index:]...)

	fmt.Println("Item berhasil dihapus")
}
