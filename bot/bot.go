package bot

const (
	SLACK_TOKEN = "SLACK_TOKEN"
	DB_PATH = "DB_PATH"
)

var (
	Ultramagnus	*ChatBot
)

func Setup(c *Config) *ChatBot {
	Ultramagnus = Ultramagnus.New()
	Ultramagnus.Init(c)

	return Ultramagnus
}
