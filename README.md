[![License](https://img.shields.io/badge/GPL_v3.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![GitHub all releases](https://img.shields.io/github/downloads/rgglez/nagios-check-domain-ai/total)
![GitHub issues](https://img.shields.io/github/issues/rgglez/nagios-check-domain-ai)
![GitHub commit activity](https://img.shields.io/github/commit-activity/y/rgglez/nagios-check-domain-ai)
[![Go Report Card](https://goreportcard.com/badge/github.com/rgglez/nagios-check-domain-ai)](https://goreportcard.com/report/github.com/rgglez/nagios-check-domain)
[![GitHub release](https://img.shields.io/github/release/rgglez/nagios-check-domain-ai.svg)](https://github.com/rgglez/nagios-check-domain-ai/releases/)
![GitHub stars](https://img.shields.io/github/stars/rgglez/nagios-check-domain-ai?style=social)
![GitHub forks](https://img.shields.io/github/forks/rgglez/nagios-check-domain-ai?style=social)

# nagios-check-domain-ai

**check_domain_ai** is a Nagios plugin written in [Go](https://go.dev/) that checks the expiration date of a given domain and notifies you if it is about to expire.

This plugin queries public [WHOIS](https://www.rfc-editor.org/rfc/rfc3912.txt) servers. It strives to query the appropriate server for the public suffix or TLD. Then, it parses the WHOIS output using the Microsoft Azure AI gateway to OpenAI API.

## Command line options

* `--domain` `-D` string, the domain name to check.
* `--warn` `-w` integer, the number of days after which a warning will be considered a warning condition. Default: 30.
* `--crit` `-c` integer, the number of days after which a warning will be considered a critical condition. Default: 15.
* `--servers` `-s` string, the path to the file containing the list of WHOIS servers.

## Required enviroment variables

You must get this values from the Azure AI Foundry console:

* `AZURE_OPENAI_KEY` your key .
* `AZURE_OPENAI_ENDPOINT` the "Target URL" of endpoint to be used.
* `AZURE_OPENAI_MODEL` the name of the model you deployed. `gpt-4o-mini` works fine and it is not too expensive.

## Build and installation

### Build

* Get the source code:

```bash
$ git clone https://github.com/rgglez/nagios-check-domain-ai.git
```

* Compile the code:

```bash
$ cd nagios-check-domain-ai
$ make build
```

### Installation

To install the binary to the default path (```/usr/local/nagios/libexec```), execute:

```bash
$ sudo make install
```

Or just copy the executable to your regular Nagios plugins directory.

## Execution

First, export the required enviroment variables.

Basic example using default server (whois.iana.org):

```bash
check_domain_ai -D example.com
OK: Domain will expire in 159 days|expires=2025-08-13
```

Using the `servers.json` file:

```bash
check_domain_ai -D example.com --servers=/path/to/servers.json
```

## Server list

A list of WHOIS servers is included in the [data/servers.json](data/servers.json) file.
This is a JSON file which has the [TLD](https://en.wikipedia.org/wiki/Top-level_domain)
as the key and the corresponding WHOIS server as the value.

You can provide your own file. See the command line options above.

## Dependencies

This program has the following external dependencies:

* [github.com/likexian/whois](https://github.com/likexian/whois)
* [github.com/spf13/pflag](https://github.com/spf13/pflag)
* [github.com/xorpaul/go-nagios](https://github.com/xorpaul/go-nagios)
* [github.com/ztrue/tracerr](https://github.com/ztrue/tracerr)
* [golang.org/x/net/publicsuffix](golang.org/x/net/publicsuffix)
* [Microsoft Azure OpenAI](https://azure.microsoft.com/es-mx/pricing/details/cognitive-services/openai-service/)

## Azure OpenAI setup

1. Sign in to [Azure Portal](https://portal.azure.com).
1. Create a Resource:
  * Search for "Azure OpenAI" in the Marketplace.
  * Click **Create** ans choose your subscription,
  resource group, and region (for instance, `West US 3`).
  * Set a **deployment name** (for instance, `gpt-4o-mini-whois`) and choose the model, for example `gpt-4o-mini`.
1. Wait for deployment. Once deployed, note your endpoint URL, API keys and deployment name.

## Notes

- Of course, I am not affiliated with Microsoft in any way. Use their services at your own account and risk. Alternatively, you may modify the program to use the API of your preferred service. You can also use OpenAI's API directly, without Azure, by making only minor changes.

- Be aware that you should be polite and run the checker once a week or so, because some WHOIS servers don't like to be queried too often and may block your IP address permanently.


## License

Copyright 2025 Rodolfo González González.

This program is licensed under the terms of the [GPL v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html). Please read the [LICENSE](LICENSE.md) file.
