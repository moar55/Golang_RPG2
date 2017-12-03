package scan

type _Nothing struct {
	Message string `json:"message"`
	Mode    string `json:"mode"`
}

var Nothing = &_Nothing{Message: "You found nothing here", Mode: "Nothing"}
