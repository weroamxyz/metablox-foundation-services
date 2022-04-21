package controllers

import (
	"time"

	"github.com/metabloxDID/errval"
)

//If user provided correct nonce, then assign them a new one and return that value
func CreateNonce(ip string) string {
	NonceLookup[ip] = time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST")
	return NonceLookup[ip]
}

//Compare the nonce a user has given with the one they are assigned. Current time must also be within 1 minute of the nonce's value
func CheckNonce(ip, givenNonce string) error {
	assignedNonce, found := NonceLookup[ip]
	if !found {
		return errval.ErrNoNonce
	}

	nonceTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", assignedNonce)
	if err != nil {
		return err
	}
	if nonceTime.Add(time.Minute).Before(time.Now()) {
		delete(NonceLookup, ip) //remove expired nonces
		return errval.ErrExpiredNonce
	}

	if assignedNonce != givenNonce {
		return errval.ErrWrongNonce
	}

	return nil
}

//delete a nonce after it has been successfully used in an operation
func DeleteNonce(ip string) {
	delete(NonceLookup, ip)
}
