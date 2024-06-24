package primary

import (
	//Importaciones de go (vienen incluidas al instalar)
	"fmt"
	"strconv"
    "encoding/json"
	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"
	//importaciones personalizadas (creadas desde cero)
    "github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/utils"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/domain/use_cases"

)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayProxyRequest) (int, string) {

	fmt.Println("Voy a procesar " + path + " > " + method)

	id := request.PathParameters["id"]

	fmt.Println("El ID tiene:")
	fmt.Println(id)
	fmt.Println("Se imprimio ID :")

	idn, _ := strconv.Atoi(id)

    fmt.Println("Mostramos IDN: ")
    fmt.Println(idn)
	fmt.Println("Mostramos METHOD: " + method)
	fmt.Println("Mostramos PATH: " + path)

    eventBytes, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir el evento a JSON:", err)
	} else {
		fmt.Println("Evento recibido 00002:")
        fmt.Println(request.Resource) 
		fmt.Println(string(eventBytes))
	}

    //validamos la autorización del token
	isOk, statusCode, user := utils.ValidoAuthorization(path, method, headers)

    fmt.Println("El IsOk: ")
	fmt.Println(isOk)

    if !isOk {
		return statusCode, user
	}
    fmt.Println("Vamos bien con cognito")
	
	fmt.Println("Path1: " + path[0:19])
	fmt.Println("Path2: " + path[16:20])

	//Validamos y analizamos que nos viene en el path
    switch path[16:20] {
	case "user":
		fmt.Println("Entramos a Users")
		return ProcesoUsers(body, path, method, user, id, request)
		//return ProcesoProduct(body, path, method, user, idn, request)
    default:
		fmt.Println("No es la ruta correcta: No es por categoria, Detección Sospechosa 02")
		//	return ProcesoOrder(body, path, method, user, idn, request)
	}
	return 400, "Method Invalid loquisimo 10"
}

func ProcesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayProxyRequest) (int, string) {

	fmt.Println("Si entramos A PUT de Users")

	if(request.Resource == "/ecommerceEscom/user/normal"){

		switch method{
		case "PUT":
			return use_cases.UpdateUserUC(body, user)
		case "GET":
			return use_cases.GetUserUC(body, user)
		}
		
	}else if (request.Resource == "/ecommerceEscom/user/admin"){

		switch method{
		case "GET":
			return use_cases.ListUserUC(body, user, request)
		}
	}
	

	fmt.Println("LA RUTA: request.Resource" )
	fmt.Println(request.Resource )
	return 400, "Method Invalid  dentro"
}

/*
func ProcesoProduct(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {
	//Validamos el metodo Que estamos Recibiendo
	pathParams01 := request.PathParameters["id"] //Alfanumericos
	pathParams02, _ := strconv.Atoi(pathParams01) //Numerico

	switch method {
	case "POST":
		fmt.Println("Si entramos A POST de Producto")
		return use_cases.AddProductUC(body, user)
		//return use_cases.AddCategoryUC(body, user)
	case "PUT":
		fmt.Println("Si entramos A PUT de Category")
		return use_cases.UpdateProductUC(body, user, pathParams02)
		//return use_cases.UpdateCategoryUC(body, user, pathParams02)
	case "DELETE":
		fmt.Println("Si entramos A DELETE de Category")
		return use_cases.DeleteProductUC(body, user, pathParams02)
		//return use_cases.DeleteCategoryUC(body, user, pathParams02)
		//return routers.DeleteCategory(body, user, id)
	case "GET":
		fmt.Println("Si entramos A GET de Category")
        if(request.Resource == "/ecommerceEscom/product"){
            fmt.Println("Se deben de traer todas las categorias")
			return use_cases.ListProductUC(body, request)
        }else if(request.Resource == "/ecommerceEscom/product/{id}"){
			id := request.PathParameters["id"]
            fmt.Println("Se debe de traer una categoria en especifico")
			fmt.Println(id)
			return use_cases.GetProductUC(body, request, pathParams02)
			//return use_cases.ListProductsUC(body, request)
        }else{
            fmt.Println("Algo esta mal con el Metodo Get No entra en ningun caso")
        }
		//return routers.SelectCategories(body, request)
	}
	return 400, "Method Invalid Para Categorias, revisar en el codigo"
}
	
*/

/*

func ProcesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayProxyRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoProducts(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {

	//Validamos el metodo Que estamos Recibiendo
	switch method {
	case "POST":
		fmt.Println("Si entramos A POST de Product")
		return routers.InsertProduct(body, user)
	case "PUT":
		fmt.Println("Si entramos A PUT de Product")
		return routers.UpdateProduct(body, user, id)
	case "DELETE":
		fmt.Println("Si entramos A DELETE de Product")
		return routers.DeleteProduct(user, id)
	case "GET":
		fmt.Println("Si entramos A GET de Product")
		return routers.SelectProduct(request)
	}

	return 400, "Method Invalid"
}

func ProcesoStock(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoAddress(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {

	//Validamos el metodo Que estamos Recibiendo
	switch method {
	case "POST":
		fmt.Println("Si entramos A POST de Category")
		return routers.InsertCategory(body, user)
	case "PUT":
		fmt.Println("Si entramos A PUT de Category")
		return routers.UpdateCategory(body, user, id)
	case "DELETE":
		fmt.Println("Si entramos A DELETE de Category")
		return routers.DeleteCategory(body, user, id)
	case "GET":
		fmt.Println("Si entramos A GET de Category")
		return routers.SelectCategories(body, request)
	}
	return 400, "Method Invalid Para Categorias, revisar en el codigo"
}

func ProcesoOrder(body string, path string, method string, user string, id int, request events.APIGatewayProxyRequest) (int, string) {

	return 400, "Method Invalid"
}

*/


//N EDRP

// GCP

//EMILIO, ERICKA SOFIA