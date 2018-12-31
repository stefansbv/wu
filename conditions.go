/*
* conditions.go
*
* This file is part of wu.  It contains functions related to
* the -conditions switch (current conditions).
*
* Written and maintained by Stephen Ramsay <sramsay@protonmail.com>
* and Anthony Starks.
*
* Last Modified: Sat Aug 13 13:16:18 CDT 2016
*
* Copyright © 2010-2016 by Stephen Ramsay and Anthony Starks.
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
	"strconv"
	"strings"
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
	Wind_string          string
	Pressure_mb          string
	Pressure_in          string
	Pressure_trend       string
	Dewpoint_string      string
	Heat_index_string    string
	Windchill_string     string
	Visibility_mi        string
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
		//fmt.Println("   Temperature:", Convert(current.Temperature_string))
		fmt.Printf("   Temperature: %d (%d)\n", current.Temp_c, current.Temp_f)
	} else {
		fmt.Printf("   Temperature: %d (%d)\n", current.Temp_f, current.Temp_c)
	}
	if current.Heat_index_string != "NA" {
		fmt.Println("   Heat Index: ", current.Heat_index_string)
	}
	fmt.Println("   Sky Conditions:", current.Weather)
	fmt.Println("   Wind:", current.Wind_string)
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
		fmt.Print("   Dewpoint: ", Convert(current.Dewpoint_string))
	} else {
		fmt.Print("   Dewpoint: ", current.Dewpoint_string)
	}

	dp_components := strings.Split(current.Dewpoint_string, " ")
	dp, _ := strconv.Atoi(dp_components[0])
	if dp < 50 {
		fmt.Println(" (dry)")
	} else if dp >= 50 && dp <= 54 {
		fmt.Println(" (very comfortable)")
	} else if dp >= 55 && dp <= 59 {
		fmt.Println(" (comfortable)")
	} else if dp >= 60 && dp <= 64 {
		fmt.Println(" (okay for most)")
	} else if dp >= 65 && dp <= 69 {
		fmt.Println(" (somewhat uncomfortable)")
	} else if dp >= 70 && dp <= 74 {
		fmt.Println(" (very humid)")
	} else if dp >= 75 && dp <= 80 {
		fmt.Println(" (oppressive)")
	} else if dp >= 80 {
		fmt.Println(" (dangerously high)")
	}
	if current.Windchill_string != "NA" {
		fmt.Println("   Windchill: ", current.Windchill_string)
	}
	fmt.Printf("   Visibility: %s miles\n", current.Visibility_mi)
	if m, _ := regexp.MatchString("0.0", current.Precip_today_string); !m {
		fmt.Println("   Precipitation today: ", current.Precip_today_string)
	}
}
