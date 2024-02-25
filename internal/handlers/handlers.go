package handlers

import (
	"bbaktyke/lubetrack-analog.git/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	service service.Service
	Router  *gin.Engine
}

func New(service service.Service) Server {

	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	router.HandleMethodNotAllowed = true

	server := Server{
		Router:  router,
		service: service,
	}
	server.listen()

	return server
}

func (s Server) listen() {
	s.cors()

	oilAnalysis := s.Router.Group("/oil-analysis")
	{
		oilAnalysis.GET("/", s.GetOilAnalysis)
	}

	s.Router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, []ErrorResponse{{
			Description: "Unsupported method",
			Code:        http.StatusMethodNotAllowed,
		},
		})
	})

	s.Router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, []ErrorResponse{{
			Description: "Route not found",
			Code:        http.StatusNotFound,
		},
		})
	})

}

func (s Server) cors() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"PUT", "PATCH", "GET", "DELETE", "POST", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	s.Router.Use(cors.New(corsConfig))
}

type ErrorResponse struct {
	Code        int    `json:"code,omitempty" example:"400"`
	Error       error  `json:"error,omitempty"`
	Msg         string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

func (s Server) GetOilAnalysis(ctx *gin.Context) {

	ctx.Status(http.StatusNoContent)

}
