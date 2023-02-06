package controllers

import (
	"strconv"
	"time"

	"github.com/MetaBloxIO/metablox-foundation-services/errval"
)

const nanoMinute = 60000000000 //number of nanoseconds in 1 minute

// If user provided correct nonce, then assign them a new one and return that value
func CreateNonce(ip string) string {
	time.Now().UnixNano()
	NonceLookup[ip] = strconv.Itoa(int(time.Now().UnixNano()))
	return NonceLookup[ip]
}

// Compare the nonce a user has given with the one they are assigned. Current time must also be within 1 minute of the nonce's value
func CheckNonce(ip, givenNonce string) error {
	assignedNonce, found := NonceLookup[ip]
	if !found {
		return errval.ErrNoNonce
	}

	nanoTimestamp, _ := strconv.Atoi(assignedNonce)

	if int64(nanoTimestamp+nanoMinute) < time.Now().UnixNano() {
		delete(NonceLookup, ip) //remove expired nonces
		return errval.ErrExpiredNonce
	}

	if assignedNonce != givenNonce {
		return errval.ErrWrongNonce
	}

	return nil
}

// delete a nonce after it has been successfully used in an operation
func DeleteNonce(ip string) {
	delete(NonceLookup, ip)
}
