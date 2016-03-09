package models

import (
	"time"

	"modules/auth"
	"modules/log"
)

func (m model) GetUserById(id uint64) *User {
	user := User{}
	if err := m.db.Where("id = ?", id).First(&user).Error; err != nil {
		log.DebugPrint("GetUserById error: %v", err)
		return nil
	}

	return &user
}

func (m model) GetUserByNicknamePwd(nickname string, pwd string) *User {
	user := User{}
	if err := m.db.Where("nickname = ? AND password = ?", nickname, pwd).First(&user).Error; err != nil {
		log.DebugPrint("GetUserByNicknamePwd error: %v", err)
		return nil
	}
	return &user
}

func (m model) AddUserWithNicknamePwd(nickname string, pwd string) *User {
	user := User{Nickname: nickname, Password: pwd, Birthday: time.Now()}

	if err := m.db.Create(&user).Error; err != nil {
		return nil
	}
	return &user
}

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Nickname  string    `form:"nickname" json:"nickname,omitempty"`
	Password  string    `form:"password" json:"-"`
	Gender    int64     `json:"gender,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
	CreatedAt time.Time `gorm:"column:created_time" json:"created_time,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_time" json:"updated_time,omitempty"`

	authenticated bool `form:"-" db:"-" json:"-"`
}

// GetAnonymousUser should generate an anonymous user model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() auth.User {
	return &User{}
}

func (u User) TableName() string {
	return "user"
}

// Login will preform any actions that are required to make a user model
// officially authenticated.
func (u *User) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	u.authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user.
func (u *User) Logout() {
	// Remove from logged-in user's list
	// etc ...
	u.authenticated = false
}

func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) UniqueId() interface{} {
	return u.Id
}

// GetById will populate a user object from a database model with
// a matching id.
func (u *User) GetById(id interface{}) error {
	u.Id = id.(uint64)
	// err := dbmap.SelectOne(u, "SELECT * FROM users WHERE id = $1", id)
	// if err != nil {
	// 	return err
	// }

	return nil
}
