package search

import "json"

type Country struct {
	LocalizedName	string	`json:"LocalizedName"`
}

type AdministrativeArea struct {
	ID 			  string	`json:"ID"`
	LocalizedName string	`json:"LocalizedName"`
}

type Search struct {
	Version 				int32 				`json:"Version"`
	Key 					string	  			`json:"Key"`
	Type 					string				`json:"Type"`
	Rank 					int32				`json:"Rank"`
	LocalizedName 			string				`json:"LocalizedName"`
	Country 				Country				`json:"Country"`
	AdministrativeArea 		AdministrativeArea 	`json:"AdministrativeArea"`
}