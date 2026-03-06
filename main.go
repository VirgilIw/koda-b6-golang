package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/virgilIw/koda-b6-golang/internal/menu"
	"github.com/virgilIw/koda-b6-golang/model"
)

func main() {
	fmt.Println("Selamat Datang di Pizza Koda Indonesia")

	menuData, err := loadMenuData("assets/data/menu.json", "https://raw.githubusercontent.com/adityabastyas/koda-b6-golang/refs/heads/main/assets/data/menu.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var keranjang []model.ItemKeranjang

	if len(menuData) == 0 {
		fmt.Fprintln(os.Stderr, "data menu kosong")
		os.Exit(1)
	}

	menu.ShowMenu(menuData, &keranjang)
}

func loadMenuData(localPath, fallbackURL string) ([]model.Menu, error) {
	body, err := os.ReadFile(localPath)
	if err == nil {
		return decodeMenu(body)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Get(fallbackURL)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca %q dan gagal request menu dari URL: %w", localPath, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request gagal: %s", res.Status)
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca response body: %w", err)
	}

	return decodeMenu(body)
}

func decodeMenu(body []byte) ([]model.Menu, error) {
	var menuData []model.Menu
	if err := json.Unmarshal(body, &menuData); err != nil {
		return nil, fmt.Errorf("error decode json: %w", err)
	}

	return menuData, nil
}
