package models

type Mots struct {
	IdMot        int    `json:idMot`
	NomMot       string `json:nomMot`
	DefMot       string `json:defMot`
	NbLettresMot int    `json:nbLettresMot`
}
