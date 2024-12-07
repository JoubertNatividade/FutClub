package requests

type PlayerRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	LastName string `json:"last_name" binding:"required,min=3,max=50"`
	Position string `json:"position" binding:"required,min=3,max=50"`
}
