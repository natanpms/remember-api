package usecase

import (
	"github.com/natanpms/remember-api/model"
	"github.com/natanpms/remember-api/repository"
)

type MemoryUseCase struct {
	repository      repository.MemoryRepository
	groupRepository repository.MemoryGroupRepository
}

func NewMemoryUseCase(repo repository.MemoryRepository, groupRepo repository.MemoryGroupRepository) MemoryUseCase {
	return MemoryUseCase{
		repository:      repo,
		groupRepository: groupRepo,
	}
}

func (mu *MemoryUseCase) GetMemoriesByGroupID(groupID int, userID int) ([]model.Memory, error) {
	_, err := mu.groupRepository.GetGroupByID(groupID, userID)

	if err != nil {
		return []model.Memory{}, err
	}

	return mu.repository.GetMemoriesByGroupID(groupID, userID)
}

func (mu *MemoryUseCase) GetMemoryByID(id int, userID int) (model.Memory, error) {
	return mu.repository.GetMemoryByID(id, userID)
}

func (mu *MemoryUseCase) CreateMemory(memory model.Memory, userID int) (model.Memory, error) {
	_, err := mu.groupRepository.GetGroupByID(memory.GroupID, userID)

	if err != nil {
		return model.Memory{}, err
	}

	return mu.repository.CreateMemory(memory)
}

func (mu *MemoryUseCase) DeleteMemory(id int, userID int) error {
	return mu.repository.DeleteMemory(id, userID)
}
