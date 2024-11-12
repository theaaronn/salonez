package models

import "fmt"

func (hall *Hall) PrecioString() string {
	return fmt.Sprintf("%.2f", hall.Precio)
}

func (loc *Location) String() string {
	return fmt.Sprintf("%s %s %s %d %s %s", loc.Calle, loc.Numero, loc.Colonia, loc.CP, loc.Ciudad, loc.Estado)
}

func (res *Reservation) PercentagePaidString() string {
	return fmt.Sprintf("%.2f", res.PercentagePaid)
}
