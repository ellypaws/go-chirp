package models

type Follow struct {
    ID         int `json:"id"`
    FollowerID int `json:"follower_id"`
    FollowedID int `json:"followed_id"`
}
