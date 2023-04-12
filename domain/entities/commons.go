package entities

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JwtToken struct {
	Token string `json:"token"`
}
