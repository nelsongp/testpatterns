package credentialValidator

type CredentialValidator interface {
	Validate(ID string, password string) error
}
