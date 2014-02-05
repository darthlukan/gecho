/*
Gecho is a Go implementation of the commonly used *nix command line tool 'echo'.

Copyright (C) 2014  Brian C. Tomlinson

Contact: brian.tomlinson@linux.com

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program; if not, write to the Free Software Foundation, Inc.,
51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
*/
package main

import (
	"flag"
	"os"
)

const (
	newLine = "\n"
)

var (
	version        string = "Gecho version 1.0 \nWritten by Brian Tomlinson <brian.tomlinson@linux.com>"
	noNewline      *bool  = flag.Bool("n", false, "Ignore trailing newline.")
	displayVersion *bool  = flag.Bool("version", false, "Output version information and exit.")
)

func main() {

	var str string

	flag.Parse()

	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			str += " "
		}
		str += flag.Arg(i)
	}

	if !*noNewline {
		str += newLine
	}

	if *displayVersion {
		str = version
	}

	os.Stdout.WriteString(str)
}
