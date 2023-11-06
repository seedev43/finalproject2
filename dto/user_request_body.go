package dto

type UserRegister struct {
	Username string `json:"username" valid:"required~Your username is required"`
	Email    string `json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `json:"password" valid:"required~Your password is required,minstringlength(6)~Minimum password length is 6 characters"`
	Age      int    `json:"age" valid:"required~Your age is required,range(8|99)~Minimum age 8 years"`
}

type UserLogin struct {
	Email    string `json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `json:"password" valid:"required~Your password is required,minstringlength(6)~Minimum password length is 6 characters"`
}

type UserUpdate struct {
	Email    string `json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Username string `json:"username" valid:"required~Your username is required"`
}
