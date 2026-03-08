package listfunction

import (
	"fmt"
	"sync"

	"github.com/virgilIw/koda-b6-golang/model"
)

var histori []model.ItemKeranjang
var mu sync.Mutex

func Checkout(keranjang *[]model.ItemKeranjang) {

	mu.Lock()
	defer mu.Unlock()

	if len(*keranjang) == 0 {
		fmt.Println("Keranjang kosong")
		return
	}

	histori = append(histori, *keranjang...)

	*keranjang = []model.ItemKeranjang{}

	fmt.Println("Checkout berhasil")
}
