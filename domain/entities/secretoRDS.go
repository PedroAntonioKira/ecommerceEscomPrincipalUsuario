package entities

//Estructura del secreto de Secret Mananger que recibiremos

type SecretRDSJson struct {
	Username            string `json:"username"` //Alt izq + 96
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}