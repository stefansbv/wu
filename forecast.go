/*
* forecast.go
*
* This file is part of wu.  It contains functions related to
* the -forecast switch (3-day forecast).
*
* Written and maintained by Stephen Ramsay <sramsay@protonmail.com>
* and Anthony Starks.
*
* Last Modified: Thu Jan 03 15:37:23 CST 2019
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
)

type ForecastConditions struct {
	Forecast Forecast
}

type Forecast struct {
	Txt_forecast Txt_forecast
}

type Txt_forecast struct {
	Date        string
	Forecastday []Forecastday
}

type Forecastday struct {
	Title          string
	Fcttext        string
	Fcttext_metric string
}

// printForecast prints the forecast for a given station to standard out
func PrintForecast(obs *ForecastConditions, stationId string, degrees string) {
	t := obs.Forecast.Txt_forecast
	fmt.Printf("Forecast for %s\n", stationId)
	fmt.Printf("Issued at %s\n", t.Date)
	for _, f := range t.Forecastday {
		if degrees == "C" {
			fmt.Printf("%s: %s\n", f.Title, f.Fcttext_metric)
		} else {
			fmt.Printf("%s: %s\n", f.Title, f.Fcttext)
		}
	}
}
