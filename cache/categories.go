package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
	"psw/pb"
)

type CategoryField struct {
	Name      map[string]string `json:"name"`
	IsPrimary bool              `json:"isPrimary"`
	Type      pb.FieldType      `json:"type"`
	MinLines  int32             `json:"minLines"`
	MaxLines  int32             `json:"maxLines"`
}

type category struct {
	ProtoId int32             `json:"protoId"`
	Name    map[string]string `json:"name"`   // 分类名称
	Fields  []*CategoryField  `json:"fields"` // 默认信息要素
}

type categories struct {
	Default    []*CategoryField `json:"default"`
	Categories []*category      `json:"categories"`
}

func loadCategories() ([]*CategoryField, []*category, error) {
	categories := new(categories)

	bs, err := os.ReadFile(filepath.Join("store", "proto", "categories.json"))
	if err != nil {
		return nil, nil, err
	}

	// JSON
	if err := json.Unmarshal(bs, &categories); err != nil {
		return nil, nil, err
	}

	return categories.Default, categories.Categories, nil
}
