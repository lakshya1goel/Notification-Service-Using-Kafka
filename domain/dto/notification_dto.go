package dto

type NotificationRequestDto struct {
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
}
