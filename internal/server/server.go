package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetUpServer(router *gin.Engine, port string) {
	log.Info().Msg("Servers starts")
	if err := router.Run(port); err != nil {
		log.Error().Msg("Failed to start server")
	}
}
