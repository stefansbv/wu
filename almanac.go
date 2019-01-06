/*
* almanac.go
*
* This file is part of wu.  It contains functions related to
* the -almanac switch (historical data).
*
* Written and maintained by Stephen Ramsay <sramsay@protonmail.com>
* and Anthony Starks.
*
* Last Modified: Thu Jan 03 16:39:23 CST 2019
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

type AlmanacConditions struct {
	Almanac Almanac
}

type Almanac struct {
	Temp_high Temp_high
	Temp_low  Temp_low
}

type Temp_high struct {
	Normal     Normal
	Record     Record
	Recordyear string
}

type Temp_low struct {
	Normal     Normal
	Record     Record
	Recordyear string
}

type Normal struct {
	F string
	C string
}

type Record struct {
	F string
	C string
}

// printAlmanac prints the Almanac for a given station to standard out
func PrintAlmanac(obs *AlmanacConditions, stationId string, degrees string) {

	normalHighF := obs.Almanac.Temp_high.Normal.F
	normalHighC := obs.Almanac.Temp_high.Normal.C
	normalLowF := obs.Almanac.Temp_low.Normal.F
	normalLowC := obs.Almanac.Temp_low.Normal.C

	recordHighF := obs.Almanac.Temp_high.Record.F
	recordHighC := obs.Almanac.Temp_high.Record.C
	recordHYear := obs.Almanac.Temp_high.Recordyear
	recordLowF := obs.Almanac.Temp_low.Record.F
	recordLowC := obs.Almanac.Temp_low.Record.C
	recordLYear := obs.Almanac.Temp_low.Recordyear

	if degrees == "C" {
		fmt.Printf("Normal High: %s\u00B0C (%s\u00B0F)\n", normalHighC, normalHighF)
		fmt.Printf("Record High: %s\u00B0C (%s\u00B0F) [%s]\n", recordHighC, recordHighF, recordHYear)
		fmt.Printf("Normal Low : %s\u00B0C (%s\u00B0F)\n", normalLowC, normalLowF)
		fmt.Printf("Record Low : %s\u00B0C (%s\u00B0F) [%s]\n", recordLowC, recordLowF, recordLYear)

	} else {
		fmt.Printf("Normal High: %s\u00B0F (%s\u00B0C)\n", normalHighF, normalHighC)
		fmt.Printf("Record High: %s\u00B0F (%s\u00B0C) [%s]\n", recordHighF, recordHighC, recordHYear)
		fmt.Printf("Normal Low : %s\u00B0F (%s\u00B0C)\n", normalLowF, normalLowC)
		fmt.Printf("Record Low : %s\u00B0F (%s\u00B0C) [%s]\n", recordLowF, recordLowC, recordLYear)
	}

}
