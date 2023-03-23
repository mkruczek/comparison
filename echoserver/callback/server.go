package callback

import (
	"comparasion/resources"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	*echo.Echo
	service resources.Service
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

func createResources(s resources.Service) func(echo.Context) error {
	return func(e echo.Context) error {
		var newResources resources.Resources

		if err := e.Bind(&newResources); err != nil {
			return e.JSON(http.StatusBadRequest, err)
		}

		result, err := s.Create(newResources)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, err)
		}

		return e.JSON(http.StatusOK, result)
	}
}

func getResources(s resources.Service) func(echo.Context) error {
	return func(e echo.Context) error {
		id := e.Param("id")

		result, err := s.Get(uuid.MustParse(id))
		if err != nil {
			return e.JSON(http.StatusInternalServerError, err)
		}

		return e.JSON(http.StatusOK, result)
	}
}

func getAllResources(s resources.Service) func(e echo.Context) error {
	return func(e echo.Context) error {
		result, err := s.GetAll()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, err)
		}

		return e.JSON(http.StatusOK, result)
	}
}

func updateResources(s resources.Service) func(e echo.Context) error {
	return func(e echo.Context) error {
		var newResources resources.Resources

		if err := e.Bind(&newResources); err != nil {
			return e.JSON(http.StatusBadRequest, err)
		}

		result, err := s.Update(newResources)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, err)
		}

		return e.JSON(http.StatusOK, result)
	}
}

func deleteResources(s resources.Service) func(e echo.Context) error {
	return func(e echo.Context) error {
		id := e.Param("id")

		err := s.Delete(uuid.MustParse(id))
		if err != nil {
			return e.JSON(http.StatusInternalServerError, err)
		}

		return e.JSON(http.StatusNoContent, nil)
	}
}
