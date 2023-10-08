package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func loadProducts() (map[string]int64, error) {
	purchases := make(map[string]int64)

	bs, err := os.ReadFile(filepath.Join("store", "proto", "products.json"))
	if err != nil {
		return purchases, err
	}

	// JSON
	if err := json.Unmarshal(bs, &purchases); err != nil {
		return purchases, err
	}

	return purchases, nil
}
