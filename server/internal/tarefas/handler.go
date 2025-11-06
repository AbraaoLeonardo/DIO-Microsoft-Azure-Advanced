package tarefas

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// GET /tarefas
func (h *Handler) GetTarefasHandler(c *gin.Context) {
	tarefas, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tarefas)
}

// GET /tarefas/:id
func (h *Handler) GetTarefasByIdHandler(c *gin.Context) {
	id := c.Param("id")
	tarefa, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if tarefa == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tarefa não encontrada"})
		return
	}
	c.JSON(http.StatusOK, tarefa)
}

// POST /tarefas
func (h *Handler) CreateTarefasHandler(c *gin.Context) {
	var t Tarefa
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTarefa, err := h.repo.Create(c.Request.Context(), &t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTarefa)
}

// PUT /tarefas/:id
func (h *Handler) UpdateTarefasIdHandler(c *gin.Context) {
	id := c.Param("id")
	var t Tarefa
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Update(c.Request.Context(), id, &t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "tarefa atualizada com sucesso"})
}

// DELETE /tarefas/:id
func (h *Handler) DeleteTarefasHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "tarefa deletada com sucesso"})
}
