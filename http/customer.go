package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ibraheamkh/clinicy"
	"github.com/ibraheamkh/clinicy/security"
)

func (h *Handler) customerAuthHandler(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("content-type", "application/json")
	requestTemplate := struct {
		Mobile string `json:"mobile"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&requestTemplate); err != nil {
		SendErrorResponse(w, err.Error(), 400)
		return
	}
	// TODO: validate user input, it should be a mobile number.

	// ###############
	customer, _ := h.CustomerService.GetCustomerByMobile(requestTemplate.Mobile)
	if customer == nil {
		// check is user exist or create one for him
		newCustomer := &clinicy.Customer{}
		newCustomer.Account.Mobile = requestTemplate.Mobile
		err := h.CustomerService.CreateCustomerAccount(newCustomer)
		if err != nil {
			SendErrorResponse(w, err.Error(), 400)
			return
		}
		customer = newCustomer
	}

	// ###############
	// verify the user logic
	otp, err := security.GenerateOTP()

	sessionID, err := security.GenerateRandomString(32)
	if err != nil {
		SendErrorResponse(w, err.Error(), 400)
		return
	}
	log.Println(customer.ID)
	newSession := &clinicy.Session{
		AccountID: customer.ID,
		SessionID: sessionID,
		OTP:       otp,
		OTPStatus: "new",
		Status:    "deactivated",
	}

	err = h.SessionService.CreateSession(newSession)
	if err != nil {
		SendErrorResponse(w, err.Error(), 400)
		return
	}

	w.Header().Set("authorization", "Bearer "+sessionID)

	// TODO send sms message to the client
	SendResponseWithMessage(w, nil, "sms verification code is "+otp, 200)
}

func (h *Handler) customerAuthVerifySMSHandler(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("authorization")
	s := strings.Split(authorizationHeader, " ")
	var sessionID string
	if len(s) == 2 {
		sessionID = s[1]
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	session, err := h.SessionService.GetSession(sessionID)
	if err != nil || session == nil { // if no session or there is an error
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	requestTemplate := struct {
		OTP string `json:"OTP"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&requestTemplate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	if session.OTP == requestTemplate.OTP && session.OTPStatus == "new" {
		session.OTPStatus = "verified"
		session.Status = "activated"
		h.SessionService.UpdateSession(session)
		SendResponseWithMessage(w, nil, "you have successfully verified your mobile number", 200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
