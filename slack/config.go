package slack

import (
	logger "github.com/corny90/social-app-pkg/logger"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"strings"
)

type Adapter struct {
	client             *socketmode.Client
	channelLogActivity string
	channelLogChats    string
}

// NewSlackAdapter initializes and returns a new SlackLogger.
func NewSlackAdapter(appToken, botToken, channelLogActivity, channelLogChats string) *Adapter {

	//appToken := os.Getenv("SLACK_SOCKET_MODE_TOKEN")
	if appToken == "" {
		logger.Event("FATAL", "SLACK_SOCKET_MODE_TOKEN must be set.")
	}
	if !strings.HasPrefix(appToken, "xapp-") {
		logger.Event("FATAL", "SLACK_SOCKET_MODE_TOKEN must have the prefix \"xapp-\".")
	}

	//botToken := os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN")
	if botToken == "" {
		logger.Event("FATAL", "SLACK_BOT_USER_OAUTH_TOKEN must be set.")
	}
	if !strings.HasPrefix(botToken, "xoxb-") {
		logger.Event("FATAL", "SLACK_BOT_USER_OAUTH_TOKEN must have the prefix \"xoxb-\".")
	}

	api := slack.New(
		botToken,
		slack.OptionDebug(false),
		//slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
		slack.OptionAppLevelToken(appToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(false),
		//socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	logger.Event("INFO", ">>> Adapter Slack: Initialized")

	return &Adapter{
		client:             client,
		channelLogActivity: channelLogActivity,
		channelLogChats:    channelLogChats,
	}
}
