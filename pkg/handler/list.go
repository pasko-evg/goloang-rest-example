package handler

import (
	"github.com/gin-gonic/gin"
	goRest "go-rest"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input goRest.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": listId,
	})
}

type getAllListsResponse struct {
	Data []goRest.TodoList `json:"data"`
}

func (h *Handler) getAllList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.TodoList.GetAll(userId)

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, idErr := strconv.Atoi(c.Param("id"))
	if idErr != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, listErr := h.services.TodoList.GetById(userId, listId)
	if listErr != nil {
		newErrorResponse(c, http.StatusInternalServerError, listErr.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, idErr := strconv.Atoi(c.Param("id"))
	if idErr != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input goRest.UpdateListInput
	if inputErr := c.BindJSON(&input); inputErr != nil {
		newErrorResponse(c, http.StatusBadRequest, inputErr.Error())
	}

	if err = h.services.TodoList.Update(userId, listId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, idErr := strconv.Atoi(c.Param("id"))
	if idErr != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	listErr := h.services.TodoList.Delete(userId, listId)
	if listErr != nil {
		newErrorResponse(c, http.StatusInternalServerError, listErr.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
