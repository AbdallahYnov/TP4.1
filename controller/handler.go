package controller

import (
	model "TP4/data"
	"TP4/template"
	"TP4/utility"
	"fmt"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	peoplez, _ := utility.ReadPeople("data/perso.json")
	template.Temp.ExecuteTemplate(w, "home", peoplez)
}

func Create(w http.ResponseWriter, r *http.Request) {
	template.Temp.ExecuteTemplate(w, "create", nil)
}

func Person(w http.ResponseWriter, r *http.Request) {
	template.Temp.ExecuteTemplate(w, "person", nil)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	Poids, _ := strconv.Atoi(r.FormValue("poids"))
	Age, _ := strconv.Atoi(r.FormValue("age"))
	Taille, _ := strconv.ParseFloat(r.FormValue("taille"), 64)

	peoplez, _ := utility.ReadPeople("data/perso.json")
	// Get the next available ID
	nextID := utility.GetNextAvailableID(getIDs(peoplez))
	// Read JSON from request body
	var person = model.Person{
		ID:         nextID,
		Nom:        r.FormValue("nom"),
		Prenom:     r.FormValue("prenom"),
		Taille:     Taille,
		Sexe:       r.FormValue("sexe"),
		Poids:      Poids,
		Age:        Age,
		Aime:       r.FormValue("aime"),
		Competence: r.FormValue("competence"),
		Deteste:    r.FormValue("deteste"),
		Autres:     r.FormValue("autres"),
	}
	fmt.Println(person.ID)
	// Add the person to the utility or database
	utility.AddPerson(person.ID, person.Nom, person.Prenom, person.Taille, person.Sexe, person.Poids, person.Age, person.Competence, person.Aime, person.Deteste, person.Autres)

	// Respond with success
	template.Temp.ExecuteTemplate(w, "success", nil)
}

/*
	func Edit(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		_ = id
		//récupérérer les valeurs des input r.FormvValue
		//open json
		//parcourir le json et break json => .ID= id
		//mise à jour du
		template.Temp.ExecuteTemplate(w, "edit", nil)
	}
*/
func Edit(w http.ResponseWriter, r *http.Request) {
	// Print the entire URL for debugging
	fmt.Println("URL:", r.URL.String())
	// Get the "id" parameter from the URL query
	idStr := r.URL.Query().Get("id")

	// Convert the ID string to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Debugging statement
		fmt.Println("Invalid ID:", idStr)

		// Handle the error (e.g., log it, return an error response, etc.)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Check if any form values are present and update the corresponding fields
	if r.Method == http.MethodPost {
		r.ParseForm()

		Poids, _ := strconv.Atoi(r.FormValue("poids"))
		Age, _ := strconv.Atoi(r.FormValue("age"))
		Taille, _ := strconv.ParseFloat(r.FormValue("taille"), 64)
		// Update individual fields based on form values
		err := utility.ModifyPerson(
			id,
			r.FormValue("nom"),
			r.FormValue("prenom"),
			Taille,
			r.FormValue("sexe"),
			Poids,
			Age,
			r.FormValue("competence"),
			r.FormValue("aime"),
			r.FormValue("deteste"),
			r.FormValue("autres"),
		)
		if err != nil {
			// Handle the error (e.g., log it, return an error response, etc.)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Redirect to the person's details page or render a success template
		//http.Redirect(w, r, "/person?id="+idStr, http.StatusSeeOther)
		template.Temp.ExecuteTemplate(w, "success", nil)
		return
	}

	// Retrieve the person's data for rendering the "edit" template
	personToEdit, err := utility.GetPersonByID(id)
	if err != nil {
		// Handle the error (e.g., log it, return an error response, etc.)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Render the "edit" template with the person data for editing
	template.Temp.ExecuteTemplate(w, "edit", personToEdit)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Perform deletion (you should implement this function in your utility package)
	err = utility.DeletePerson(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Redirect to the home page after deletion
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// Helper function to get existing IDs
func getIDs(people []model.Person) []int {
	var ids []int
	for _, person := range people {
		ids = append(ids, person.ID)
	}
	return ids
}
