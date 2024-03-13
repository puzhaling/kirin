package backends

import "fmt"

type Headline struct {
	EffectiveDate 		string		`json:"EffectiveDate"`
	EffectiveEpochDate 	int64    	`json:"EffectiveEpochDate"`
	Severity 			int32       `json:"Severity"`
	Text 				string	 	`json:"Text"`
	Category 			string		`json:"Category"`
	EndDate 			string		`json:"EndDate"`
	EndEpochDate 		int64       `json:"EndEpochDate"`
	MobileLink 			string      `json:"MobileLink"`
	Link 				string	    `json:"Link"`
}

type Maximum struct {
	Value		float32		`json:"Value"`
	Unit 		string      `json:"Unit"`
	UnitType 	int32      	`json:"UnitType"`
}

type Minimum struct {
	Value 		float32		`json:"Value"`
	Unit 		string      `json:"Unit"`
	UnitType 	int32      	`json:"UnitType"`
}

type Temperature struct {
	Minimum		Minimum 	`json:"Minimum"`
	Maximum 	Maximum 	`json:"Maximum"`
}

type Day struct {
	Icon 					int32    `json:"Icon"`
	IconPhrase 				string   `json:"IconPhrase"`
	HasPrecipitation 		bool   	 `json:"HasPrecipitation"`
	PrecipitationType 		string	 `json:"PrecipitationType"`
	PrecipitationIntensity  string	 `json:"PrecipitationIntensity"`
}

type Night struct {
	Icon 			 int32 		`json:"Icon"`        
	IconPhrase 		 string		`json:"IconPhrase"`
	HasPrecipitation bool		`json:"HasPrecipitation"`
}

type DailyForecast struct {
	Date 			string     		`json:"Date"`
	EpochDate 		int64           `json:"EpochDate"`
	Temperature 	Temperature		`json:"Temperature"`
	Day 			Day 			`json:"Day"`
	Night 			Night			`json:"Night"`
	Sources 		[]string		`json:"Sources"`
	MobileLink 		string			`json:"MobileLink"`
	Link 			string			`json:"Link"`
}

type Weather struct {
	Headline 	   Headline 	   `json:"Headline"`
	DailyForecasts []DailyForecast `json:"DailyForecasts"`
}	


func (w Weather) Print() {
	fmt.Println(w.Headline.EffectiveDate)
	fmt.Println(w.Headline.EffectiveDate)
	fmt.Println(w.Headline.EffectiveDate)
	fmt.Println(w.Headline.Text)
	fmt.Println(w.DailyForecasts[0].Temperature.Minimum.Value, "F")
 	fmt.Println(w.DailyForecasts[0].Temperature.Maximum.Value, "F")
}