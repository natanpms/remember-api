package repository

import (
	"database/sql"
	"fmt"

	"github.com/natanpms/remember-api/model"
)

type MemoryRepository struct {
	connection *sql.DB
}

func NewMemoryRepository(connection *sql.DB) MemoryRepository {
	return MemoryRepository{
		connection: connection,
	}
}

func (mr *MemoryRepository) GetMemoriesByGroupID(groupID int, userID int) ([]model.Memory, error) {
	query := `
		SELECT m.id, m.group_id, m.title, m.description, m.created_at
		FROM memories m
		JOIN memory_groups g ON g.id = m.group_id
		WHERE m.group_id = $1 AND g.user_id = $2
		ORDER BY m.created_at DESC`

	rows, err := mr.connection.Query(query, groupID, userID)

	if err != nil {
		fmt.Println(err)
		return []model.Memory{}, err
	}

	var memoryList []model.Memory
	var memoryObject model.Memory

	for rows.Next() {
		err = rows.Scan(
			&memoryObject.ID,
			&memoryObject.GroupID,
			&memoryObject.Title,
			&memoryObject.Description,
			&memoryObject.CreatedAt)

		if err != nil {
			fmt.Println(err)
			return []model.Memory{}, err
		}
		memoryList = append(memoryList, memoryObject)
	}

	rows.Close()

	return memoryList, nil
}

func (mr *MemoryRepository) GetMemoryByID(id int, userID int) (model.Memory, error) {
	const query = `
		SELECT m.id, m.group_id, m.title, m.description, m.created_at
		FROM memories m
		JOIN memory_groups g ON g.id = m.group_id
		WHERE m.id = $1 AND g.user_id = $2`

	var memory model.Memory

	err := mr.connection.QueryRow(query, id, userID).Scan(
		&memory.ID,
		&memory.GroupID,
		&memory.Title,
		&memory.Description,
		&memory.CreatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return model.Memory{}, err
	}

	return memory, nil
}

func (mr *MemoryRepository) CreateMemory(memory model.Memory) (model.Memory, error) {
	const query = `
		INSERT INTO memories (group_id, title, description)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	err := mr.connection.QueryRow(query, memory.GroupID, memory.Title, memory.Description).Scan(
		&memory.ID,
		&memory.CreatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return model.Memory{}, err
	}

	return memory, nil
}

func (mr *MemoryRepository) DeleteMemory(id int, userID int) error {
	const query = `
		DELETE FROM memories
		WHERE id = $1 AND group_id IN (
			SELECT id FROM memory_groups WHERE user_id = $2
		)`

	result, err := mr.connection.Exec(query, id, userID)

	if err != nil {
		fmt.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
