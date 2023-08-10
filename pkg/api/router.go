package api

import (
	"lib-manager/pkg/controllers"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func Start() {
	fmt.Println("Starting")
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./templates/"))
    r.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", fs))

	r.HandleFunc("/", controllers.Welcome).Methods("GET")
	
	r.HandleFunc("/loginCheck", controllers.ChecklogIn).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")
	r.HandleFunc("/login", controllers.LogIn).Methods("GET")

	r.HandleFunc("/register", controllers.RegisterPage).Methods("GET")
	r.HandleFunc("/register" , controllers.Register).Methods("POST")

	r.HandleFunc("/admin" , controllers.GetAdmin).Methods("GET")
	r.HandleFunc("/client" , controllers.GetClient).Methods("GET")

	r.HandleFunc("/admin/checkin/", controllers.AdminCheckin).Methods("GET")
	r.HandleFunc("/admin/checkin" , controllers.AdminCheckinSubmit).Methods("POST")

	r.HandleFunc("/admin/add", controllers.AdminAdd).Methods("GET")
	r.HandleFunc("/admin/add" , controllers.AdminAddSubmit).Methods("POST")

	r.HandleFunc("/admin/remove", controllers.AdminRemove).Methods("GET")
	r.HandleFunc("/admin/remove" , controllers.AdminRemoveSubmit).Methods("POST")
	
	r.HandleFunc("/admin/checkout", controllers.AdminCheckout).Methods("GET")
	r.HandleFunc("/admin/checkout" , controllers.AdminCheckoutSubmit).Methods("POST")

	r.HandleFunc("/admin/choose", controllers.AdminChoose).Methods("GET")
	r.HandleFunc("/admin/choose/accept" , controllers.AdminAccept).Methods("POST")
	r.HandleFunc("/admin/choose/deny" , controllers.AdminDeny).Methods("POST")
	
	r.HandleFunc("/checkout", controllers.Checkout).Methods("GET")
	r.HandleFunc("/checkout" , controllers.CheckoutSubmit).Methods("POST")

	r.HandleFunc("/checkin", controllers.Checkin).Methods("GET")
	r.HandleFunc("/checkin" , controllers.CheckinSubmit).Methods("POST")

	http.ListenAndServe(":8000", r)
}
