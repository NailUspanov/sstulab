package handlers

import (
	"data-service/internal/helpers"
	"data-service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/add", h.add)
		api.GET("/report1", h.report)
		api.GET("/report2", h.report2)
		api.GET("/report3", h.report3)
	}
	return router
}

func (h *Handler) add(c *gin.Context) {

	var input service.AddRequest
	if err := c.BindJSON(&input); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// check if url is valid
	//_, err := url.ParseRequestURI(input.LongUrl)
	//if err != nil {
	//	helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid URI for request")
	//	return
	//}

	h.services.FilmService.QueueImplementation(input)

	c.JSON(http.StatusOK, map[string]any{
		"status": "ok",
	})
}

func (h *Handler) report(c *gin.Context) {

	res, err := h.services.FilmService.GetTopDirectorsByFilms()
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) report2(c *gin.Context) {

	res, err := h.services.FilmService.GetTopFilmsByRating()
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) report3(c *gin.Context) {

	res, err := h.services.FilmService.GetTopFilmsByReviewsCount()
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}