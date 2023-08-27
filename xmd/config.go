package xmd

type BetMode string

func (m BetMode) String() string {
	switch m {
	case BetModeCustom:
		return "Custom - 在特定的时间段进行投注"
	case BetModeModeAll:
		return "ModeAll - 在45秒～50秒时，选择权重值最大的自动投注模式（大小奇偶中边大尾小尾）和其它符合统计的数字，且权重值大于350的投注"
	case BetModeModeOnly:
		return "ModeOnly - 在45秒～50秒时，仅选择权重值最大的自动投注模式（大小奇偶中边大尾小尾），且权重值大于350的投注"
	case BetModeHalf:
		return "Half - 按幸运值从高到低，只选择约50%的数字"
	default:
		return "<<< Undefined >>>"
	}
}

var (
	BetModeCustom   BetMode = "Custom"
	BetModeModeAll  BetMode = "ModeAll"
	BetModeModeOnly BetMode = "ModeOnly"
	BetModeHalf     BetMode = "Half"
)

type ConfigCookies struct {
	Bet   string `json:"betting"`
	Prize string `json:"prize"`
}

type CustomTime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Config struct {
	BetMode  BetMode       `json:"bet_mode"`
	Custom   []CustomTime  `json:"custom"`
	Seed     int64         `json:"seed"`
	Secs     float64       `json:"secs"`
	Cookies  ConfigCookies `json:"cookies"`
	Agent    string        `json:"agent"`
	UserId   string        `json:"user_id"`
	Token    string        `json:"token"`
	Unix     string        `json:"unix"`
	KeyCode  string        `json:"key_code"`
	DeviceId string        `json:"device_id"`
}
