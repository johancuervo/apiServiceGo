package domain

// AuthService define el contrato para la autenticación
type AuthService interface {
	GetToken() string
	SetToken(token string)
	IsTokenValid() bool
}
