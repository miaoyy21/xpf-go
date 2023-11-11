package cache

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"os"
	"path/filepath"
	"xpf/pb"
)

type application struct {
	AppId   string            `json:"appId"`   // 应用ID
	AppName map[string]string `json:"appName"` // 应用名称
}

func loadApplications() (map[string]int32, []*pb.Application, error) {
	as := make([]*application, 0)

	mas := make(map[string]int32)
	pbs := make([]*pb.Application, 0)

	bs, err := os.ReadFile(filepath.Join("store", "proto", "applications.json"))
	if err != nil {
		return nil, nil, err
	}

	// JSON
	if err := json.Unmarshal(bs, &as); err != nil {
		return nil, nil, err
	}

	for _, a := range as {
		p := &pb.Application{
			AppId:   proto.String(a.AppId),
			AppName: a.AppName,
		}

		pbs = append(pbs, p)
	}

	return mas, pbs, nil
}
