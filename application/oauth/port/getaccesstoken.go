package port

type GetAccessTokenInput struct {
	OAuthToken    string
	OAuthVerifier string
}

type GetAccessTokenOutput struct {
	OAuthToken       string
	OAuthTokenSecret string
}

type GetAccessToken interface {
	Invoke(input GetAccessTokenInput) (GetAccessTokenOutput, error)
}
