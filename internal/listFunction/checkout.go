package listfunction

import (
	"fmt"

	"github.com/virgilIw/koda-b6-golang/model"
)

var histori []model.ItemKeranjang

func Checkout(keranjang *[]model.ItemKeranjang) {

	if len(*keranjang) == 0 {
		fmt.Println("Keranjang kosong")
		return
	}

	histori = append(histori, *keranjang...)

	*keranjang = []model.ItemKeranjang{}

	fmt.Println("Checkout berhasil")
}
