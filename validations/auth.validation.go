package validations

type LoginStruct struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type RegisterStruct struct {
	UserName string `validate:"required,min=3"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
