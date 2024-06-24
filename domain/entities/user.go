// Estructuras para category
package entities

//Estructura para Usuarios

type User struct {
	UserUUID       string  `json:"userUUID"`
	UserEmail      string  `json:"userEmail"`
	UserFirstName  string  `json:"userFirstName"`
	UserLastName   string  `json:"userLastName"`
	UserStatus     int     `json:"userStatus"`
	UserDateAdd    string  `json:"userDateAdd"`
	UserDateUpd    string  `json:"userDateUpd"`
}

type ListUser struct{
	TotalItems	int 	`json:"totalItems"`
	Date		[]User	`json:"data"`
}

type Product struct {
	ProdId         int     `json:"prodID"`
	ProdTitle      string  `json:"prodTitle"`
	ProdDescrition string  `json:"prodDescription"`
	ProdCreateAt   string  `json:"prodCreateAt"`
	ProdUpdated    string  `json:"prodUpdated"`
	ProdPrice      float64 `json:"prodPrice,omitempty"`
	ProdStock      int     `json:"prodStock"`
	ProdCategId    int     `json:"prodCategId"`
	ProdPath       string  `json:"prodPath"`
	ProdSearch     string  `json:"search,omitempty"`
	ProdCategPath  string  `json:"categPath,omitempty"`
}

//Nuevo modelo para Select Productos

type ProductResp struct {
	TotalItems int       `json:"totalItems"`
	Data       []Product `json:"data"` //Nuestro slice (arreglo) de productos
}