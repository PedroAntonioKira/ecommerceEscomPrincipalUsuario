package repositories

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/domain/queries_category"
)

func UpdateUserRepositories(body string, user string) (int, string) {
	status := 200
	response := "Vacio Update Category Repositories"

	status, response = queries_category.UpdateUserQuery(body, user)

	return status, response
}