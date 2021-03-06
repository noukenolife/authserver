package port

type GetAuthURLInput struct {
}

type GetAuthURLOutput struct {
	URL string
}

type GetAuthURL interface {
	Invoke(input GetAuthURLInput) (GetAuthURLOutput, error)
}
