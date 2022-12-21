package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/matthi01/go-gin-api-starter/db"
	"net/http"
	"strconv"
)

func GetItems(data *db.List) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		results := data.GetAll()
		ctx.JSON(http.StatusOK, results)
	}
}

func GetItem(data *db.List) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sId := ctx.Param("id")
		id, err := strconv.Atoi(sId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("bad item id: %s", sId),
			})
			return
		}
		result, ok := data.Get(id)
		if !ok {
			ctx.JSON(http.StatusNotFound, map[string]string{
				"error": fmt.Sprintf("item with id %d not found", id),
			})
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

type PostRequestData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateItem(data *db.List) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody := PostRequestData{}
		err := ctx.Bind(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "request body could not be parsed",
			})
			return
		}
		// TODO: need to run some data validation on requestBody data before adding item
		addedItem, ok := data.Add(requestBody.Name, requestBody.Description)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error": "an error has occurred",
			})
			return
		}
		ctx.JSON(http.StatusOK, addedItem)
	}
}

func UpdateItem(data *db.List) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// id validation...
		sId := ctx.Param("id")
		id, err := strconv.Atoi(sId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("bad item id: %s", sId),
			})
			return
		}
		_, ok := data.Get(id)
		if !ok {
			ctx.JSON(http.StatusNotFound, map[string]string{
				"error": fmt.Sprintf("item with id %d not found", id),
			})
			return
		}

		// request body validation
		requestBody := PostRequestData{}
		err = ctx.Bind(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "request body could not be parsed",
			})
			return
		}

		// request is clean - update the item
		updatedItem, ok := data.Update(id, requestBody.Name, requestBody.Description)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error": "an error has occurred",
			})
			return
		}
		ctx.JSON(http.StatusOK, updatedItem)
	}
}

// TODO: Need to get rid of some of the validation code duplication
func DeleteItem(data *db.List) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// id validation...
		sId := ctx.Param("id")
		id, err := strconv.Atoi(sId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("bad item id: %s", sId),
			})
			return
		}
		_, ok := data.Get(id)
		if !ok {
			ctx.JSON(http.StatusNotFound, map[string]string{
				"error": fmt.Sprintf("item with id %d not found", id),
			})
			return
		}

		// id is valid, delete item
		removedItem, ok := data.Delete(id)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error": "an error has occurred",
			})
			return
		}
		ctx.JSON(http.StatusOK, removedItem)
	}
}
