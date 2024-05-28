package user

type CreateUser struct {
	Name     string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

type AuthUser struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

type CheckEmailUser struct {
	Email string `binding:"required,email"`
}
