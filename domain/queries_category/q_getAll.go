package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"strconv"
//	"strings"
//	"fmt"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/domain/entities"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func ListUserQuery(body string, User string, request events.APIGatewayProxyRequest) (int, string) {

	var Page int
	if len(request.QueryStringParameters["page"]) == 0 {
		Page = 1
	}else{
		Page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}

	//Verificamos si User Is Admin
	isAdmin, msg := secundary.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	//Actualizamos el Producto.
	user, err2 := database.ListUserDatabase(Page)

	//Verificamos no exista un error al momento en que actualizamos el producto.
	if err2 != nil {
		return 400, "Ocurrio un error al intentar obtener la lista de usuarios" + err2.Error()
	}

	respJson, err2 := json.Marshal(user)

	if err2 != nil {
		return 400, "Ocurrio un error al formatear los datos de los usuarios como Json"
	}

	return 200, string(respJson)
}