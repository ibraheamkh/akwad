package mongo

import (
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
)

func Dial(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)

	session.SetMode(mgo.Monotonic, true)
	return session, err
}

func PopulateDB(s *mgo.Session) error {
	//populateDB(s)
	return nil
}

//TODO create more users
/*
func populateDB(session *mgo.Session) {
	log.Println("Populating DB")
	//Crete the services,
	//init the service
	//service := &Service{DB: db}

	usersRepo := &UserRepository{Session: session}

	//Create Some users

	//setting up dummy time
	availableFromTime := time.Time{}
	duration, _ := time.ParseDuration("8h30m")
	availableFromTime = availableFromTime.Add(duration)

	admin := &carwash.User{
		Username:    "admin",
		Password:    "admin",
		Role:        "admin",
		FullName:    "G of the House Ghonaim",
		Email:       "admin@job.com",
		PhoneNumber: "0562262286",

		City: "Riyadh",

		Nationality: "Saudi",
	}

	othman := &carwash.User{
		Username:    "othman",
		Password:    "123",
		Role:        "user",
		FullName:    "Otham Ascehibany",
		Email:       "othman@akwad.com",
		PhoneNumber: "0555555555",

		City: "Riyadh",

		Nationality: "Saudi",
	}

	ibraheam := &carwash.User{
		Username:    "ibraheam",
		Password:    "123",
		Role:        "user",
		FullName:    "Ibraheam Alkhalifah",
		Email:       "iheemoo@gmail.com",
		PhoneNumber: "0555555555",
		City:        "Riyadh",
		Nationality: "Saudi",
	}

	//Insert users
	err := usersRepo.CreateUser(ibraheam)

	if err != nil {
		log.Println(err)
	}

	//Insert users
	err = usersRepo.CreateUser(admin)

	if err != nil {
		log.Println(err)
	}

	err = usersRepo.CreateUser(othman)

	if err != nil {
		log.Println(err)
	}

	fako.Register("nationality", generateNationality)
	fako.Register("available_from", generateAvailableFrom)
	fako.Register("available_to", generateAvailableTo)
	fako.Register("qualifications", generateQualifications)

	//generate users
	for i := 0; i < 50; i++ {
		uu := &carwash.User{}
		uu.Role = "user"
		fako.Fill(uu)
		err = usersRepo.CreateUser(uu)
		if err != nil {
			log.Println("Error inserting generated user to db")
			log.Println(err)
		}
	}

}
*/
//TODO improve dummy data generators

//Helper methods to hash passwords
func hashPassword(pass string) string {
	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error Creating hash: %v", err)
		return ""
	}

	return string(hash)
}

func compareHash(pass, hashFromDatabase string) bool {
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(hashFromDatabase), []byte(pass))
	if err != nil {
		// TODO: Properly handle error
		log.Println(err)
		return false
	}

	return true
}

func generatePassword() string {
	return "123"
}

func generateNationality() string {
	return "Saudi"
}

func generateQualifications() string {
	return "Bachelor"
}

func generateAvailableFrom() string {
	return "8:00 AM"
}

func generateAvailableTo() string {
	return "8:00 PM"
}

func generateIndustry() string {
	return "Technology"
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
