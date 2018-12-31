/*
* conditions.go
*
* This file is part of wu.  It contains functions related to
* the -conditions switch (current conditions).
*
* Written and maintained by Stephen Ramsay <sramsay@protonmail.com>
* and Anthony Starks.
*
* Last Modified: Mon Dec 31 16:26:56 CST 2018
*
* Copyright © 2010-2019 by Stephen Ramsay and Anthony Starks.
*
* wu is free software; you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation; either version 3, or (at your option)
* any later version.
*
* wu is distributed in the hope that it will be useful, but WITHOUT
* ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
* or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public
* License for more details.
*
* You should have received a copy of the GNU General Public License
* along with wu; see the file COPYING.  If not see
* <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"regexp"
)

type Conditions struct {
	Current_observation Current
}

type Current struct {
	Observation_time     string
	Observation_location Location
	Station_id           string
	Weather              string
	Temp_f							 int
	Temp_c							 int
	Relative_humidity    string
	Wind_dir						 string
	Wind_degrees				 int
	Wind_mph						 int
	Wind_gust_mph				 string
	Wind_kph						 int
	Wind_gust_kph				 string
	Pressure_mb          string
	Pressure_in          string
	Pressure_trend       string
	Dewpoint_f					 int
	Dewpoint_c					 int
	Heat_index_string    string
	Windchill_string		 string
	Windchill_f					 string
	Windchill_c					 string
	Visibility_mi        string
	Visibility_km        string
	Feelslike_f					 string
	Feelslike_c					 string
	Precip_today_string  string
}

type Location struct {
	Full string
}

// printConditions prints the conditions to standard output
func PrintConditions(obs *Conditions, degrees string) {
	current := obs.Current_observation
	fmt.Printf("Current conditions at %s (%s)\n%s\n",
		current.Observation_location.Full, current.Station_id, current.Observation_time)
	if degrees == "C" {
		fmt.Printf("   Temperature: %d° C (%d° F)\n", current.Temp_c, current.Temp_f)
	} else {
		fmt.Printf("   Temperature: %d° F (%d° C)\n", current.Temp_f, current.Temp_c)
	}
	if current.Heat_index_string != "NA" {
		fmt.Println("   Heat Index: ", current.Heat_index_string)
	}
	fmt.Println("   Sky Conditions:", current.Weather)

	var windstring = ""
	if degrees == "C" {
		windstring = fmt.Sprintf("   Wind: From the %s (%d°) at %d km/h", current.Wind_dir, current.Wind_degrees, current.Wind_kph)
		if (current.Wind_gust_kph != "0") {
			windstring = fmt.Sprintf("%s gusting to %s km/h", windstring, current.Wind_gust_kph)
		}
	} else {
		windstring = fmt.Sprintf("   Wind: From the %s (%d°) at %d mph", current.Wind_dir, current.Wind_degrees, current.Wind_mph)
		if (current.Wind_gust_mph != "0") {
			windstring = fmt.Sprintf("%s gusting to %s mph", windstring, current.Wind_gust_mph)
		}
	}

	fmt.Println(windstring)

	var pstring = ""
	if degrees == "C" {
		pstring = fmt.Sprintf("   Pressure: %s mb (%s in) and", current.Pressure_mb, current.Pressure_in)
	} else {
		pstring = fmt.Sprintf("   Pressure: %s in (%s mb) and", current.Pressure_in, current.Pressure_mb)
	}
	switch current.Pressure_trend {
	case "+":
		fmt.Println(pstring, "rising")
	case "-":
		fmt.Println(pstring, "falling")
	case "0":
		fmt.Println(pstring, "holding steady")
	}

	fmt.Println("   Relative humidity:", current.Relative_humidity)

	if degrees == "C" {
		fmt.Printf("   Dewpoint: %d° C", current.Dewpoint_c);
	} else {
		fmt.Printf("   Dewpoint: %d° F", current.Dewpoint_f);
	}

	switch dp := current.Dewpoint_f; {
	case dp < 50:
		fmt.Println(" (dry)")
	case dp >= 50 && dp <= 54:
		fmt.Println(" (very comfortable)")
	case dp >= 55 && dp <= 59:
		fmt.Println(" (comfortable)")
	case dp >= 60 && dp <= 64:
		fmt.Println(" (okay for most)")
	case dp >= 65 && dp <= 69:
		fmt.Println(" (somewhat uncomfortable)")
	case dp >= 70 && dp <= 74:
		fmt.Println(" (very humid)")
	case dp >= 75 && dp <= 80:
		fmt.Println(" (oppressive)")
	case dp >= 80:
		fmt.Println(" (dangerously high)")
	}

	if current.Windchill_string != "NA" {
		if degrees == "C" {
			fmt.Printf("   Windchill: %s° C (%s° F)\n", current.Windchill_c, current.Windchill_f)
		} else {
			fmt.Printf("   Windchill: %s° F (%s° C)\n", current.Windchill_f, current.Windchill_c)
		}
	}

	if degrees == "C" {
		fmt.Printf("   Visibility: %s kilometers\n", current.Visibility_km)
	} else {
		fmt.Printf("   Visibility: %s miles\n", current.Visibility_mi)
	}

	if m, _ := regexp.MatchString("0.0", current.Precip_today_string); !m {
		fmt.Println("   Precipitation today: ", current.Precip_today_string)
	}
}
