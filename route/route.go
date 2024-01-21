package route

import (
	"TP4/controller"
	"net/http"
)

// RegisterRoutes registers all the routes for the application.
func InitServe() {
	//lance ton serveur
	http.HandleFunc("/home", controller.Home)
	http.HandleFunc("/create", controller.Create)
	http.HandleFunc("/person", controller.Person)
	http.HandleFunc("/create_submit", controller.CreatePerson)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/delete", controller.Delete)

	//charge les fichiers statics
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./asset/static"))))

	http.ListenAndServe(":8080", nil)

	/* 	http.HandleFunc("/create", CreateHandler)
	   	http.HandleFunc("/home", HomeHandler)
	   	http.HandleFunc("/person", PersonHandler) */
}
