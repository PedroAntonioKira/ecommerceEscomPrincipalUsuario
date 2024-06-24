package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
//	"strconv"
	"fmt"
//	"strings"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/utils"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func GetUserQuery(body string, User string) (int, string) {
	//Actualizamos el User.
	_, encontrado := utils.UserExists(User)

	//Verificamos no exista un error al momento en que actualizamos el producto.
	if !encontrado {
		return 400, "No existe un usuario con es UUID '" + User + "' > " 
	}

	row, err := database.GetUserDatabase(User)
	fmt.Println(row)

	if err != nil {
		return 400, "Ocurrio un error al intentar realizar el Select del Usuario " + User + " > " + err.Error()
	}

	respJson, err := json.Marshal(row)

	if err != nil {
		return 500, "Error al formatear los datos del usuario como Json  "
	}

	return 200, string(respJson)
}