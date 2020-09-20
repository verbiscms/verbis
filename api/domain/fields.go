package domain

import "github.com/google/uuid"

type FieldGroup struct {
	UUID 		uuid.UUID 					`json:"uuid"`
	Title 		string 						`json:"title"`
	Fields 		[]Field 					`json:"fields"`
	Locations 	[][]FieldLocation 			`json:"location,omitempty"`
}

type Field struct {
	UUID 			uuid.UUID 				`json:"uuid"`
	Label 			string 					`json:"label"`
	Name 			string 					`json:"name"`
	Type 			string 					`json:"type"`
	Instructions 	string 					`json:"instructions"`
	Required 		bool 					`json:"required"`
	Options 	 	map[string]interface{} 	`json:"options"`
}

type FieldLocation struct {
	Param string
	Operator string
	Value string
}