package duffel

import (
	"net/http"
	"time"
)

const userAgentString = "duffel-go/1.0"
const defaultHost = "https://api.duffel.com/"

type (
	Duffel interface {
		OfferRequestClient
	}

	OfferRequestInput struct {
		Passengers   []Passenger         `json:"passengers"`
		Slices       []OfferRequestSlice `json:"slices"`
		CabinClass   CabinClass          `json:"cabin_class"`
		ReturnOffers bool                `json:"-"`
	}

	OfferRequestSlice struct {
		DepartureDate Date   `json:"departure_date"`
		Destination   string `json:"destination"`
		Origin        string `json:"origin"`
	}

	OfferResponse struct {
		Offers     []Offer     `json:"offers"`
		Slices     []Slice     `json:"slices"`
		Passengers []Passenger `json:"passengers"`
	}

	Offer struct {
		Slices           []Slice     `json:"slices"`
		Passengers       []Passenger `json:"passengers"`
		UpdatedAt        time.Time   `json:"updated_at"`
		TotalEmissionsKg float64     `json:"total_emissions_kg"`
		TotalCurrency    string      `json:"total_currency"`
		TotalAmount      float64     `json:"total_amount"`
		TaxCurrency      string      `json:"tax_currency"`
		TaxAmount        float64     `json:"tax_amount"`
	}

	Slice struct {
		OriginType      string    `json:"origin_type"`
		Origin          Location  `json:"origin"`
		DestinationType string    `json:"destination_type"`
		Destination     Location  `json:"destination"`
		DepartureDate   Date      `json:"departure_date,omitempty"`
		CreatedAt       time.Time `json:"created_at,omitempty"`
		Segments        []Flight  `json:"segments,omitempty"`
	}

	Flight struct {
		Passengers                   []SegmentPassenger `json:"passengers"`
		Origin                       Location           `json:"origin"`
		OriginTerminal               string             `json:"origin_terminal"`
		OperatingCarrierFlightNumber string             `json:"operating_carrier_flight_number"`
		OperatingCarrier             Airline            `json:"operating_carrier"`
		MarketingCarrierFlightNumber string             `json:"marketing_carrier_flight_number"`
		MarketingCarrier             Airline            `json:"marketing_carrier"`
		Duration                     time.Duration      `json:"duration"`
		Distance                     float64            `json:"distance, omitempty"`
		DestinationTerminal          string             `json:"destination_terminal"`
		Destination                  Location           `json:"destination"`
		DepartingAt                  time.Time          `json:"departing_at"`
		ArrivingAt                   time.Time          `json:"arriving_at"`
		Aircraft                     Equipment          `json:"aircraft"`
	}

	SegmentPassenger struct {
		ID                      string     `json:"id"`
		FareBasisCode           string     `json:"fare_basis_code"`
		CabinClassMarketingName string     `json:"cabin_class_marketing_name"`
		CabinClass              CabinClass `json:"cabin_class"`
		Baggages                []Baggage  `json:"baggages"`
	}

	Equipment struct {
		IATACode string `json:"iata_code"`
		ID       string `json:"id"`
		Name     string `json:"name"`
	}

	Airline struct {
		Name     string `json:"name"`
		IATACode string `json:"iata_code"`
		ID       string `json:"id"`
	}

	Baggage struct {
		Quantity int    `json:"quantity"`
		Type     string `json:"type"`
	}

	Location struct {
		ID              string     `json:"id"`
		Type            string     `json:"type"`
		TimeZone        string     `json:"time_zone"`
		Name            string     `json:"name"`
		Longitude       *float64   `json:"longitude,omitempty"`
		Latitude        *float64   `json:"latitude,omitempty"`
		ICAOCode        *string    `json:"icao_code,omitempty"`
		IATACountryCode *string    `json:"iata_country_code,omitempty"`
		IATACode        *string    `json:"iata_code,omitempty"`
		IATACityCode    *string    `json:"iata_city_code,omitempty"`
		CityName        *string    `json:"city_name,omitempty"`
		City            *Location  `json:"city,omitempty"`
		Airports        []Location `json:"airports,omitempty"`
	}

	Passenger struct {
		ID                       string                    `json:"id,omitempty"`
		FamilyName               string                    `json:"family_name"`
		GivenName                string                    `json:"given_name"`
		Age                      *int                      `json:"age,omitempty"`
		Type                     *PassengerType            `json:"type,omitempty"`
		LoyaltyProgrammeAccounts []LoyaltyProgrammeAccount `json:"loyalty_programme_accounts,omitempty"`
	}

	LoyaltyProgrammeAccount struct {
		IATACode string `json:"iata_code"`
	}

	PassengerType string

	CabinClass string

	Offers []Offer

	Option  func(*Options)
	Options struct {
		Version   string
		Host      string
		UserAgent string
		HttpDoer  *http.Client
	}

	client struct {
		httpDoer *http.Client
		APIToken string
		options  *Options
	}
)

const (
	PassengerTypeAdult             PassengerType = "adult"
	PassengerTypeChild             PassengerType = "child"
	PassengerTypeInfantWithoutSeat PassengerType = "infant_without_seat"

	CabinClassEconomy  CabinClass = "economy"
	CabinClassPremium  CabinClass = "premium"
	CabinClassBusiness CabinClass = "business"
	CabinClassFirst    CabinClass = "first"
)

func New(apiToken string, opts ...Option) Duffel {
	options := &Options{
		Version:   "beta",
		UserAgent: userAgentString,
		Host:      defaultHost,
		HttpDoer:  http.DefaultClient,
	}
	for _, opt := range opts {
		opt(options)
	}
	c := &client{
		httpDoer: options.HttpDoer,
		APIToken: apiToken,
		options:  options,
	}
	return c
}

// Assert that our interface matches
var (
	_ Duffel = (*client)(nil)
)

func (p PassengerType) String() string {
	return string(p)
}
