package requests

type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type AddRoleRequest struct {
	UserID string `json:"user_id" binding:"required"`
	RoleID uint   `json:"role_id" binding:"required"`
}

type RemoveRoleRequest struct {
	UserID string `json:"user_id" binding:"required"`
	RoleID uint   `json:"role_id" binding:"required"`
}
