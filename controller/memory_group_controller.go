package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/natanpms/remember-api/handlers"
	"github.com/natanpms/remember-api/model"
	"github.com/natanpms/remember-api/usecase"
)

type memoryGroupController struct {
	memoryGroupUseCase usecase.MemoryGroupUseCase
}

func NewMemoryGroupController(usecase usecase.MemoryGroupUseCase) memoryGroupController {
	return memoryGroupController{
		memoryGroupUseCase: usecase,
	}
}

// GetGroups
// @Summary Get all memory groups
// @Description Get a list of all memory groups of the authenticated user
// @Tags groups
// @Produce json
// @Success 200 {object} handlers.APIResponse
// @Failure 401 {object} handlers.APIResponse
// @Failure 500 {object} handlers.APIResponse
// @Security BearerAuth
// @Router /api/v1/groups [get]
func (gc *memoryGroupController) GetGroups(ctx *gin.Context) {
	userID := ctx.GetInt("userID")

	groups, err := gc.memoryGroupUseCase.GetGroups(userID)

	if err != nil {
		handlers.Error(ctx, 500, handlers.ErrCodeInternal, "Erro ao buscar grupos")
		return
	}

	handlers.Success(ctx, 200, "", groups)
}

// GetGroupByID
// @Summary Get a memory group by ID
// @Description Get a single memory group by its ID
// @Tags groups
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} handlers.APIResponse
// @Failure 400 {object} handlers.APIResponse
// @Failure 401 {object} handlers.APIResponse
// @Failure 404 {object} handlers.APIResponse
// @Security BearerAuth
// @Router /api/v1/groups/{id} [get]
func (gc *memoryGroupController) GetGroupByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handlers.InvalidRequest(ctx, "ID inválido")
		return
	}

	userID := ctx.GetInt("userID")

	group, err := gc.memoryGroupUseCase.GetGroupByID(id, userID)

	if err != nil {
		handlers.NotFound(ctx, "Grupo não encontrado")
		return
	}

	handlers.Success(ctx, 200, "", group)
}

// CreateGroup
// @Summary Create a memory group
// @Description Create a new memory group for the authenticated user
// @Tags groups
// @Accept json
// @Produce json
// @Param group body model.MemoryGroup true "Group payload"
// @Success 201 {object} handlers.APIResponse
// @Failure 400 {object} handlers.APIResponse
// @Failure 401 {object} handlers.APIResponse
// @Failure 500 {object} handlers.APIResponse
// @Security BearerAuth
// @Router /api/v1/groups [post]
func (gc *memoryGroupController) CreateGroup(ctx *gin.Context) {
	var group model.MemoryGroup

	err := ctx.BindJSON(&group)

	if err != nil {
		handlers.InvalidRequest(ctx, "Erro ao ler dados do grupo")
		return
	}

	userID := ctx.GetInt("userID")

	groupInserted, err := gc.memoryGroupUseCase.CreateGroup(group, userID)

	if err != nil {
		handlers.Error(ctx, 500, handlers.ErrCodeInternal, "Erro ao criar grupo")
		return
	}

	handlers.Success(ctx, 201, "Grupo criado com sucesso", groupInserted)
}

// DeleteGroup
// @Summary Delete a memory group
// @Description Delete a memory group by its ID (memories inside it are also deleted)
// @Tags groups
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} handlers.APIResponse
// @Failure 400 {object} handlers.APIResponse
// @Failure 401 {object} handlers.APIResponse
// @Failure 404 {object} handlers.APIResponse
// @Security BearerAuth
// @Router /api/v1/groups/{id} [delete]
func (gc *memoryGroupController) DeleteGroup(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handlers.InvalidRequest(ctx, "ID inválido")
		return
	}

	userID := ctx.GetInt("userID")

	err = gc.memoryGroupUseCase.DeleteGroup(id, userID)

	if err != nil {
		handlers.NotFound(ctx, "Grupo não encontrado")
		return
	}

	handlers.Success(ctx, 200, "Grupo excluído com sucesso", nil)
}
