package database

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"database/sql"
	"fmt"
//	"strconv"
//	"errors"
//	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/domain/entities"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/utils"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func UpdateUserDatabase(UField entities.User, User string) error {
	fmt.Println(" Comienza Update User")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	//Declaramos la sentencia SQL para insertar el Producto.
	sentencia := "Update users SET "

	coma := ""

	if len(UField.UserFirstName) > 0 {
		coma = ","
		sentencia += "User_FirstName = '" + UField.UserFirstName + "'"
	}

	if len(UField.UserLastName) > 0{
		sentencia += coma + "User_LastName = '" + UField.UserLastName + "'"
	}

	//Terminamos la sentencia indicando el id del registro que se va actualizar
	sentencia += ", User_DateUpg = '" + utils.FechaMySQL() + "' WHERE User_UUID='"+ User + "'"

	//Ejecutamos la sentencia SQL
	_, err = secundary.Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update User > Ejecución Exitosa!")

	return nil

}