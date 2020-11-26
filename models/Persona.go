package models

import (
	"gorm.io/gorm"
)

type Persona struct {
	gorm.Model
	Nombre          string `json:"nombre"`
	Paterno         string `json:"paterno"`
	Materno         string `json:"materno"`
	Edad            string `json:"edad"`
	Fechanacimiento string `json:"fechanacimiento"`
	Esatadocivil    string `json:"estadocivil"`
}
