package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
//	"strconv"
//	"fmt"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/secundary/database"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/utils"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func UpdateUserQuery(body string, User string) (int, string) {
	//creamos la variable de la estructura que almacenará los datos del usuario
	var t entities.User

	//Decodificamos el json que nos mandan en el endpoint en la estructura del producto para poder guardarla.
	err := json.Unmarshal([]byte(body), &t)

	//Verificamos que no tengamos un error al decodificar la información en la estructura.
	if err != nil {
		return 400, "Error en los datos recibidos con el error: " + err.Error()
	}

	if len(t.UserFirstName) == 0 && len(t.UserLastName) == 0 {
		return 400, "Se debe especificar el Nombre (FirstName) o el apellido (LastName) del usuario"
	}

	//Actualizamos el User.
	_, encontrado := utils.UserExists(User)

	//Verificamos no exista un error al momento en que actualizamos el producto.
	if !encontrado {
		return 400, "No existe un usuario con es UUID '" + User + "' > " 
	}

	err = database.UpdateUserDatabase(t, User)

	//Validamos que no exista algun error
	if err != nil {
		return 400, "Ocurrio un error al intentar realizar la actualización del usuario " + User + " > " + err.Error() 
	}

	return 200, "Update User OK User"

}
