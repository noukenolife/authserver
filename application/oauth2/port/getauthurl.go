package port

type GetAuthURLInput struct {
	Scopes []string
}

type GetAuthURLOutput struct {
	URL string
}

type GetAuthURL interface {
	Invoke(input GetAuthURLInput) (GetAuthURLOutput, error)
}
