package pointers

import (
	"comparasion/resources"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Server struct {
	*gin.Engine
	service resources.Service
}

func NewServer() *Server {
	g := gin.New()
	gin.SetMode(gin.ReleaseMode)
	gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/resources")
	return &Server{
		Engine: g,
	}
}

func (s *Server) Start(port string) error {
	return s.Run(port)
}

func (s *Server) SetService(service resources.Service) {
	s.service = service
}

func (s *Server) SetRouters(version string) {
	apiV1Group := s.Group(fmt.Sprintf("/api/%s", version))
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
