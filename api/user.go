package api

// UserController is the user controller interface
type UserController interface {
	SignIn(*UserSignInRequest) (*UserSignInResponse, error)
	SignOut(*UserSignOutRequest) (*UserSignOutResponse, error)
	UserChangePassword(*UserChangePasswordRequest) (*UserChangePasswordResponse, error)
	Get(int) (*User, error)
}

// UserSignInRequest is the sign in request for user controller interface
type UserSignInRequest struct {
	Username string
	Password string
}

// UserSignInRequest is the sigin in  response for user controller interface
type UserSignInResponse struct {
	Token  string
	UserID int
}

// UserSignOutRequest is the sign out request for user controller interface
type UserSignOutRequest struct {
	Token string
}

// UserSignOutResponse is the
type UserSignOutResponse struct { 
}

// UserChangePasswordRequest is the change password request input for UserPassword Method
type UserChangePasswordRequest struct {
	OldPassword string
	NewPasseord string
}

// UserChangePasswordResponse is the change password response output for UserPassword Method
type UserChangePasswordResponse struct {
}

// User is user struct
type User struct {
	Username string
	Secret   string
}
