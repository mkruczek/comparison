package callback

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
	return &Server{
		Engine: gin.Default(),
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
	apiV1Group.POST("/resources", createResources(s.service))
	apiV1Group.GET("/resources", getAllResources(s.service))
	apiV1Group.GET("/resources/:id", getResources(s.service))
	apiV1Group.PUT("/resources", updateResources(s.service))
	apiV1Group.DELETE("/resources/:id", deleteResources(s.service))
}

func createResources(s resources.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newResources resources.Resources

		if err := c.BindJSON(&newResources); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}

		result, err := s.Create(newResources)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, result)
	}
}

func getResources(s resources.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		result, err := s.Get(uuid.MustParse(id))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, result)
	}
}

func getAllResources(s resources.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		result, err := s.GetAll()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, result)
	}
}

func updateResources(s resources.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newResources resources.Resources

		if err := c.BindJSON(&newResources); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}

		result, err := s.Update(newResources)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, result)
	}
}

func deleteResources(s resources.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := s.Delete(uuid.MustParse(id))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusNoContent, nil)
	}
}
