package listfunction

import (
	"fmt"

	"github.com/virgilIw/koda-b6-golang/model"
)

func ShowCart(keranjang []model.ItemKeranjang) {
	if len(keranjang) == 0 {
		fmt.Println("Keranjang kosong")
		return
	}

	fmt.Println("== KERANJANG ==")

	total := 0

	for i, k := range keranjang {

		subtotal := k.Harga * k.Jumlah
		total += subtotal

		fmt.Printf("%d. %s | %d x %d = %d\n",
			i+1,
			k.Nama,
			k.Jumlah,
			k.Harga,
			subtotal)
	}

	fmt.Println("Total:", total)
}
