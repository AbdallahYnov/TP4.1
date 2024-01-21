package main

import (
	"TP4/route"
	"TP4/template"
)

func main() {

	template.Inittemplate()
	route.InitServe()
	/* //import setroutes
	web.RegisterRoutes()




	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		person, err := LoadPerson()
		fmt.Println(person)
		if err != nil {
			fmt.Println("erreur load person")
			return
		}
		personChoices := getRandomPerson(person, 9)
		temp.ExecuteTemplate(w, "home", personChoices)
	})

	filePath := "person.json"
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var people []Person
	err = json.Unmarshal(jsonData, &people)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("JSON:", string(jsonData))

	for _, person := range people {
		fmt.Println(person)
	} */
	/*for _, person := range people {
		fmt.Println("Nom:", person.Nom)
		fmt.Println("Prenom:", person.Prenom)
		fmt.Println("Taille:", person.Taille)
		fmt.Println("Sexe:", person.Sexe)
		fmt.Println("Poids:", person.Poids)
		fmt.Println("Age:", person.Age)
		fmt.Println("ID:", person.ID)
		fmt.Println("Competence:", person.Competence)
		fmt.Println("Aime:", person.Aime)
		fmt.Println("Deteste:", person.Deteste)
		fmt.Println("Autres:", person.Autres)
	}*/
}
