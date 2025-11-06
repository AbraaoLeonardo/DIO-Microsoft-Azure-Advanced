package router

import (
	"log"
	"server/db"
	"server/internal/tarefas"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Panic("Nao foi possivel conectar no banco de dados")
	}
	repo := tarefas.NewRepository(db.GetDB())
	handler := tarefas.NewHandler(repo)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/tarefas", handler.GetTarefasHandler)
	r.GET("/tarefas/:id", handler.GetTarefasByIdHandler)
	r.POST("/tarefas", handler.CreateTarefasHandler)
	r.PUT("/tarefas/:id", handler.UpdateTarefasIdHandler)
	r.DELETE("/tarefas/:id", handler.DeleteTarefasHandler)
}
