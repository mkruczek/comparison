package pointers

import (
	"comparasion/resources"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	*echo.Echo
	service *resources.Service
}

func NewServer() *Server {
	return &Server{
		Echo: echo.New(),
	}
}

func (s *Server) Start(port string) error {
	return s.Echo.Start(port)
}

func (s *Server) SetService(service resources.Service) {
	s.service = &service
}

func (s *Server) SetRouters(version string) {
	apiV1Group := s.Group(fmt.Sprintf("/api/%s", version))
	apiV1Group.POST("/resources", s.createResources)
	apiV1Group.GET("/resources", s.getAllResources)
	apiV1Group.GET("/resources/:id", s.getResources)
	apiV1Group.PUT("/resources", s.updateResources)
	apiV1Group.DELETE("/resources/:id", s.deleteResources)
}

func (s *Server) createResources(e echo.Context) error {

	var newResources resources.Resources

	if err := e.Bind(&newResources); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	result, err := s.service.Create(newResources)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, result)
}

func (s *Server) getResources(e echo.Context) error {

	id := e.Param("id")

	result, err := s.service.Get(uuid.MustParse(id))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, result)
}

func (s *Server) getAllResources(e echo.Context) error {

	result, err := s.service.GetAll()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, result)
}

func (s *Server) updateResources(e echo.Context) error {

	var newResources resources.Resources

	if err := e.Bind(&newResources); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	result, err := s.service.Update(newResources)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, result)
}

func (s *Server) deleteResources(e echo.Context) error {

	id := e.Param("id")

	err := s.service.Delete(uuid.MustParse(id))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusNoContent, nil)
}
