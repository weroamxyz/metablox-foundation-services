package errval

import "errors"

var ErrInvalidIssuer = errors.New("provided did is not a valid issuer")
var ErrAuthFailed = errors.New("authentication failed")
var ErrNoNonce = errors.New("no nonce assigned to user")
var ErrExpiredNonce = errors.New("nonce has expired")
var ErrWrongNonce = errors.New("provided nonce is incorrect")
var ErrDIDFormat = errors.New("improperly formatted did")
var ErrDIDNotIssuer = errors.New("provided did does not match issuer of credential")
var ErrVerifyPresent = errors.New("failed to verify presentation")
var ErrRenewRevoked = errors.New("VC has been revoked, cannot renew")
var ErrUnknownIssuer = errors.New("unknown issuer")
var ErrSecp256k1WrongVMType = errors.New("must use a verification method with a type of 'EcdsaSecp256k1RecoveryMethod2020' to verify a 'EcdsaSecp256k1Signature2019' proof")
var ErrUnknownProofType = errors.New("unable to verify unknown proof type")
var ErrUnknownVMType = errors.New("unable to resolve unknown verification method type")
var ErrJWSAuthentication = errors.New("square/go-jose: error in cryptographic primitive")
var ErrInvalidSecp256k1PubKey = errors.New("invalid secp256k1 public key")
var ErrMissingAuthentication = errors.New("failed to find authentication method")
var ErrMissingVM = errors.New("failed to find verification method")
var ErrWrongAddress = errors.New("provided public key does not match issuer address")
