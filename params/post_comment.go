package params

type Comment struct {
	Message string `json:"message"`
	PhotoId uint   `json:"photo_id"`
}

type UpdateComment struct {
	Message string `json:"message"`
}
