package models

type Tweet struct {
    ID        int    `json:"id"`
    UserID    int    `json:"user_id"`
    Content   string `json:"content"`
    CreatedAt string `json:"created_at"`
}