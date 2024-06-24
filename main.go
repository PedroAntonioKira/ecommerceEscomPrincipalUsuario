package main

import (
	//Importaciones de go (vienen incluidas al instalar)
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"
	lambda02 "github.com/aws/aws-lambda-go/lambda"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/primary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/secundary"
	//"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/utils"
)

func main() {
	lambda02.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	fmt.Println("Aqui inicia el programa de Usuarios 02")

	secundary.InicializoAWS()

	if !ValidoParametros() {
		panic("Error en los parametros, debe enviar 'SecretName', 'UrlPrefix'")
	} else {
		fmt.Println("Se validaron correctamente los Parametros")
	}

	fmt.Println("Imprimiremos el evento completo")

	// Imprimir el evento completo para depuración
	eventBytes, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir el evento a JSON:", err)
	} else {
		fmt.Println("Evento recibido:")
		fmt.Println(string(eventBytes))
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.Path, os.Getenv("UrlPrefix"), "", -1)
	method := request.HTTPMethod
	body := request.Body
	header := request.Headers
	stage := request.RequestContext.Stage

	fmt.Println("Información importante:")
	fmt.Println("Path: " + path)
	fmt.Println("Method: " + method)
	fmt.Println("Body: " + body)
	fmt.Println("Stage: " + stage)
	fmt.Println("HEADERS: ")
	for key, value := range header {
		fmt.Printf("%s: %s\n", key, value)
	}

	//Leemos el secreto de RDS
	secundary.ReadSecret()

	status, message := primary.Manejadores(path, method, body, header, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}

	return res, nil

}

func ValidoParametros() bool {
	var traeParametro bool

	_, traeParametro = os.LookupEnv("SecretName")
	if !traeParametro {
		fmt.Println("Algo fallo en la parte de SecretName")
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		fmt.Println("Algo fallo en la parte de UrlPrefix")
		return traeParametro
	}

	return traeParametro
}