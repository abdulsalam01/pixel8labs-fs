package entity

type Account struct {
	ID           int64  `json:"id"`
	AvatarUrl    string `json:"avatar_url"`
	Name         string `json:"name"`
	Username     string `json:"login"`
	Email        string `json:"email"`
	Bio          string `json:"bio"`
	Follower     int64  `json:"followers"`
	Following    int64  `json:"following"`
	FollowersUrl string `json:"followers_url"`
}

type Email struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Visibility string `json:"visibility"`
}
