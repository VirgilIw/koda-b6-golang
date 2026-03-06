package listfunction

import "fmt"

func ShowHistory() {
	if len(histori) == 0 {
		fmt.Println("Histori kosong")
		return
	}

	fmt.Println("== HISTORI PESANAN ==")

	total := 0

	for i, h := range histori {

		subtotal := h.Harga * h.Jumlah
		total += subtotal

		fmt.Printf("%d. %s | %d x %d = %d\n",
			i+1,
			h.Nama,
			h.Jumlah,
			h.Harga,
			subtotal)
	}

	fmt.Println("Total histori:", total)
}
