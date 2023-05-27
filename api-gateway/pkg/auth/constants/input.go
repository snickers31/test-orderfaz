package constants

type LoginRequestBody struct {
	Msisdn   string `json:"msisdn"`
	Password string `json:"password"`
}

type RegisterRequestBody struct {
	Msisdn   string `json:"msisdn"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
