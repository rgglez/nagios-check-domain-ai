/*
Copyright 2025 Rodolfo González González

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
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
