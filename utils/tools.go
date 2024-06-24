package utils

import (

	//Importaciones de go (vienen incluidas al instalar)
	"fmt"
	"strconv"
	"strings"
	"time"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalUsuario/adapters/secundary"
)

func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
}

func EscapeString(t string) string {
	desc := strings.ReplaceAll(t, "'", "")
	desc = strings.ReplaceAll(t, "'\"", "")
	return desc
}

func ArmoSentencia(s string, fieldName string, typeField string, ValueN int, ValueF float64, ValueS string) string {
	if (typeField == "S" && len(ValueS) == 0) ||
		(typeField == "F" && ValueF == 0) ||
		(typeField == "N" && ValueN == 0) {
		return s
	}

	if !strings.HasSuffix(s, "SET ") {
		s += ", "
	}

	switch typeField {
	case "S":
		s += fieldName + " = '" + EscapeString(ValueS) + "'"
	case "N":
		s += fieldName + " = " + strconv.Itoa(ValueN)
	case "F":
		s += fieldName + " = " + strconv.FormatFloat(ValueF, 'e', -1, 64)
	}

	return s

}

func UserExists(UserUUID string) (error, bool){
	fmt.Println("Comienza")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err, false
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	var sentencia string

	//Armamos sentencia sql
	sentencia = "SELECT 1 FROM users WHERE User_UUID='" + UserUUID + "'"

	//Imprimimos la sentencia SQL
	fmt.Println(sentencia)

	rows, err := secundary.Db.Query(sentencia)

	if err != nil {
		return err, false
	}

	var valor string

	rows.Next()
	rows.Scan(&valor)

	fmt.Println("UserExists > Ejecución exitosa - valor devuelto " + valor)

	if valor == "1"{
		return nil, true
	}

	return nil, false
}
