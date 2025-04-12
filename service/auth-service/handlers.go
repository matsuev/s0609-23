package main

import (
	"github.com/gofiber/fiber/v2"
)

// ConnectRequestData ...
type ConnectRequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ConnectRequest ...
type ConnectRequest struct {
	Client    string             `json:"client"`
	Transport string             `json:"transport"`
	Protocol  string             `json:"protocol"`
	Encoding  string             `json:"encoding"`
	Name      string             `json:"name,omitempty"`
	Version   string             `json:"version,omitempty"`
	Data      ConnectRequestData `json:"data"`
	B64Data   string             `json:"b64data,omitempty"`
	Channels  []string           `json:"channels,omitempty"`
}

// Error ...
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

// ConnectResult ...
type ConnectResult struct {
	User string `json:"user"`
}

// ConnectResponse ...
type ConnectResponse struct {
	Error  *Error         `json:"error,omitempty"`
	Result *ConnectResult `json:"result,omitempty"`
}

// connectHandler ...
func connectHandler(ctx *fiber.Ctx) error {
	request := &ConnectRequest{}

	if err := ctx.BodyParser(request); err != nil {
		return ctx.JSON(ConnectResponse{
			Error: &Error{
				Code:    107,
				Message: "bad request",
			},
		})
	}

	if request.Data.Username != "alex" || request.Data.Password != "password" {
		return ctx.JSON(ConnectResponse{
			Error: &Error{
				Code:    101,
				Message: "unauthorized",
			},
		})
	}

	return ctx.JSON(ConnectResponse{
		Result: &ConnectResult{
			User: "alex",
		},
	})
}
