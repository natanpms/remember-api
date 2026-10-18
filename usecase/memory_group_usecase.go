package usecase

import (
	"github.com/natanpms/remember-api/model"
	"github.com/natanpms/remember-api/repository"
)

type MemoryGroupUseCase struct {
	repository repository.MemoryGroupRepository
}

func NewMemoryGroupUseCase(repo repository.MemoryGroupRepository) MemoryGroupUseCase {
	return MemoryGroupUseCase{
		repository: repo,
	}
}

func (gu *MemoryGroupUseCase) GetGroups(userID int) ([]model.MemoryGroup, error) {
	return gu.repository.GetGroupsByUserID(userID)
}

func (gu *MemoryGroupUseCase) GetGroupByID(id int, userID int) (model.MemoryGroup, error) {
	return gu.repository.GetGroupByID(id, userID)
}

func (gu *MemoryGroupUseCase) CreateGroup(group model.MemoryGroup, userID int) (model.MemoryGroup, error) {
	group.UserID = userID
	return gu.repository.CreateGroup(group)
}

func (gu *MemoryGroupUseCase) DeleteGroup(id int, userID int) error {
	return gu.repository.DeleteGroup(id, userID)
}
