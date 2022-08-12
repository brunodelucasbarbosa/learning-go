package userRepository

type UserInput struct {
	ID    int    `json:"-"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Taxid string `json:"taxid"` //cpf
	Password string `json:"password"` 
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Taxid string `json:"taxid"` //cpf
}