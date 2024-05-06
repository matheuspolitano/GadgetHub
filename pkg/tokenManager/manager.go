package tokenManager

type Manager interface {
	GenerateToken(payloadParameters PayloadParameters) (string, *Payload, error)
	CheckToken(token string) (*Payload, error)
}

var _ Manager = (*JWTManager)(nil)
