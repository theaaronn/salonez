package models

type User struct {
	Nombre string
	Correo string
}

type Usuario = User
type Propietario = User
type Admin = User

type Hall struct {
	Id             int32
	ImgsPath       []string
	Nombre         string
	Direccion      string
	Capacidad      string
	NumeroTelefono string
	Precio         float64
}

type Location struct {
	Calle   string
	Numero  string
	Colonia string
	CP      int
	Ciudad  string
	Estado  string
}

type Reservation struct {
	Date           string
	Time           string
	PercentagePaid float32
	Cancelled      bool
}

type Cancelation struct {
	Date string
	Time string
}
