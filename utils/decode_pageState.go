package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

func GetPageStateFromParam(ctx iris.Context, paramName string) ([]byte, bool) {
	pageStateParam := ctx.URLParamDefault(paramName, "")

	if pageStateParam == "" {
		return nil, true // Return true as it's not an error to have no page state
	}

	decodedState, err := base64.URLEncoding.DecodeString(pageStateParam)
	if err != nil {
		golog.Errorf("Error validating %s: %v", paramName, err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": fmt.Sprintf("invalid %s", paramName)})
		return nil, false
	}

	return decodedState, true
}
