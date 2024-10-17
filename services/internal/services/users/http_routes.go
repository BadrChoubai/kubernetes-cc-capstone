package users

import "net/http"

func (a *UserService) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("/register", a.RegisterUserHandler)
	router.HandleFunc("/login", a.LoginUserHandler)
}

func (a *UserService) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (a *UserService) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
