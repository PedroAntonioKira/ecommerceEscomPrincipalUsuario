package use_cases

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/domain/repositories"
)

func UpdateUserUC(body string, user string) (int, string) {
	status := 200
	response := "Vacio Update Category Use Case"

	status, response = repositories.UpdateUserRepositories(body, user)

	return status, response
}