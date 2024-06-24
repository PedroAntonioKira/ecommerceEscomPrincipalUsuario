package secundary
//Viene siendo lo que era antes awsgo.go

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InicializoAWS() {
	fmt.Println("Entramos a AWSGO")
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	fmt.Println("vamos a la mitad de AWSGO")

	if err != nil {
		panic("Error al cargar la configutations de .aws/config " + err.Error())
	}

	fmt.Println("Salimos de AWSGO")
}


/*
Este código es una función de inicialización para una aplicación basada en AWS Lambda 
que configura el AWS SDK para Go (v2) con la configuración predeterminada necesaria 
para interactuar con los servicios de AWS. La configuración se carga en la región us-east-1, 
y si ocurre algún error durante la carga de esta configuración, 
la función termina con un pánico, mostrando el error.

En el contexto de una Lambda de AWS, esta función se llamaría típicamente al inicio del 
manejo de la función Lambda para asegurarse de que todas las configuraciones necesarias 
estén listas antes de realizar cualquier operación con los servicios de AWS.
*/