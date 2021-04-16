package movr

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type RideFilter struct {
	Limit  int
	Offset int
}

type PromoCodeFilter struct {
	Limit  int
	Offset int
}

type UserFilter struct {
	Limit  int
	Offset int
}

type VehicleFilter struct {
	Limit  int
	Offset int
}

type AddUserRequest struct {
	City             string
	Name             string
	Address          string
	CreditCardNumber string
}

type AddUserResponse struct {
	User *User
}

type AddVehicleRequest struct {
	City            string
	Owner           uuid.UUID
	CurrentLocation string
	Type            string
	Metadata        VehicleMetadata
	Status          string
}

type AddVehicleResponse struct {
	Vehicle *Vehicle
}

type ApplyPromoCodeRequest struct {
	City      string
	User      uuid.UUID
	PromoCode string
}

type ApplyPromoCodeResponse struct {
	Code *UserPromoCode
}

type CreatePromoCodeRequest struct {
	Code        string
	Description string
	Expires     time.Time
	Rules       PromoCodeRules
}

type CreatePromoCodeResponse struct {
	Code *PromoCode
}

type EndRideRequest struct {
	City string
	Ride uuid.UUID
}

type EndRideResponse struct {
	Ride *Ride
}

type GetActiveRidesRequest struct {
	City   string
	Filter RideFilter
}

type GetActiveRidesResponse struct {
	Limit  int
	Offset int
	Rides  []*Ride
}

type GetPromoCodesRequest struct {
	Filter PromoCodeFilter
}

type GetPromoCodesResponse struct {
	Limit  int
	Offset int
	Codes  []*PromoCode
}

type GetUsersRequest struct {
	City   string
	Filter UserFilter
}

type GetUsersResponse struct {
	Limit  int
	Offset int
	Users  []*User
}

type GetVehiclesRequest struct {
	City   string
	Filter VehicleFilter
}

type GetVehiclesResponse struct {
	Limit    int
	Offset   int
	Vehicles []*Vehicle
}

type StartRideRequest struct {
	City    string
	Rider   uuid.UUID
	Vehicle uuid.UUID
}

type StartRideResponse struct {
	Ride *Ride
}

type UpdateRideLocationRequest struct {
	City      string
	Ride      uuid.UUID
	Latitude  float64
	Longitude float64
}

type UpdateRideLocationResponse struct {
	Location *VehicleLocationHistory
}

type Service interface {
	StartRide(context.Context, StartRideRequest) (*StartRideResponse, error)
	EndRide(context.Context, EndRideRequest) (*EndRideResponse, error)
	UpdateRideLocation(context.Context, UpdateRideLocationRequest) (*UpdateRideLocationResponse, error)
	AddUser(context.Context, AddUserRequest) (*AddUserResponse, error)
	AddVehicle(context.Context, AddVehicleRequest) (*AddVehicleResponse, error)
	GetUsers(context.Context, GetUsersRequest) (*GetUsersResponse, error)
	GetVehicles(context.Context, GetVehiclesRequest) (*GetVehiclesResponse, error)
	GetActiveRides(context.Context, GetActiveRidesRequest) (*GetActiveRidesResponse, error)
	GetPromoCodes(context.Context, GetPromoCodesRequest) (*GetPromoCodesResponse, error)
	CreatePromoCode(context.Context, CreatePromoCodeRequest) (*CreatePromoCodeResponse, error)
	ApplyPromoCode(context.Context, ApplyPromoCodeRequest) (*ApplyPromoCodeResponse, error)
}
