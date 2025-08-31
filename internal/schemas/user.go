package schemas

type UserResponse struct {
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

type UserContext struct {
	ID        uint    `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Address   *string `json:"address"`
	Cellphone *string `json:"cellphone"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	IsAdmin   bool    `json:"is_admin"`
	IsActive  bool    `json:"is_active"`
	RoleID    uint    `json:"role_id"`
	Role      string  `json:"role"`
}
