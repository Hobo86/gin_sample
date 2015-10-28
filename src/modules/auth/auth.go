package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	DefaultKey  = "modules/auth"
	errorFormat = "[modules] ERROR! %s\n"
)

var (
	// RedirectUrl should be the relative URL for your login route
	RedirectUrl string = "/login"

	// RedirectParam is the query string parameter that will be set
	// with the page the user was trying to visit before they were
	// intercepted.
	RedirectParam string = "return_url"

	// SessionKey is the key containing the unique ID in your session
	SessionKey string = "AUTHUNIQUEID"
)

type User interface {
	// Return whether this user is logged in or not
	IsAuthenticated() bool

	// Set any flags or extra data that should be available
	Login()

	// Clear any sensitive data out of the user
	Logout()

	// Return the unique identifier of this user object
	UniqueId() interface{}

	// Populate this user object with values
	GetById(id interface{}) error
}

type auth struct {
	User
}

func Auth(newUser func() User) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		userId := s.Get(SessionKey)
		user := newUser()

		if userId != nil {
			err := user.GetById(userId)
			if err != nil {
				log.Printf("Login Error: %v\n", err)
			} else {
				user.Login()
			}
		} else {
			log.Printf("Login Error: No UserId")
		}

		auth := auth{user}
		c.Set(DefaultKey, auth)
		c.Next()
	}
}

// shortcut to get Auth
func Default(c *gin.Context) auth {
	return c.MustGet(DefaultKey).(auth)
}

// AuthenticateSession will mark the session and user object as authenticated. Then
// the Login() user function will be called. This function should be called after
// you have validated a user.
func AuthenticateSession(s sessions.Session, user User) error {
	user.Login()
	return UpdateUser(s, user)
}

func (a auth) LogoutTest(s sessions.Session) {
	a.User.Logout()
	s.Delete(SessionKey)
	s.Save()
}

// Logout will clear out the session and call the Logout() user function.
func Logout(s sessions.Session, user User) {
	user.Logout()
	s.Delete(SessionKey)
	s.Save()
}

// LoginRequired verifies that the current user is authenticated. Any routes that
// require a login should have this handler placed in the flow. If the user is not
// authenticated, they will be redirected to /login with the "next" get parameter
// set to the attempted URL.
func LoginRequired(c *gin.Context) {
	a := Default(c)
	if a.User.IsAuthenticated() == false {
		path := fmt.Sprintf("%s?%s=%s", RedirectUrl, RedirectParam, c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, path)
	}
}

// UpdateUser updates the User object stored in the session. This is useful incase a change
// is made to the user model that needs to persist across requests.
func UpdateUser(s sessions.Session, user User) error {
	s.Set(SessionKey, user.UniqueId())
	s.Save()
	return nil
}
