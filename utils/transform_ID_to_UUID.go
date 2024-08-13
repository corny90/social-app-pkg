package utils

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

func TransformParamIdToUUID(ctx iris.Context, paramName string) (gocql.UUID, bool) {
	paramValue := ctx.URLParam(paramName)
	if paramValue == "" {
		golog.Errorf("[%s] is required, but received empty", paramName)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": fmt.Sprintf("[%s] is required", paramName)})
		return gocql.UUID{}, true
	}

	parsedUUID, err := gocql.ParseUUID(paramValue)
	if err != nil {
		golog.Errorf(err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return gocql.UUID{}, true
	}

	return parsedUUID, false
}
