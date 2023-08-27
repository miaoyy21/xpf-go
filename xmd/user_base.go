package xmd

import (
	"fmt"
	"strings"
)

type UserBase struct {
	Seed    int64
	BetMode BetMode
	Custom  []CustomTime

	cookies ConfigCookies
	agent   string
	unix    string
	code    string
	device  string
	id      string
	token   string
}

func (o UserBase) CustomString() string {
	ss := make([]string, 0, len(o.Custom))
	for _, cs := range o.Custom {
		ss = append(ss, fmt.Sprintf("%s ~ %s", cs.Start, cs.End))
	}

	return strings.Join(ss, "、")
}

func NewUserBase(seed int64, betMode BetMode, custom []CustomTime, cookies ConfigCookies, agent string, unix string, code string, device string, id string, token string) UserBase {
	return UserBase{
		Seed:    seed,
		BetMode: betMode,
		Custom:  custom,

		cookies: cookies,
		agent:   agent,
		unix:    unix,
		code:    code,
		device:  device,
		id:      id,
		token:   token,
	}
}
