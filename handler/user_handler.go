package handler

import (
	"1-loyiha/model"
	"1-loyiha/service"
	"1-loyiha/validation"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)



type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) * UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine){
	r.GET("/users", h.GetAll)
	r.GET("/users/:id", h.GetByID)
	r.POST("/users", h.Create)
	r.PUT("/users/:id", h.Update)
	r.DELETE("/users/:id", h.Delete)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetAll())
} 

func (h *UserHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := h.service.GetByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	}
	c.JSON(http.StatusOK, user)
}


func (h *UserHandler) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, h.service.Create(user))
}

func (h *UserHandler) Update(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validation.ValidationUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return 
	}
	updated := h.service.Update(id, user)
	if updated == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if !h.service.Delete(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.Status(http.StatusNoContent)
}