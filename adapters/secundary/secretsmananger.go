package secundary

import (

	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"fmt"

	//importaciones externas (descargadas)
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	//importaciones personalizadas (creadas desde cero)
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/awsgo"
	//"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/domain/entities"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func GetSecret(nombreSecret string) (entities.SecretRDSJson, error) {
	var datosSecret entities.SecretRDSJson
	fmt.Println(" > Pido Secreto " + nombreSecret)

	svc := secretsmanager.NewFromConfig(Cfg)

	clave, err := svc.GetSecretValue(Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	//Preguntamos si exisitio un error
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	//Trabajamos que clave que es la que tiene los valores

	//Unmarshal parsea el json codiicado que nos devuelve a la estructura nuestra
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)

	fmt.Println(" > Lectura Secret OK " + nombreSecret)
	return datosSecret, nil
}
