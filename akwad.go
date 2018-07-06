package akwad

//@author ibrahim
// Clinicy models

import "time"

type Session struct {
	SessionID string `bson:"session_id"`
	AccountID string `bson:"user_id"`
	DeviceID  string `bson:"device_id"`
	Status    string `bson:"status"` // active, inactive
	OTP       string `bson:"-"`      // should have validiry period
	OTPStatus string `bson:"-"`      // new, verified, expired
}

//TODO: these are domain model users and will be modified to fit our needs

//User is main system user
type Account struct {
	FullName    string `json:"full_name,omitempty" bson:"full_name,omitempty" db:"full_name" valid:"-" fako:"full_name"`
	Username    string `json:"username,omitempty" bson:"username,omitempty" db:"username" valid:"-" fako:"user_name"`
	Password    string `json:"password,omitempty" bson:"password,omitempty" db:"password" valid:"-" fako:"my_password"`
	Email       string `json:"email,omitempty" bson:"eamil,omitempty" db:"email" valid:"email" fako:"email_address"`
	Role        string `json:"role,omitempty" bson:"role,omitempty" db:"role" valid:"-"`
	Mobile      string `json:"mobile,omitempty" bson:"mobile,omitempty" db:"mobile" valid:"" fako:"mobile"`
	City        string `json:"city,omitempty" bson:"city,omitempty" db:"city" valid:"-" fako:"city"`
	Nationality string `json:"nationality,omitempty" bson:"nationality,omitempty" db:"nationality" valid:"-" fako:"nationality"`
}

//Clinic represent a clinic in the system
type Clinic struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	MainSpeciality   string    `json:"mainSpeciality"`
	SecondSpeciality string    `json:"secondSpeciality"`
	Rating           float64   `json:"rating"`
	Branches         []string  `json:"branches"`
	Doctors          []*Doctor `json:"doctors"`
}

//Doctor represent a doctor in the system
type Doctor struct {
	ID               int            `json:"id"`
	Title            string         `json:"title"`
	Name             string         `json:"name"`
	MainSpeciality   string         `json:"mainSpeciality"`
	SecondSpeciality string         `json:"secondSpeciality"`
	Rating           float64        `json:"rating"`
	Availablilty     []Availability `json:"availablilty"` //available times
}

// Availability represents available time
type Availability struct {
	Datetime string `json:"datetime"`
	Branch   string `json:"branch"`
	Status   string `json:"status"`
}

//Appointment represents an appointment in the system
type Appointment struct {
	ID                     int     `json:"id"`
	Name                   string  `json:"name"`
	MainSpeciality         string  `json:"mainSpeciality"`
	SecondSpeciality       string  `json:"secondSpeciality"`
	Rating                 float64 `json:"rating"`
	ClinicName             string  `json:"clinicName"`
	ClinicSpeciality       string  `json:"clinicSpeciality,omitempty"`
	ClinicSecondSpeciality string  `json:"clinicSecondSpeciality"`
	Status                 string  `json:"status"`
	Price                  string  `json:"price"`
	Datetime               string  `json:"datetime"`
	Branch                 string  `json:"branch"`
	PatientName            string  `json:"patientName,omitempty"`
	Availablilty           []struct {
		Datetime string `json:"datetime"`
		Branch   string `json:"branch"`
		Status   string `json:"status"`
	} `json:"availablilty"`
	ClincSpeciality string `json:"clincSpeciality,omitempty"`
}

//Booking represent a booking
type Booking struct {
}

//Here are car wash domain models
type Customer struct {
	ID      string  `json:"id" bson:"_id"`
	Account Account `json:"account" bson:"account"`
}

type Captain struct {
	ID             string  `json:"id" bson:"_id"`
	Account        Account `json:"account" bson:"account"`
	AssignedOrders []Order `json:"assigned_orders" bson:"assigned_orders"`
}

type Order struct {
	ID           string    `json:"id" bson:"_id"`
	AccountID    string    `json:"account_id" bson:"account_id"`
	Source       string    `json:"source" bson:"source"`
	Status       string    `json:"status" bson:"status"`
	Cars         []Car     `json:"cars" bson:"cars"`
	ReservedDate time.Time `json:"reserved_date" bson:"reserved_date"`
	Location     Location  `json:"location" bson:"location"`
	TotalCost    int       `json:"cost" bson:"cost"`
}

type Car struct {
	Type  string `json:"type" bson:"type"`
	Count int    `json:"count" bson:"count"`
	Cost  int    `json:"cost" bson:"cost"`
}

type Location struct {
	Latitude   string `json:"latitude" bson:"latitude"`
	Longitude  string `json:"longitude" bson:"longitude"`
	Directions string `json:"directions" bson:"directions"`
}
