package slack

import (
	"fmt"
	"github.com/corny90/social-app-pkg/core-domain/message"
	logger "github.com/corny90/social-app-pkg/logger"
	"github.com/slack-go/slack"
)

func (adapter *Adapter) LogRequest(logInfo logger.LogRequest) {
	go func() {
		// CONSTRUCT THE FORMATTED MESSAGE TEXT WITH EMOJIS
		messageText := fmt.Sprintf("%v `%v` ⏰ `%v` 🀄️`%v · %v` ⌛ `%v` 📍 `%v`",
			statusEmoji(logInfo.LogStatusCode),
			logInfo.LogType,
			logInfo.LogStartTimestamp,
			logInfo.LogMethod,
			logInfo.LogStatusCode,
			logInfo.LogDuration,
			logInfo.LogPath,
		)

		// SEND THE MESSAGE TO SLACK
		_, _, err := adapter.client.Client.PostMessage(adapter.channelLogActivity, slack.MsgOptionText(messageText, false))
		if err != nil {
			logger.Event("ERROR", "failed posting chat to Slack: %v", err)
		}
	}()
}

func (adapter *Adapter) LogEvent(logObj logger.LogEvent) {
	go func() {
		// CONSTRUCT THE FORMATTED MESSAGE TEXT WITH EMOJIS
		messageText := fmt.Sprintf("%v `%v` ⏰ `%v` 📝 `%v`",
			logEmoji(logObj.LogType),
			logObj.LogType,
			logObj.LogStartTimestamp,
			logObj.LogMessage,
		)

		// SEND THE MESSAGE TO SLACK
		_, _, err := adapter.client.Client.PostMessage(adapter.channelLogActivity, slack.MsgOptionText(messageText, false))
		if err != nil {
			logger.Event("ERROR", "failed posting chat to Slack: %v", err)
		}
	}()
}

func (adapter *Adapter) LogChatMessage(message message.Message, agentPersonality string, regenerated bool) {
	go func() {
		// CONSTRUCT THE FORMATTED MESSAGE TEXT WITH EMOJIS
		var messageText string
		if message.SenderType == "R" {
			messageText = fmt.Sprintf("⏰ `%v` 🧔🏻‍ `%v` 💬 `%v`",
				message.CreatedAt.Format("2006/01/02 15:04:05"),
				message.Sender,
				message.Content,
			)
		} else if message.SenderType == "V" {
			vSymbol := "👽‍"
			if regenerated {
				vSymbol = "🔄"
			}
			messageText = fmt.Sprintf("⏰ `%v` %v `%v` 💬 `%v` `%v`",
				message.CreatedAt.Format("2006/01/02 15:04:05"),
				vSymbol,
				message.Sender,
				message.Content,
				agentPersonality,
			)
		}

		// SEND THE MESSAGE TO SLACK
		_, _, err := adapter.client.Client.PostMessage(adapter.channelLogChats, slack.MsgOptionText(messageText, false))
		if err != nil {
			logger.Event("ERROR", "failed posting message to Slack: %v", err)
		}
	}()
}

func (adapter *Adapter) LogHttp(logInfo logger.LogRequest) {
	go func() {
		// CONSTRUCT THE FORMATTED MESSAGE TEXT WITH EMOJIS
		messageText := fmt.Sprintf("%v `%v` ⏰ `%v` ⌛ `%v` 🀄️`%v · %v` 📍 `%v`",
			statusEmoji(logInfo.LogStatusCode),
			logInfo.LogType,
			logInfo.LogStartTimestamp,
			logInfo.LogDuration,
			logInfo.LogMethod,
			logInfo.LogStatusCode,
			logInfo.LogPath,
		)

		// SEND THE MESSAGE TO SLACK
		_, _, err := adapter.client.Client.PostMessage(adapter.channelLogActivity, slack.MsgOptionText(messageText, false))
		if err != nil {
			logger.Event("ERROR", "failed posting message to Slack: %v", err)
		}
	}()
}

func (adapter *Adapter) LogWebsocket(logInfo logger.LogWS) {
	go func() {
		// CONSTRUCT THE FORMATTED MESSAGE TEXT WITH EMOJIS
		messageText := fmt.Sprintf("%v `%v` ⏰ `%v` ⌛ `%v` 🀄️`%v · %v` 📨 `%v`",
			statusEmoji(logInfo.LogStatusCode),
			logInfo.LogType,
			logInfo.LogStartTimestamp,
			logInfo.LogDuration,
			logInfo.LogAction,
			logInfo.LogStatusCode,
			logInfo.LogData,
		)

		// SEND THE MESSAGE TO SLACK
		_, _, err := adapter.client.Client.PostMessage(adapter.channelLogActivity, slack.MsgOptionText(messageText, false))
		if err != nil {
			logger.Event("ERROR", "failed posting message to Slack: %v", err)
		}
	}()
}

func logEmoji(logType string) string {
	switch logType {
	case "[INFO_]":
		return "🟦"
	case "[DEBUG]":
		return "🟨"
	case "[WARN_]":
		return "🟧"
	case "[ERROR]":
		return "🟥"
	case "[FATAL]":
		return "⬛️"
	default:
		return "ℹ️" // Default icon if the log type is not recognized
	}
}

func statusEmoji(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "✅" // Success
	case code >= 400 && code < 500:
		return "⚠️" // Client error
	case code >= 500:
		return "❌" // Server error
	default:
		return "ℹ️" // Default case
	}
}
