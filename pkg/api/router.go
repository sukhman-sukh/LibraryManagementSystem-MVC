package api

import (
	"lib-manager/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./templates/"))
	r.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", fs))
	r.NotFoundHandler = http.HandlerFunc(controllers.PageNotFound)

	r.HandleFunc("/", controllers.Welcome).Methods("GET")

	r.HandleFunc("/loginCheck", controllers.ChecklogIn).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")
	r.HandleFunc("/login", controllers.LogIn).Methods("GET")

	r.HandleFunc("/register", controllers.RegisterPage).Methods("GET")
	r.HandleFunc("/register", controllers.Register).Methods("POST")

	r.HandleFunc("/admin", controllers.GetAdmin).Methods("GET")
	r.HandleFunc("/client", controllers.GetClient).Methods("GET")

	r.HandleFunc("/admin/checkin", controllers.AdminCheckinSubmit).Methods("POST")

	r.HandleFunc("/admin/add", controllers.AdminAdd).Methods("GET")
	r.HandleFunc("/admin/add", controllers.AdminAddSubmit).Methods("POST")

	r.HandleFunc("/admin/remove", controllers.AdminRemove).Methods("GET")
	r.HandleFunc("/admin/remove", controllers.AdminRemoveSubmit).Methods("POST")

	r.HandleFunc("/admin/checkout", controllers.AdminCheckoutSubmit).Methods("POST")

	r.HandleFunc("/admin/choose/accept", controllers.AdminAccept).Methods("POST")
	r.HandleFunc("/admin/choose/deny", controllers.AdminDeny).Methods("POST")

	r.HandleFunc("/checkout", controllers.CheckoutSubmit).Methods("POST")

	r.HandleFunc("/checkin", controllers.CheckinSubmit).Methods("POST")

	r.HandleFunc("/error403", controllers.ForbiddenAccess).Methods("GET")
	r.HandleFunc("/error500", controllers.InternalError).Methods("GET")

	http.ListenAndServe(":8000", r)
}
