package errors

type Err struct {
	Message error `json:"message"`
}

type Error struct {
	HTTPStatus int
	Message    ErrorMessage
	Mode       string `json:"mode"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

var WrongCredentials = Error{HTTPStatus: 401, Message: ErrorMessage{Message: "Wrong username or password"}, Mode: "Error"}
var NoBot = Error{HTTPStatus: 401, Message: ErrorMessage{Message: "Please create a bot!"}, Mode: "Error"}
var NotLoggedIn = Error{HTTPStatus: 401, Message: ErrorMessage{Message: "Please log in first"}, Mode: "Error"}
var HaveBot = Error{HTTPStatus: 401, Message: ErrorMessage{Message: "You already have a bot"}, Mode: "Error"}
var InvalidParameters = Error{HTTPStatus: 422, Message: ErrorMessage{Message: "Invalid Parameters"}, Mode: "Error"}
var SearchForShop = Error{HTTPStatus: 400, Message: ErrorMessage{Message: "You need to search for near shops first"}, Mode: "Error"}
var ItemNotFound = Error{HTTPStatus: 404, Message: ErrorMessage{Message: "La walahi, 5elset men shwaya"}, Mode: "Error"}
var NoEnoughFakka = Error{HTTPStatus: 403, Message: ErrorMessage{Message: "La shofly fakka ya basha"}, Mode: "Error"}
