package pointers

import (
	"comparasion/resources"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Server struct {
	*gin.Engine
	service *resources.Service
}

func NewServer(service *resources.Service) *Server {
	return &Server{
		Engine:  gin.Default(),
		service: service,
	}
}

func (s *Server) Start(port string) error {
	return s.Run(port)
}

func (s *Server) SetRouters() {

	apiV1Group := s.Group("/api/v1")
	apiV1Group.POST("/resources", s.createResources)
	apiV1Group.GET("/resources", s.getAllResources)
	apiV1Group.GET("/resources/:id", s.getResources)
	apiV1Group.PUT("/resources", s.updateResources)
	apiV1Group.DELETE("/resources/:id", s.deleteResources)
}

func (s *Server) createResources(c *gin.Context) {
	var newResources resources.Resources

	if err := c.BindJSON(&newResources); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	result, err := s.service.Create(newResources)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func (s *Server) getResources(c *gin.Context) {
	id := c.Param("id")

	result, err := s.service.Get(uuid.MustParse(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func (s *Server) getAllResources(c *gin.Context) {
	result, err := s.service.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func (s *Server) updateResources(c *gin.Context) {
	var newResources resources.Resources

	if err := c.BindJSON(&newResources); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	result, err := s.service.Update(newResources)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func (s *Server) deleteResources(c *gin.Context) {
	id := c.Param("id")

	err := s.service.Delete(uuid.MustParse(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
