package xmd

type BetMode string

func (m BetMode) String() string {
	switch m {
	case BetModeAll:
		return "All - 每期都投注"
	case BetModeWork:
		return "Work - 在[08:30,11:30]、[14:30,17:00]、[19:30,21:30]不投注"
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
	BetModeAll      BetMode = "All"
	BetModeWork     BetMode = "Work"
	BetModeModeAll  BetMode = "ModeAll"
	BetModeModeOnly BetMode = "ModeOnly"
	BetModeHalf     BetMode = "Half"
)

type ConfigCookies struct {
	Bet   string `json:"betting"`
	Prize string `json:"prize"`
}

type Config struct {
	BetMode  BetMode       `json:"bet_mode"`
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
