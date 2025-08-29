package schemas

type UserResponseToken struct {
	ID        uint         `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Address   *string      `json:"address"`
	Cellphone *string      `json:"cellphone"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
	IsAdmin   bool         `json:"is_admin"`
	Role      RoleResponse `json:"role"`
}
