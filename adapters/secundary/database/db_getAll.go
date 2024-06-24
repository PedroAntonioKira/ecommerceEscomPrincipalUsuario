package database

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
	"strconv"
//	"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary/database"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func ListUserDatabase(Page int) (entities.ListUser, error) {

	fmt.Println("Comienza Select List Users")

	var lu entities.ListUser
	User := []entities.User{}

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return lu, err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	var offset int = (Page * 10) -10
	var sentencia string
	var sentenciaCount string = "SELECT count(*) as registros FROM users"

	sentencia = "SELECT * FROM users LIMIT 10"
	if offset > 0 {
		sentencia += " OFFSET " + strconv.Itoa(offset)
	}

	var rowsCount *sql.Rows

	rowsCount, err = secundary.Db.Query(sentenciaCount)

	if err != nil {
		return lu, err
	}

	defer rowsCount.Close()

	rowsCount.Next()

	var registros int
	rowsCount.Scan(&registros)
	lu.TotalItems= registros

	var rows *sql.Rows
	rows, err = secundary.Db.Query(sentencia)
	if( err != nil) {
		fmt.Println(err.Error())
		return lu, err
	}

	for rows.Next() {
		var u entities.User
		var firstName sql.NullString
		var lastName sql.NullString
		var dateUpg sql.NullString

		rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateAdd, &dateUpg)

		u.UserFirstName = firstName.String
		u.UserLastName = lastName.String
		u.UserDateUpd = dateUpg.String

		User = append(User, u)
	}
	
	fmt.Println("Select Users > Ejecución Exitosa!")

	lu.Date = User
	return lu, nil
}