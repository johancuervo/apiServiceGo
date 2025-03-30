package application

import "github.com/johancuervo/apiServiceGo/src/domain"

// AuthUseCase maneja la lógica de aplicación para autenticación
type AuthUseCase struct {
	authService domain.AuthService
}

// NewAuthUseCase crea un nuevo caso de uso de autenticación
func NewAuthUseCase(authService domain.AuthService) *AuthUseCase {
	return &AuthUseCase{authService: authService}
}

// RefreshToken actualiza el token si es inválido
func (uc *AuthUseCase) RefreshToken(newToken string) {
	if !uc.authService.IsTokenValid() {
		uc.authService.SetToken(newToken)
	}
}

// GetToken obtiene el token
func (uc *AuthUseCase) GetToken() string {
	return uc.authService.GetToken()
}
