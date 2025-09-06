package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (r *RoleService) RoleGetAll() ([]*schemas.RoleResponse, error) {
	roles, err := r.RoleRepository.RoleGetAll()
	if err != nil {
		return nil, err
	}

	var rolesResponse []*schemas.RoleResponse
	_ = copier.Copy(&rolesResponse, &roles)

	return rolesResponse, nil
}
