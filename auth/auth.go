package auth

type Auth struct {
	UserDB string
	Keys   map[string]string
}

func (a *Auth) saltHash() (err error) {}

func (a *Auth) SignUp(user User) (err error) {
	return nil
}

func (a *Auth) Login() (jwt []byte, err error) {}

func (a *Auth) Verify() (err error) {}
