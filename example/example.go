package example

import (
	"fmt"
	"time"

	calc "github.com/MSA-Software-LLC/adhan-go/pkg/calc"
	data "github.com/MSA-Software-LLC/adhan-go/pkg/data"
	util "github.com/MSA-Software-LLC/adhan-go/pkg/util"
)

func CalculateNewYorkPrayers() {
	date := data.NewDateComponents(time.Date(2015, time.Month(7), 12, 0, 0, 0, 0, time.UTC))
	params := calc.GetMethodParameters(calc.NORTH_AMERICA)
	params.Madhab = calc.HANAFI

	coords, err := util.NewCoordinates(35.7750, -78.6336)
	if err != nil {
		fmt.Printf("got error %+v", err)
		return
	}

	prayerTimes, err := calc.NewPrayerTimes(coords, date, params)
	if err != nil {
		fmt.Printf("got error %+v", err)
		return
	}

	err = prayerTimes.SetTimeZone("America/New_York")
	if err != nil {
		fmt.Printf("got error %+v", err)
		return
	}

	fmt.Printf("Fajr: %+v\n", prayerTimes.Fajr)       // Fajr: 2015-07-12 04:42:00 -0400 EDT
	fmt.Printf("Sunrise: %+v\n", prayerTimes.Sunrise) // Sunrise: 2015-07-12 06:08:00 -0400 EDT
	fmt.Printf("Dhuhr: %+v\n", prayerTimes.Dhuhr)     // Dhuhr: 2015-07-12 13:21:00 -0400 EDT
	fmt.Printf("Asr: %+v\n", prayerTimes.Asr)         // Asr: 2015-07-12 18:22:00 -0400 EDT
	fmt.Printf("Maghrib: %+v\n", prayerTimes.Maghrib) // Maghrib: 2015-07-12 20:32:00 -0400 EDT
	fmt.Printf("Isha: %+v\n", prayerTimes.Isha)       // Isha: 2015-07-12 21:57:00 -0400 EDT
}

func CalculateNewYorkPrayersWithBuilder() {
	date := data.NewDateComponents(time.Date(2015, time.Month(7), 12, 0, 0, 0, 0, time.UTC))
	params := calc.NewCalculationParametersBuilder().
		SetMadhab(calc.HANAFI).
		SetMethod(calc.NORTH_AMERICA).
		SetFajrAngle(15.0).
		SetIshaAngle(15.0).
		SetMethodAdjustments(calc.PrayerAdjustments{
			DhuhrAdj: 1,
		}).
		Build()

	coords, err := util.NewCoordinates(35.7750, -78.6336)
	if err != nil {
		fmt.Printf("got error %+v", err)
		return
	}

	prayerTimes, err := calc.NewPrayerTimes(coords, date, params)
	if err != nil {
		fmt.Printf("got error %+v", err)
		return
	}

	err = prayerTimes.SetTimeZone("America/New_York")
	if err != nil {
		fmt.Printf("got error %+v", err)
		return
	}

	fmt.Printf("Fajr: %+v\n", prayerTimes.Fajr)       // Fajr: 2015-07-12 04:42:00 -0400 EDT
	fmt.Printf("Sunrise: %+v\n", prayerTimes.Sunrise) // Sunrise: 2015-07-12 06:08:00 -0400 EDT
	fmt.Printf("Dhuhr: %+v\n", prayerTimes.Dhuhr)     // Dhuhr: 2015-07-12 13:21:00 -0400 EDT
	fmt.Printf("Asr: %+v\n", prayerTimes.Asr)         // Asr: 2015-07-12 18:22:00 -0400 EDT
	fmt.Printf("Maghrib: %+v\n", prayerTimes.Maghrib) // Maghrib: 2015-07-12 20:32:00 -0400 EDT
	fmt.Printf("Isha: %+v\n", prayerTimes.Isha)       // Isha: 2015-07-12 21:57:00 -0400 EDT
}

func CalculateLondonPrayers() {
	date := data.NewDateComponents(time.Date(2025, time.June, 24, 0, 0, 0, 0, time.UTC))
	params := calc.GetMethodParameters(calc.NORTH_AMERICA)
	params.Madhab = calc.HANAFI
	params.HighLatitudeRule = calc.TWILIGHT_ANGLE // ANGLE_BASED

	coords, err := util.NewCoordinates(51.5285262, -0.2664052)
	if err != nil {
		fmt.Printf("got error %+v", err)
		return
	}

	prayerTimes, err := calc.NewPrayerTimes(coords, date, params)
	if err != nil {
		fmt.Printf("got error %+v", err)
		return
	}

	err = prayerTimes.SetTimeZone("Europe/London")
	if err != nil {
		fmt.Printf("got error %+v", err)
		return
	}

	fmt.Printf("Fajr:    %v\n", prayerTimes.Fajr)
	fmt.Printf("Sunrise: %v\n", prayerTimes.Sunrise)
	fmt.Printf("Dhuhr:   %v\n", prayerTimes.Dhuhr)
	fmt.Printf("Asr:     %v\n", prayerTimes.Asr)
	fmt.Printf("Maghrib: %v\n", prayerTimes.Maghrib)
	fmt.Printf("Isha:    %v\n", prayerTimes.Isha)
}
