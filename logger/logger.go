package pkg_logger

import (
	"fmt"
	environment "github.com/corny90/social-app-pkg/environment"
	"os"
	"time"
)

// ANSI COLOR CODES
const (
	ColorRed           = "\033[31m"
	ColorRedBright     = "\033[91m"
	ColorGreen         = "\033[32m"
	ColorGreenBright   = "\033[92m"
	ColorYellow        = "\033[33m"
	ColorYellowBright  = "\033[93m"
	ColorBlue          = "\033[34m"
	ColorBlueBright    = "\033[94m"
	ColorMagenta       = "\033[35m"
	ColorMagentaBright = "\033[95m"
	ColorCyan          = "\033[36m"
	ColorCyanBright    = "\033[96m"
	ColorWhite         = "\033[37m"
	ColorWhiteBright   = "\033[97m"
	ColorGray          = "\033[90m"
	ColorGrayBright    = "\033[37m"
	ColorReset         = "\033[0m"
)

type LogRequest struct {
	LogType           string        // Type of log
	LogStartTimestamp string        // Timestamp when the request started
	LogDuration       time.Duration // Duration of the request
	LogStatusCode     int           // HTTP status code of the response
	LogMethod         string        // HTTP method of the request
	LogPath           string        // Path of the HTTP request
	LogEnvironment    string        // Environment in which the log was generated
}

type LogEvent struct {
	LogType           string
	LogEnvironment    string
	LogStartTimestamp string
	LogMessage        string
}

func Event(logType, msg string, args ...interface{}) LogEvent {
	if environment.EnvName == "PRODUCTION" && logType == "DEBUG" {
		return LogEvent{}
	}

	logEvent := LogEvent{
		LogType:           GetLogTypeLabel(logType),
		LogEnvironment:    environment.EnvName,
		LogStartTimestamp: time.Now().Format("2006/01/02 15:04:05"),
		LogMessage:        fmt.Sprintf(msg, args...),
	}

	log := fmt.Sprintf("%s%s%s %s%s%s %v%v%v %v%v%v",
		GetLogTypeColor(logType), logEvent.LogType, ColorReset,
		ColorWhite, logEvent.LogStartTimestamp, ColorReset,
		ColorGray, "|", ColorReset,
		ColorYellow, logEvent.LogMessage, ColorReset,
	)

	fmt.Println(log)

	// EXIT THE PROGRAM IF THE LOG TYPE IS FATAL
	if logType == "FATAL" {
		os.Exit(1)
	}

	return logEvent
}

func ReqLog(logStatusCode int, logMethod string, logPath string) LogRequest {
	start := time.Now()

	logInfo := LogRequest{
		LogType:           GetLogTypeLabel("REQ"),
		LogStartTimestamp: start.Format("2006/01/02 15:04:05"),
		LogDuration:       time.Since(start),
		LogStatusCode:     logStatusCode,
		LogMethod:         logMethod,
		LogPath:           logPath,
		LogEnvironment:    environment.EnvName,
	}

	logMessage := fmt.Sprintf("%s%s%s %s%s%s %v%v%v %s%s%s %s%d%s %v%v%v %s%s%s %v%v%v %s%s%s",
		ColorCyan, logInfo.LogType, ColorReset,
		ColorWhite, logInfo.LogStartTimestamp, ColorReset,
		ColorGray, "|", ColorReset,

		GetStatusCodeColor(logInfo.LogStatusCode), logInfo.LogMethod, ColorReset,
		GetStatusCodeColor(logInfo.LogStatusCode), logInfo.LogStatusCode, ColorReset,
		ColorGray, "|", ColorReset,

		ColorWhite, logInfo.LogDuration, ColorReset,
		ColorGray, "|", ColorReset,

		ColorYellow, logInfo.LogPath, ColorReset,
	)
	fmt.Println(logMessage)
	return logInfo
}

func GetStatusCodeColor(statusCode int) string {
	var statusColor string
	switch {
	case statusCode >= 200 && statusCode < 300:
		statusColor = ColorGreenBright
	case statusCode >= 300 && statusCode < 400:
		statusColor = ColorCyanBright
	case statusCode >= 400 && statusCode < 500:
		statusColor = ColorYellowBright
	case statusCode >= 500:
		statusColor = ColorRedBright
	default:
		statusColor = ColorReset
	}
	return statusColor
}

func GetMethodColor(methodType string) string {
	var methodColor string
	switch methodType {
	case "GET":
		methodColor = ColorBlueBright
	case "POST":
		methodColor = ColorGreenBright
	case "PUT":
		methodColor = ColorYellowBright
	case "DELETE":
		methodColor = ColorRedBright
	case "PATCH":
		methodColor = ColorMagentaBright
	case "OPTIONS":
		methodColor = ColorWhiteBright
	default:
		methodColor = ColorReset
	}
	return methodColor
}

func GetLogTypeLabel(logType string) string {
	var label string
	switch logType {
	case "REQ":
		label = "[REQ  ]"
	case "INFO":
		label = "[INFO ]"
	case "DEBUG":
		label = "[DEBUG]"
	case "WARN":
		label = "[WARN ]"
	case "ERROR":
		label = "[ERROR]"
	case "FATAL":
		label = "[FATAL]"
	default:
		label = "[?????]"
	}
	return label
}

func GetLogTypeColor(logType string) string {
	switch logType {
	case "REQ":
		return ColorCyan
	case "INFO":
		return ColorCyan
	case "DEBUG":
		return ColorYellow
	case "WARN":
		return ColorYellowBright
	case "ERROR":
		return ColorRed
	case "FATAL":
		return ColorRedBright
	default:
		return ""
	}
}
