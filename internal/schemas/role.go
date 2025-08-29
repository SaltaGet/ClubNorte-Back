package schemas

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name" validate:"oneof=admin vendedor repositor"`
}