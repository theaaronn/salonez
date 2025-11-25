package db

import (
	"time"
)

type ReservationData struct {
	Id             int
	HallId         int
	HallName       string
	UserId         int
	UserName       string
	UserEmail      string
	Date           string
	Time           string
	PercentagePaid float32
	Cancelled      bool
}

func GetAllReservations() ([]ReservationData, error) {
	db := GetDb()
	defer db.Close()

	query := `
		SELECT 
			r.id,
			r.hall_id,
			h.nombre as hall_name,
			r.user_id,
			u.nombre as user_name,
			u.correo as user_email,
			r.date,
			r.time,
			r.percentage_paid,
			r.cancelled
		FROM reservations r
		JOIN halls h ON r.hall_id = h.id
		JOIN users u ON r.user_id = u.id
		ORDER BY r.date DESC, r.time DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reservations := make([]ReservationData, 0)
	for rows.Next() {
		var r ReservationData
		var date time.Time
		var timeVal time.Time

		err := rows.Scan(&r.Id, &r.HallId, &r.HallName, &r.UserId, &r.UserName, &r.UserEmail,
			&date, &timeVal, &r.PercentagePaid, &r.Cancelled)
		if err != nil {
			return nil, err
		}

		r.Date = date.Format("2006-01-02")
		r.Time = timeVal.Format("15:04")
		reservations = append(reservations, r)
	}

	return reservations, nil
}

func GetReservationsByUser(userId int) ([]ReservationData, error) {
	db := GetDb()
	defer db.Close()

	query := `
		SELECT 
			r.id,
			r.hall_id,
			h.nombre as hall_name,
			r.user_id,
			u.nombre as user_name,
			u.correo as user_email,
			r.date,
			r.time,
			r.percentage_paid,
			r.cancelled
		FROM reservations r
		JOIN halls h ON r.hall_id = h.id
		JOIN users u ON r.user_id = u.id
		WHERE r.user_id = $1
		ORDER BY r.date DESC, r.time DESC
	`

	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reservations := make([]ReservationData, 0)
	for rows.Next() {
		var r ReservationData
		var date time.Time
		var timeVal time.Time

		err := rows.Scan(&r.Id, &r.HallId, &r.HallName, &r.UserId, &r.UserName, &r.UserEmail,
			&date, &timeVal, &r.PercentagePaid, &r.Cancelled)
		if err != nil {
			return nil, err
		}

		r.Date = date.Format("2006-01-02")
		r.Time = timeVal.Format("15:04")
		reservations = append(reservations, r)
	}

	return reservations, nil
}

func GetReservationsByHall(hallId int) ([]ReservationData, error) {
	db := GetDb()
	defer db.Close()

	query := `
		SELECT 
			r.id,
			r.hall_id,
			h.nombre as hall_name,
			r.user_id,
			u.nombre as user_name,
			u.correo as user_email,
			r.date,
			r.time,
			r.percentage_paid,
			r.cancelled
		FROM reservations r
		JOIN halls h ON r.hall_id = h.id
		JOIN users u ON r.user_id = u.id
		WHERE r.hall_id = $1
		ORDER BY r.date DESC, r.time DESC
	`

	rows, err := db.Query(query, hallId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reservations := make([]ReservationData, 0)
	for rows.Next() {
		var r ReservationData
		var date time.Time
		var timeVal time.Time

		err := rows.Scan(&r.Id, &r.HallId, &r.HallName, &r.UserId, &r.UserName, &r.UserEmail,
			&date, &timeVal, &r.PercentagePaid, &r.Cancelled)
		if err != nil {
			return nil, err
		}

		r.Date = date.Format("2006-01-02")
		r.Time = timeVal.Format("15:04")
		reservations = append(reservations, r)
	}

	return reservations, nil
}

func CreateReservation(hallId, userId int, date, timeStr string, percentage float64) error {
	db := GetDb()
	defer db.Close()

	query := `
		INSERT INTO reservations (hall_id, user_id, date, time, percentage_paid, cancelled)
		VALUES ($1, $2, $3, $4, $5, false)
	`

	_, err := db.Exec(query, hallId, userId, date, timeStr, percentage)
	return err
}

func CancelReservation(reservationId int) error {
	db := GetDb()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Marcar reservación como cancelada
	updateQuery := `UPDATE reservations SET cancelled = true WHERE id = $1`
	_, err = tx.Exec(updateQuery, reservationId)
	if err != nil {
		return err
	}

	// Crear registro de cancelación
	insertQuery := `
		INSERT INTO cancelations (reservation_id, date, time)
		VALUES ($1, CURRENT_DATE, CURRENT_TIME)
	`
	_, err = tx.Exec(insertQuery, reservationId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func ProcessPayment(reservationId int, additionalPercentage float64) error {
	db := GetDb()
	defer db.Close()

	query := `
		UPDATE reservations 
		SET percentage_paid = LEAST(percentage_paid + $1, 100),
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $2 AND cancelled = false
	`

	_, err := db.Exec(query, additionalPercentage, reservationId)
	return err
}

func GetUserIdByEmail(email string) (int, error) {
	db := GetDb()
	defer db.Close()

	var id int
	err := db.QueryRow("SELECT id FROM users WHERE correo = $1", email).Scan(&id)
	return id, err
}
