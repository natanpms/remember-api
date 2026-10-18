package repository

import (
	"database/sql"
	"fmt"

	"github.com/natanpms/remember-api/model"
)

type MemoryGroupRepository struct {
	connection *sql.DB
}

func NewMemoryGroupRepository(connection *sql.DB) MemoryGroupRepository {
	return MemoryGroupRepository{
		connection: connection,
	}
}

func (gr *MemoryGroupRepository) GetGroupsByUserID(userID int) ([]model.MemoryGroup, error) {
	query := `
		SELECT id, user_id, name, description, created_at
		FROM memory_groups
		WHERE user_id = $1
		ORDER BY created_at DESC`

	rows, err := gr.connection.Query(query, userID)

	if err != nil {
		fmt.Println(err)
		return []model.MemoryGroup{}, err
	}

	var groupList []model.MemoryGroup
	var groupObject model.MemoryGroup

	for rows.Next() {
		err = rows.Scan(
			&groupObject.ID,
			&groupObject.UserID,
			&groupObject.Name,
			&groupObject.Description,
			&groupObject.CreatedAt)

		if err != nil {
			fmt.Println(err)
			return []model.MemoryGroup{}, err
		}
		groupList = append(groupList, groupObject)
	}

	rows.Close()

	return groupList, nil
}

func (gr *MemoryGroupRepository) GetGroupByID(id int, userID int) (model.MemoryGroup, error) {
	const query = `
		SELECT id, user_id, name, description, created_at
		FROM memory_groups
		WHERE id = $1 AND user_id = $2`

	var group model.MemoryGroup

	err := gr.connection.QueryRow(query, id, userID).Scan(
		&group.ID,
		&group.UserID,
		&group.Name,
		&group.Description,
		&group.CreatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return model.MemoryGroup{}, err
	}

	return group, nil
}

func (gr *MemoryGroupRepository) CreateGroup(group model.MemoryGroup) (model.MemoryGroup, error) {
	const query = `
		INSERT INTO memory_groups (user_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	err := gr.connection.QueryRow(query, group.UserID, group.Name, group.Description).Scan(
		&group.ID,
		&group.CreatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return model.MemoryGroup{}, err
	}

	return group, nil
}

func (gr *MemoryGroupRepository) DeleteGroup(id int, userID int) error {
	const query = "DELETE FROM memory_groups WHERE id = $1 AND user_id = $2"

	result, err := gr.connection.Exec(query, id, userID)

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
