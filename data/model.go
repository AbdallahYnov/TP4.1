package model

type Person struct {
	ID         int     `json:"id"`
	Nom        string  `json:"nom"`
	Prenom     string  `json:"prenom"`
	Taille     float64 `json:"taille"`
	Sexe       string  `json:"sexe"`
	Poids      int     `json:"poids"`
	Age        int     `json:"age"`
	Competence string  `json:"competence"`
	Aime       string  `json:"aime"`
	Deteste    string  `json:"deteste"`
	Autres     string  `json:"autres"`

}