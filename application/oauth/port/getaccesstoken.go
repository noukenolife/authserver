package port

type GetAccessTokenInput struct {
	Code string
}

type GetAccessTokenOutput struct {
	AccessToken string
}

type GetAccessToken interface {
	Invoke(input GetAccessTokenInput) (GetAccessTokenOutput, error)
}
