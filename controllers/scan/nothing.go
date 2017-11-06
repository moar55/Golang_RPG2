package scan

type _Nothing struct {
	Message string `json:"message"`
}

var Nothing = &_Nothing{Message: "You found nothing here"}
