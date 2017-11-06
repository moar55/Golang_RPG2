package errors

type Err struct {
	Message error `json:"message"`
}

type Error struct {
	HTTPStatus int
	Message    ErrorMessage
}

type ErrorMessage struct {
	Message string `json:"error"`
}

var WrongCredentials = Error{HTTPStatus: 401, Message: ErrorMessage{Message: "Wrong username or password"}}
var NoBot = Error{HTTPStatus: 401, Message: ErrorMessage{Message: "Please create a bot!"}}
var NotLoggedIn = Error{HTTPStatus: 401, Message: ErrorMessage{Message: "Please log in first"}}
var HaveBot = Error{HTTPStatus: 401, Message: ErrorMessage{Message: "You already have a bot"}}
var InvalidParameters = Error{HTTPStatus: 422, Message: ErrorMessage{Message: "Invalid Parameters"}}
var SearchForShop = Error{HTTPStatus: 400, Message: ErrorMessage{Message: "You need to search for near shops first"}}
var ItemNotFound = Error{HTTPStatus: 404, Message: ErrorMessage{Message: "La walahi, 5elset men shwaya"}}
var NoEnoughFakka = Error{HTTPStatus: 403, Message: ErrorMessage{Message: "La shofly fakka ya basha"}}
