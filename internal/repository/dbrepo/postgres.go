package dbrepo

import (
	"context"
	"time"

	"github.com/JeanCntrs/bookings/internal/models"
)

func (pdb *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (pdb *postgresDBRepo) InsertReservation(r models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int
	stmt := `insert into reservations (first_name, last_name, email, phone, start_date,
			end_date, room_id, created_at, updated_at)
			values($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := pdb.DB.QueryRowContext(ctx, stmt,
		r.FirstName,
		r.LastName,
		r.Email,
		r.Phone,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (pdb *postgresDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into room_restrictions (start_date, end_date, room_id, reservation_id,
			created_at, updated_at, restriction_id)
			values($1, $2, $3, $4, $5, $6, $7)`

	_, err := pdb.DB.ExecContext(ctx, stmt,
		rr.StartDate,
		rr.EndDate,
		rr.RoomID,
		rr.ReservationID,
		time.Now(),
		time.Now(),
		rr.RestrictionID,
	)

	if err != nil {
		return err
	}

	return nil
}
