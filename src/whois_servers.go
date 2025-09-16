/*
   Copyright (C) 2025 Rodolfo González González.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package main

import (
	"encoding/json"
	"os"

	"golang.org/x/net/publicsuffix"
)

//-----------------------------------------------------------------------------

type WhoisServers struct {
	Servers map[string]string
}

// See README.md file in the data directory for more information.
func NewWhoisServers(servers string) *WhoisServers {
	data, err := os.ReadFile(servers)
	if err != nil {
		return &WhoisServers{}
	}

	var serverList map[string]string
	err = json.Unmarshal(data, &serverList)
	if err != nil {
		return &WhoisServers{}
	}

	return &WhoisServers{
		Servers: serverList,
	}
}

// GetWhoisServer finds the WHOIS server for a given domain
func (w *WhoisServers) GetWhoisServer(domain string) (string, bool) {
	// Get the public suffix for the given domain
	tld, _ := publicsuffix.PublicSuffix(domain)

	// Return the server, if any, from the list, or exists=false if
	// it was not found
	server, exists := w.Servers[tld]

	return server, exists
}
