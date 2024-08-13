package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

func DecodePageState(ctx iris.Context, pageState string) []byte {
	pageStateParam := ctx.URLParamDefault(pageState, "")

	if pageStateParam == "" {
		return nil
	}

	decodedState, err := base64.URLEncoding.DecodeString(pageStateParam)
	if err != nil {
		golog.Errorf("Error validating %s: %v", pageState, err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": fmt.Sprintf("invalid %s", pageState)})
		return nil
	}

	return decodedState
}

func EncodePageState(pageState []byte) string {
	return base64.URLEncoding.EncodeToString(pageState)
}
