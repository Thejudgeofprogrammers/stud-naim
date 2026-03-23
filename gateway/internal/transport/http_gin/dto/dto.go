package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=student employer"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type StudentResponse struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Skills  []string `json:"skills"`
	About   string   `json:"about"`
	Resume  string   `json:"resume_url"`
}

type EmployerResponse struct {
	ID             string   `json:"id"`
	CompanyName    string   `json:"company_name"`
	Description    string   `json:"description"`
	Representative string   `json:"representative"`
	Vacancies      []string `json:"vacancies"`
}

type UpdateStudentRequest struct {
	Name   string   `json:"name"`
	Skills []string `json:"skills"`
	About  string   `json:"about"`
}

type UpdateEmployerRequest struct {
	CompanyName    string `json:"company_name"`
	Description    string `json:"description"`
	Representative string `json:"representative"`
}
