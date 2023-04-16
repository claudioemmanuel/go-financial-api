package dtos

/**
* UserDTO is used to receive the user request body
 */
type UserDTO struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

/**
* AccountDTO is used to receive the account request body
 */
type AccountDTO struct {
	OwnerID uint `form:"owner_id" json:"owner_id" binding:"required"`
	Balance int `form:"balance" json:"balance" binding:"required"`
}

/**
* LoginDTO is used to receive the login request body
 */
type LoginDTO struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
