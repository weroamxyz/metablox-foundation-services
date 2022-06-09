package errval

import "errors"

var (
	ErrInvalidIssuer          = errors.New("provided did is not a valid issuer")
	ErrAuthFailed             = errors.New("authentication failed")
	ErrNoNonce                = errors.New("no nonce assigned to user")
	ErrExpiredNonce           = errors.New("nonce has expired")
	ErrWrongNonce             = errors.New("provided nonce is incorrect")
	ErrDIDFormat              = errors.New("improperly formatted did")
	ErrDIDNotIssuer           = errors.New("provided did does not match issuer of credential")
	ErrVerifyPresent          = errors.New("failed to verify presentation")
	ErrRenewRevoked           = errors.New("VC has been revoked, cannot renew")
	ErrUnknownIssuer          = errors.New("unknown issuer")
	ErrSecp256k1WrongVMType   = errors.New("must use a verification method with a type of 'EcdsaSecp256k1RecoveryMethod2020' to verify a 'EcdsaSecp256k1Signature2019' proof")
	ErrUnknownProofType       = errors.New("unable to verify unknown proof type")
	ErrUnknownVMType          = errors.New("unable to resolve unknown verification method type")
	ErrJWSAuthentication      = errors.New("square/go-jose: error in cryptographic primitive")
	ErrInvalidSecp256k1PubKey = errors.New("invalid secp256k1 public key")
	ErrMissingAuthentication  = errors.New("failed to find authentication method")
	ErrMissingVM              = errors.New("failed to find verification method")
	ErrWrongAddress           = errors.New("provided public key does not match issuer address")
	ErrETHAddress             = errors.New("provided address is not a correct ETH address")
	ErrVerifySignature        = errors.New("failed to verify signature")
)
