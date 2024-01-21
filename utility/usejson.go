package utility

import (
	model "TP4/data"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// search
func ReadPerson(filename string, id int) model.Person {
	var people []model.Person

	content, err := os.ReadFile(filename)
	if err != nil {
		return model.Person{}
	}

	err = json.Unmarshal(content, &people)
	if err != nil {
		return model.Person{}
	}

	for _, person := range people {
		if person.ID == id {
			return person
		}
	}
	return model.Person{}
}

// menu de perso
func ReadPeople(filename string) ([]model.Person, error) {
	var people []model.Person

	content, err := os.ReadFile(filename)
	if err != nil {
		return []model.Person{}, err
	}

	err = json.Unmarshal(content, &people)
	if err != nil {
		fmt.Println(err)
		return []model.Person{}, err
	}
	//fmt.Println(people)
	return people, nil
}

func AddPerson(nouvelID int, nom string, prenom string, taille float64, sexe string, poids int, age int, competence string, aime string, deteste string, autres string) error {
	filePath := "data/perso.json"

	// Read file
	contenu, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var personnages []model.Person
	err = json.Unmarshal(contenu, &personnages)
	if err != nil {
		return err
	}

	nouveauPersonnage := model.Person{
		ID:         nouvelID,
		Nom:        nom,
		Prenom:     prenom,
		Taille:     taille,
		Sexe:       sexe,
		Poids:      poids,
		Age:        age,
		Competence: competence,
		Aime:       aime,
		Deteste:    deteste,
		Autres:     autres,
	}

	// Add the new character to the list
	personnages = append(personnages, nouveauPersonnage)

	// Serialize the updated list to JSON
	newContent, err := json.MarshalIndent(personnages, "", "    ")
	if err != nil {
		return err
	}

	// Write the updated content to the file perso.json
	err = os.WriteFile(filePath, newContent, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Personnage ajouté avec succès!")

	return nil
}

func ModifyPerson(id int, nom string, prenom string, taille float64, sexe string, poids int, age int, competence string, aime string, deteste string, autres string) error {
	filePath := "data/perso.json"

	// Read file
	contenu, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var personnages []model.Person
	err = json.Unmarshal(contenu, &personnages)
	if err != nil {
		return err
	}

	// Find the person with the specified ID
	found := false
	for i, personnage := range personnages {
		if personnage.ID == id {
			personnages[i].Nom = nom
			personnages[i].Prenom = prenom
			personnages[i].Taille = taille
			personnages[i].Sexe = sexe
			personnages[i].Poids = poids
			personnages[i].Age = age
			personnages[i].Competence = competence
			personnages[i].Aime = aime
			personnages[i].Deteste = deteste
			personnages[i].Autres = autres
			found = true
			break
		}
	}

	if !found {
		return errors.New("person not found")
	}

	// Serialize the updated list to JSON
	newContent, err := json.MarshalIndent(personnages, "", "    ")
	if err != nil {
		return err
	}

	// Write the updated content to the file perso.json
	err = os.WriteFile(filePath, newContent, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Personnage modifié avec succès!")

	return nil
}

func GetPersonByID(id int) (model.Person, error) {
	filePath := "data/perso.json"

	// Read file
	contenu, err := os.ReadFile(filePath)
	if err != nil {
		return model.Person{}, err
	}

	var personnages []model.Person
	err = json.Unmarshal(contenu, &personnages)
	if err != nil {
		return model.Person{}, err
	}

	// Find the person with the specified ID
	for _, personnage := range personnages {
		if personnage.ID == id {
			return personnage, nil // Person found, return the person's information
		}
	}

	// Person not found, return an error
	return model.Person{}, errors.New("person not found")
}

var idCounter = 0

// GenerateNewID generates a new unique identifier
func GenerateNewID() int {
	idCounter++
	return idCounter
}

// IsEmpty checks if the data file is empty
func GetLastID() byte {
	data, _ := ioutil.ReadFile("data/perso.json")

	return data[len(data)-1]
}

// DeletePerson deletes a person with the given ID.
func DeletePerson(id int) error {
	filePath := "data/perso.json"

	// Read file
	contenu, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var personnages []model.Person
	err = json.Unmarshal(contenu, &personnages)
	if err != nil {
		return err
	}

	// Find the index of the person with the given ID
	index := -1
	for i, person := range personnages {
		if person.ID == id {
			index = i
			break
		}
	}

	// If the person is not found, return an error
	if index == -1 {
		return errors.New("person not found")
	}

	// Remove the person from the slice
	personnages = append(personnages[:index], personnages[index+1:]...)

	// Serialize the updated list to JSON
	newContent, err := json.MarshalIndent(personnages, "", "    ")
	if err != nil {
		return err
	}

	// Write the updated content to the file perso.json
	err = os.WriteFile(filePath, newContent, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Personnage supprimé avec succès!")

	return nil
}

func GetNextAvailableID(existingIDs []int) int {
	// Sort the existing IDs in ascending order
	sort.Ints(existingIDs)

	// Iterate through the sorted IDs and find the first gap
	for i, id := range existingIDs {
		if i+1 < id {
			// Gap found, return the first available ID
			return i + 1
		}
	}

	// If no gap is found, return the next consecutive ID
	return len(existingIDs) + 1
}
