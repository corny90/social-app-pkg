package utils

func HttpRespOk(message string, data ...interface{}) interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   data,
		"chat":   message,
	}
}

func HttpRespError(code int, message string) interface{} {
	return map[string]interface{}{
		"status": "error",
		"code":   code,
		"chat":   message,
	}
}
