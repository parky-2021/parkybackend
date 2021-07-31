package token

//because i will mock signin in the future while writing test cases, it will be defined in an interface:
type sigInInterface interface {
	SignIn(AuthDetails) (string, error)
}

type signInStruct struct{}

//let expose this interface:
var (
	Authorize sigInInterface = &signInStruct{}
)

func (si *signInStruct) SignIn(authD AuthDetails) (string, error) {
	token, err := CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}
