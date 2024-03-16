package backends

import "fmt"

type Headline struct {
	EffectiveDate 	   string		                                 `json:"EffectiveDate"`
	EffectiveEpochDate int64    	                                 `json:"EffectiveEpochDate"`
	Severity 		   int32                                         `json:"Severity"`
	Text 			   string	   	                                 `json:"Text"`
	Category 		   string		                                 `json:"Category"`
	EndDate 		   string		                                 `json:"EndDate"`
	EndEpochDate 	   int64                                         `json:"EndEpochDate"`
	MobileLink 		   string                                        `json:"MobileLink"`
	Link 			   string	                                     `json:"Link"`
}

type Direction struct {
    Degrees   float32                                                `json:"Degrees"`     
    Localized string                                                 `json:"Localized"`
    English   string                                                 `json:"English"`
}  // +++++++++++++++++++++++++

type Wind struct {
    Speed     MinMax                                                 `json:"Speed"`  
    Direction Direction                                              `json:"Direction"`  
}  // +++++++++++++++++++++++++

// Day/Night
type Period struct {
    Icon int32                                                       `json:"Icon"` 
    IconPhrase string                                                `json:"IconPhrase"` 
    HasPrecipitation bool                                            `json:"HasPrecipitation"` 
    ShortPhrase string                                               `json:"ShortPhrase"` 
    LongPhrase string                                                `json:"LongPhrase"` 
    PrecipitationProbability int32                                   `json:"PrecipitationProbability"` 
    ThunderstormProbability int32                                    `json:"ThunderstormProbability"` 
    RainProbability int32                                            `json:"RainProbability"` 
    SnowProbability int32                                            `json:"SnowProbability"` 
    IceProbability int32                                             `json:"IceProbability"` 
    Wind Wind                                                        `json:"Wind"` 
    WindGust Wind                                                    `json:"WindGust"` 
    TotalLiquid
    Rain
    Snow
    Ice
    HoursOfPrecipitation
    HoursOfRain
    HoursOfSnow
    HoursOfIce
    CloudCover
    Evapotranspiration Evapotranspiration
    SolarIrradiance SolarIrradiance
    RelativeHumidity RelativeHumidity
    WetBulbTemperature WetBulbTemperature
    WetBulbGlobeTemperature WetBulbGlobeTemperature
}

type MinMax struct {
	Value    float32		                                         `json:"Value"`
	Unit     string                                                  `json:"Unit"`
	UnitType int32      	                                         `json:"UnitType"`
}  // +++++++++++++++++++++++++

type Temperature struct {
	Minimum	MinMax 	                                                 `json:"Minimum"`
	Maximum MinMax 	                                                 `json:"Maximum"`
}  // +++++++++++++++++++++++++

type Night struct {
	Icon 			 int32 		                                     `json:"Icon"`        
	IconPhrase 		 string		                                     `json:"IconPhrase"`
	HasPrecipitation bool		                                     `json:"HasPrecipitation"`
}

type Sun struct {
    Rise      string                                                 `json:"Rise"`
    EpochRise int64                                                  `json:"EpochRise"`
    Set       string                                                 `json:"Set"`
    EpochSet  int64                                                  `json:"EpochSet"`
} // +++++++++++++++++++++++++

type Moon struct {
    Rise      string                                                 `json:"Rise"`
    EpochRise int64                                                  `json:"EpochRise"`
    Set       string                                                 `json:"Set"`
    EpochSet  int64                                                  `json:"EpochSet"`
    Phase     string                                                 `json:"Phase"`
    Age       int32                                                  `json:"Age"`
} // +++++++++++++++++++++++++

type MinMaxRFT struct {
    Value    float32                                                 `json:"Value"`
    Unit     string                                                  `json:"Unit"`
    UnitType int32                                                   `json:"UnitType"`
    Phrase   string                                                  `json:"Phrase"`
} // +++++++++++++++++++++++++

type RealFeelTemperature struct {
    Minimum MinMaxRFT                                                `json:"Minimum"`
    Maximum MinMaxRFT                                                `json:"Maximum"`
} // +++++++++++++++++++++++++

type DegreeDaySummary struct {
    Heating MinMax                                                   `json:"Heating"`
    Cooling MinMax                                                   `json:"Cooling"`
} // +++++++++++++++++++++++++

type AirAndPollen struct {
    Name          string                                             `json:"Name"`
    Value         int32                                              `json:"Value"`
    Category      string                                             `json:"Category"`
    CategoryValue int32                                              `json:"CategoryValue"`
    Type          string                                             `json:"Type,omitempty"`
} // +++++++++++++++++++++++++

type DailyForecast struct {
  Date 				 	   string     			                     `json:"Date"`
  EpochDate 			   int64           		                     `json:"EpochDate"`
  Sun                      Sun                                       `json:"Sun"`
  Moon 		               Moon                                      `json:"Moon"`
  Temperature 	           Temperature			                     `json:"Temperature"`
  RealFeelTemperature 	   RealFeelTemperature 	                     `json:"RealFeelTemperature"`
  RealFeelTemperatureShade RealFeelTemperature                       `json:"RealFeelTemperatureShade"`
  HoursOfSun     		   float32				                     `json:"HoursOfSun"`
  DegreeDaySummary         DegreeDaySummary		                     `json:"DegreeDaySummary"`
  AirAndPollens            []AirAndPollen                            `json:"AirAndPollen"`
  // Day                Period                `json:"Day"`
  // Night 					   Period					       `json:"Night"`
  Sources 				   []string				                     `json:"Sources"`
  MobileLink 				 string					                 `json:"MobileLink"`
  Link 					     string					                 `json:"Link"`
}

type Weather struct {
	Headline 	   Headline 	   `json:"Headline"`
	DailyForecasts []DailyForecast `json:"DailyForecasts"`
}	

func toCelc(t float32) int32 {
  tint := int32(t)
	return ((tint-32)*5)/9
}

func (w Weather) Echo() {
	fmt.Println("Brief desc:", w.Headline.Text)
	fmt.Printf("Temperature range: %v℃ -%v℃\n", 
		toCelc(w.DailyForecasts[0].Temperature.Minimum.Value), 
		toCelc(w.DailyForecasts[0].Temperature.Maximum.Value))
	fmt.Println("Daytime info:")
	fmt.Printf("\tStatus: %v\n\tHas precipitation: %t\n", w.DailyForecasts[0].Day.IconPhrase, w.DailyForecasts[0].Day.HasPrecipitation)
	fmt.Println("Nightime info:")
	fmt.Printf("\tStatus: %v\n\tHas precipitation: %t\n", w.DailyForecasts[0].Night.IconPhrase, w.DailyForecasts[0].Night.HasPrecipitation)
}
