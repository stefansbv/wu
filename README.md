
Wu
==========

Version 3.10.3

_wu_ is a small, fast command-line application that retrieves weather data from [Weather Underground](http://www.wunderground.com).

Update (5 July 2018)
--------------------

Greetings _wu_ users!  I use _wu_ every single day, and I try to deal with issues that come up.  It hasn't been updated in a while, but mostly because it appears to be bug free and is feature complete.

There is, however, a problem building _wu_ on MacOS using any version of the Go compiler after 1.9.  The problem appears to be ultimately related to MacOS itself, but from the standpoint of someone trying to compile from source, it has to do with cgo (the tool used to create Go programs that link to C libraries).  If you compile Go without cgo, the problem goes away.  However, Go compiled _with_ cgo is the default on most platforms.

Some time in the next few months, I intend to create a new version of _wu_ using the [Nim programming language](http://nim-lang.org).  This new version will likely match _wu's_ current feature set exactly, though I may add a switch or two.  I am also considering changing the format of the configuration file (the JSON format trips up a lot of users, because a stray comma breaks the program).  

However, it will be a long time before I deprecate this version.  Nim has not reached 1.0, and this version is working fine for the vast majority of users.  It's also likely that the problem with Go will eventually get resolved.  Right now, though, I am uncomfortable with maintaining a program that can't be compiled with the latest default version of the Go compiler on all platforms -- in part, because this means _wu_ isn't keeping pace with security updates to the Go compiler itself.

I realize that a lot of people use _wu_ -- some, for fairly mission-critical purposes.  I am writing this update primarily to assure users that this program is still being actively maintained, to acknowledge this hiccup with Go, and to indicate that there may be another (nearly identical) version of the program available before the year is out.

If *you* have any insight into what is going on with cgo and the crypto/x509 library, I'd love to hear from you.  Otherwise, please continue to enjoy _wu_, report any bugs you see, and experience lovely/interesting weather wherever you are.

Steve

Description
-----------

To use _wu,_ you need to obtain an API key from Weather Underground [http://www.wunderground.com/weather/api/](http://www.wunderground.com/weather/api/).  You should then add that key, the name of your default weather station, and your preference for Fahrenheit or Celcius (new feature in 3.10!) to $HOME/.condrc.  E.g.:

	{
	  "key": "YOUR_API_KEY",
	  "station": "Lincoln, NE",
		"degrees": "F"
	}

(the above is available in the wu root directory as "condrc")

wu has the following major options:

* `--conditions` reports the current weather conditions.

* `--forecast` gives the current (3-day) forecast.

* `--forecast10` gives the current (10-day) forecast.

* `--alerts` reports any active weather alerts.

* `--lookup [STATION]` allows you to determine the codes for the various weather stations in a particular area.  The format for STATION is the same as that for the -s switch below.

* `--astronomy` reports sunrise, sunset, and lunar phase.

* `--almanac` reports average high and low temperatures, as well as record temperatures for the day.

* `--yesterday` gives detailed almanac information for the previous day.

* `--history=YYYYMMDD` gives detailed almanac information for a given day.
* `--planner=MMDDMMDD` gives averages for travel planning (30-day max).
* `--tides` reports tidal data (when available).

* `--all` generate all reports (useful for creating custom reports and for mollifying the truly weather-crazed).
	
All twelve options can be accompanied by the -s switch, which can be used to override the default location in .condrc.  The argument passed to -s can be a "city, state-abbreviation/country", a (U.S. or Canadian) zip code, a 3- or 4-letter airport code, or "lat,long".

_wu_ also has two additional switches that provide information about the program:

* `--help`
* `--version`

By itself, the _wu_ command will show the current conditions.

Compiling and Installing Wu 
---------------------------

Mac users can install wu through the [Homebrew](http://brew.sh) system by typing:

	brew install wu

It can also be compiled from source.  Wu is written in the [Go programming language](http://golang.org/) (version 1.0 or later).  If you don't have a Go compiler, [you'll need to install one.](http://golang.org/doc/install.html).

To obtain the source code for wu:

    git clone git://github.com/sramsay/wu.git

To compile the wu executable, type:

    go build

To compile and install the excutable type:

    go install

(this will install it at the location specified by the GOPATH variable).

Wu should work on any system that can compile Go programs.

You may find the following aliases useful:

    alias conditions='wu'
    alias forecast='wu -forecast'
    alias forecast10='wu -forecast10'
    alias alerts='wu -alerts'
    alias astronomy='wu -astro'
    alias yesterday='wu -yesterday'
    alias almanac='wu -almanac'
    alias tides='wu -tides'
    alias slookup='wu -lookup'

License(s)
---------

Wu is written and maintained by [Stephen Ramsay](http://stephenramsay.us/) (sramsay{at}protonmail{dot}com) and [Anthony Starks](http://mindchunk.blogspot.com/).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see [http://www.gnu.org/licenses/](http://www.gnu.org/licenses/).

Data courtesy of Weather Underground, Inc. (WUI) is subject to the [Weather Underground API Terms and Conditions of Use](http://www.wunderground.com/weather/api/d/terms.html).  The author of this software is not affiliated with WUI, and the software is neither sponsored nor endorsed by WUI.

Thanks
------

Wu was heavily inspired by Jeremy Stanley's [weather](http://fungi.yuggoth.org/weather/).  This is a lovely Python script that has more-or-less the same output format as _wu_.  I reimplemented the system because Stanley's had stopped working (for me) and I wanted a program that was faster.  I also wanted a system that takes advantage of Weather Underground's rich [JSON](http://www.json.org/) API.
