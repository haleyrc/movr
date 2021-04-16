package movr

import (
	"fmt"
	"time"

	"movr/maybe"

	"github.com/google/uuid"
)

type PromoCode struct {
	Code        string         `db:"code"`
	Description string         `db:"description"`
	Created     time.Time      `db:"creation_time"`
	Expires     time.Time      `db:"expiration_time"`
	Rules       PromoCodeRules `db:"rules"`
}

func (c PromoCode) String() string {
	return fmt.Sprintf(
		"<PromoCode(code='%s', description='%s', creation_time='%s', expiration_time='%s', rules='%s'>",
		c.Code,
		c.Description,
		c.Created,
		c.Expires,
		c.Rules,
	)
}

type PromoCodeRules map[string]interface{}

type Ride struct {
	ID           uuid.UUID     `db:"id"`
	City         string        `db:"city"`
	Rider        uuid.UUID     `db:"rider_id"`
	Vehicle      uuid.UUID     `db:"vehicle_id"`
	StartAddress string        `db:"start_address"`
	EndAddress   string        `db:"end_address"`
	StartTime    time.Time     `db:"start_time"`
	EndTime      maybe.Time    `db:"end_time"`
	Revenue      maybe.Float64 `db:"revenue"`
}

func (r Ride) String() string {
	return fmt.Sprintf(
		"<Ride(city='%s', id='%s', rider_id='%s', vehicle_id='%s')>",
		r.City,
		r.ID,
		r.Rider,
		r.Vehicle,
	)
}

type User struct {
	ID         uuid.UUID    `db:"id"`
	City       string       `db:"city"`
	Name       maybe.String `db:"name"`
	Address    maybe.String `db:"address"`
	CreditCard maybe.String `db:"credit_card"`
}

func (u User) String() string {
	return fmt.Sprintf(
		"<User(city='%s', id='%s', name='%s')>",
		u.City,
		u.ID,
		u.Name.Value,
	)
}

type UserPromoCode struct {
	City      string    `db:"city"`
	User      uuid.UUID `db:"user_id"`
	Code      string    `db:"code"`
	Timestamp time.Time `db:"timestamp"`
	Usage     int64     `db:"usage"`
}

func (c UserPromoCode) String() string {
	return fmt.Sprintf(
		"<UserPromoCode(city='%s', user_id='%s', code='%s', timestamp='%s'>",
		c.City,
		c.User,
		c.Code,
		c.Timestamp,
	)
}

type Vehicle struct {
	ID              uuid.UUID       `db:"id"`
	City            string          `db:"city"`
	Type            maybe.String    `db:"type"`
	Owner           uuid.UUID       `db:"owner_id"`
	Created         time.Time       `db:"creation_time"`
	Status          string          `db:"status"`
	CurrentLocation string          `db:"current_location"`
	Metadata        VehicleMetadata `db:"ext"`
}

func (v Vehicle) String() string {
	return fmt.Sprintf(
		"<Vehicle(city='%s', id='%s', type='%s', status='%s', ext='%s')>",
		v.City,
		v.ID,
		v.Type.Value,
		v.Status,
		v.Metadata,
	)
}

type VehicleLocationHistory struct {
	City      string    `db:"city"`
	Ride      uuid.UUID `db:"ride_id"`
	Timestamp time.Time `db:"timestamp"`
	Latitude  float64   `db:"lat"`
	Longitude float64   `db:"long"`
}

func (h VehicleLocationHistory) String() string {
	return fmt.Sprintf(
		"<VehicleLocationHistory(city='%s', ride_id='%s', timestamp='%s', lat='%f', long='%f')>",
		h.City,
		h.Ride,
		h.Timestamp,
		h.Latitude,
		h.Longitude,
	)
}

type VehicleMetadata map[string]interface{}
