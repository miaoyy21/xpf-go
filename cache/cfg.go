package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	Version                       string `json:"version"`
	AssetsVersion                 string `json:"assetsVersion"`
	AppStoreVerifyReceiptPassword string `json:"appStoreVerifyReceiptPassword"`
	Mode                          string `json:"mode"`
	Host                          string `json:"host"`
	Port                          string `json:"port"`
	DbHost                        string `json:"dbHost"`
	DbPort                        string `json:"dbPort"`
	DbDatabase                    string `json:"dbDatabase"`
	DbUser                        string `json:"dbUser"`
	DbPassword                    string `json:"dbPassword"`
	QrCode                        string `json:"qrCode"`
}

func loadConfig() (*config, error) {
	cfg := new(config)

	bs, err := os.ReadFile(filepath.Join("store", "psw.json"))
	if err != nil {
		return nil, err
	}

	// JSON
	if err := json.Unmarshal(bs, &cfg); err != nil {
		return nil, err
	}

	cfg.Mode = strings.ToLower(cfg.Mode)

	return cfg, nil
}
