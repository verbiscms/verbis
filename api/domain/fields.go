package domain

import "github.com/google/uuid"

// FieldGroup defines a group of JSON fields
type FieldGroup struct {
	UUID 		uuid.UUID 					`json:"uuid"`
	Title 		string 						`json:"title"`
	Fields 		*[]Field 					`json:"fields,omitempty"`
	Locations 	[][]FieldLocation 			`json:"location,omitempty"`
}

// Field defines an individual field type
type Field struct {
	UUID 			uuid.UUID 				`json:"uuid"`
	Label 			string 					`json:"label"`
	Name 			string 					`json:"name"`
	Type 			string 					`json:"type"`
	Instructions 	string 					`json:"instructions"`
	Required 		bool 					`json:"required"`
	Logic 			*[][]ConditionalLogic 	`json:"conditional_logic"`
	Wrapper 		*Wrapper 				`json:"wrapper"`
	Options 	 	map[string]interface{} 	`json:"options"`
	SubFields 		*[]Field 				`json:"sub_fields,omitempty"`
}

type FieldFilter struct {
	Resource 			string 				`json:"resource"`
	PageTemplate 		string 				`json:"template"`
	Layout 				string 				`json:"layout"`
	Category 			string 				`json:"category"`
}

// FieldLocation defines where the FieldGroup will appear
type FieldLocation struct {
	Param string
	Operator string
	Value string
}

type Wrapper struct {
	Width int `json:"width"`
}

type ConditionalLogic struct {
	Param string `json:"param"`
	Operator string `json:"operator"`
	Value string `json:"value"`
}