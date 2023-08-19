package xmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func (o *Cache) Save() error {
	dFileName := filepath.Join(o.dir, "data.json")
	dFile, err := os.Create(dFileName)
	if err != nil {
		log.Panicf("unexpected :: os.CreateFile(%s) failure : %s \n", dFileName, err.Error())
	}
	defer dFile.Close()

	js := json.NewEncoder(dFile)
	js.SetIndent("", "\t")
	if err := js.Encode(o.Prizes); err != nil {
		return err
	}

	return nil
}
