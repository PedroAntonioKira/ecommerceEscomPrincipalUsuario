package database

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
//	"strconv"
//	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/domain/entities"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func GetUserDatabase(UserId string) (entities.User, error) {

	fmt.Println("Comienza SelectUser")

	//var Resp entities.ProductResp
	User := entities.User{} // Con esto devolveremos una colección de estructuras

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return User, err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	var sentencia string

	sentencia = "SELECT * FROM users WHERE User_UUID = '" + UserId + "'"

	//Imprimimos la sentencia SQL
	fmt.Println(sentencia)

	var rows *sql.Rows

	rows, err = secundary.Db.Query(sentencia)

	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return User, err
	}

	rows.Next()

	var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullString

	rows.Scan(&User.UserUUID, &User.UserEmail, &firstName, &lastName, &User.UserStatus, &User.UserDateAdd, &dateUpg)

	User.UserFirstName = firstName.String
	User.UserLastName = lastName.String
	User.UserDateUpd = dateUpg.String

	fmt.Println(" Select Product > Ejecución Exitosa !")

	return User, nil
}