package bot

const (
	SLACK_TOKEN = "SLACK_TOKEN"
	DB_PATH = "DB_PATH"
	SLACK_SIGNING_SECRET = "SLACK_SIGNING_SECRET"
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
