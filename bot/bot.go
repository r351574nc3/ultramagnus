package bot

const (
	SLACK_SECRETS = "SLACK_SECRETS"
	DB_PATH = "DB_PATH"
	HTTP_CONFIG = "HTTP_CONFIG"
)

var (
	Ultramagnus	*ChatBot
	SlackHandlers *BotHandlers
)

func Setup(c *Config) *ChatBot {
	Ultramagnus = Ultramagnus.New()
	Ultramagnus.Init(c)

	SlackHandlers = SlackHandlers.New()

	return Ultramagnus
}
