package model

type User struct {
	// the id for this user.
	//
	// required: true
	ID int64 `json:"id"             gorm:"AUTO_INCREMENT;primary_key;column:user_id"`

	// Login is the username for this user.
	//
	// required: true
	Login string `json:"login"      gorm:"type:varchar(250);column:user_login"`

	// Token is the oauth2 token.
	Token string `json:"-"          gorm:"type:varchar(500);column:user_token"`
}

