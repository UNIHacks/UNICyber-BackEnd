package routes

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/views"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

func SignIn() {
	config.Router.PUT(fmt.Sprintf("/api/%s/signin", config.API_VERSION), views.SignIn_PUT)
}

func SignUp() {
	config.Router.PUT(fmt.Sprintf("/api/%s/signup", config.API_VERSION), views.SignUp_POST)
}
