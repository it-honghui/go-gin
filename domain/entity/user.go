package entity

import (
	"encoding/json"
	"go-gin/domain/embed"
	"go-gin/domain/enum"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Username string    `gorm:"uniqueIndex:idx_u;size:20;not null" json:"username"`
	Password string    `gorm:"size:90;not null" json:"password,omitempty"`
	Role     enum.Role `gorm:"size:20;not null" json:"role"`
}

func (User) TableName() string {
	return "user"
}

func (t User) MarshalJSON() ([]byte, error) {
	type TmpJSON User
	return json.Marshal(&struct {
		TmpJSON
		CreatedAt embed.DateTime `json:"created_at"`
		UpdatedAt embed.DateTime `json:"updated_at"`
	}{
		TmpJSON:   (TmpJSON)(t),
		CreatedAt: embed.DateTime(t.CreatedAt),
		UpdatedAt: embed.DateTime(t.UpdatedAt),
	})
}
