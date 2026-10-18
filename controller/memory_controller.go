package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/natanpms/remember-api/handlers"
	"github.com/natanpms/remember-api/model"
	"github.com/natanpms/remember-api/usecase"
)

type memoryController struct {
	memoryUseCase usecase.MemoryUseCase
}

func NewMemoryController(usecase usecase.MemoryUseCase) memoryController {
	return memoryController{
		memoryUseCase: usecase,
	}
}

// GetMemoriesByGroup
// @Summary Get all memories of a group
// @Description Get a list of all memories of a group
// @Tags memories
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} handlers.APIResponse
// @Failure 400 {object} handlers.APIResponse
// @Failure 401 {object} handlers.APIResponse
// @Failure 404 {object} handlers.APIResponse
// @Security BearerAuth
// @Router /api/v1/groups/{id}/memories [get]
func (mc *memoryController) GetMemoriesByGroup(ctx *gin.Context) {
	groupID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handlers.InvalidRequest(ctx, "ID inválido")
		return
	}

	userID := ctx.GetInt("userID")

	memories, err := mc.memoryUseCase.GetMemoriesByGroupID(groupID, userID)

	if err != nil {
		handlers.NotFound(ctx, "Grupo não encontrado")
		return
	}

	handlers.Success(ctx, 200, "", memories)
}

// GetMemoryByID
// @Summary Get a memory by ID
// @Description Get a single memory by its ID
// @Tags memories
// @Produce json
// @Param id path int true "Memory ID"
// @Success 200 {object} handlers.APIResponse
// @Failure 400 {object} handlers.APIResponse
// @Failure 401 {object} handlers.APIResponse
// @Failure 404 {object} handlers.APIResponse
// @Security BearerAuth
// @Router /api/v1/memories/{id} [get]
func (mc *memoryController) GetMemoryByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handlers.InvalidRequest(ctx, "ID inválido")
		return
	}

	userID := ctx.GetInt("userID")

	memory, err := mc.memoryUseCase.GetMemoryByID(id, userID)

	if err != nil {
		handlers.NotFound(ctx, "Memória não encontrada")
		return
	}

	handlers.Success(ctx, 200, "", memory)
}

// CreateMemory
// @Summary Create a memory
// @Description Create a new memory or idea inside a group
// @Tags memories
// @Accept json
// @Produce json
// @Param id path int true "Group ID"
// @Param memory body model.Memory true "Memory payload"
// @Success 201 {object} handlers.APIResponse
// @Failure 400 {object} handlers.APIResponse
// @Failure 401 {object} handlers.APIResponse
// @Failure 404 {object} handlers.APIResponse
// @Failure 500 {object} handlers.APIResponse
// @Security BearerAuth
// @Router /api/v1/groups/{id}/memories [post]
func (mc *memoryController) CreateMemory(ctx *gin.Context) {
	groupID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handlers.InvalidRequest(ctx, "ID inválido")
		return
	}

	var memory model.Memory

	err = ctx.BindJSON(&memory)

	if err != nil {
		handlers.InvalidRequest(ctx, "Erro ao ler dados da memória")
		return
	}

	userID := ctx.GetInt("userID")
	memory.GroupID = groupID

	memoryInserted, err := mc.memoryUseCase.CreateMemory(memory, userID)

	if err != nil {
		handlers.Error(ctx, 500, handlers.ErrCodeInternal, "Erro ao criar memória")
		return
	}

	handlers.Success(ctx, 201, "Memória criada com sucesso", memoryInserted)
}

// DeleteMemory
// @Summary Delete a memory
// @Description Delete a memory by its ID
// @Tags memories
// @Produce json
// @Param id path int true "Memory ID"
// @Success 200 {object} handlers.APIResponse
// @Failure 400 {object} handlers.APIResponse
// @Failure 401 {object} handlers.APIResponse
// @Failure 404 {object} handlers.APIResponse
// @Security BearerAuth
// @Router /api/v1/memories/{id} [delete]
func (mc *memoryController) DeleteMemory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handlers.InvalidRequest(ctx, "ID inválido")
		return
	}

	userID := ctx.GetInt("userID")

	err = mc.memoryUseCase.DeleteMemory(id, userID)

	if err != nil {
		handlers.NotFound(ctx, "Memória não encontrada")
		return
	}

	handlers.Success(ctx, 200, "Memória excluída com sucesso", nil)
}
