# DynDNS-Cloudflare
[![Go build and test](https://github.com/walidoow/dyndns-cloudflare/actions/workflows/golang-build.yml/badge.svg)](https://github.com/walidoow/dyndns-cloudflare/actions/workflows/golang-build.yml) [![Create Release on Tag Push](https://github.com/walidoow/dyndns-cloudflare/actions/workflows/create-release.yml/badge.svg)](https://github.com/walidoow/dyndns-cloudflare/actions/workflows/create-release.yml)
## Overview

DynDNS-Cloudflare is a dynamic DNS update tool designed to work with Cloudflare's API. This tool helps in automatically
updating DNS records based on the current IP address of your system. It's particularly useful for systems with dynamic
IP addresses that need to maintain a consistent domain name.

## Installation

Ensure you have Go installed on your system. You can then clone this repository and build the project using:

```bash
git clone https://github.com/walidoow/dyndns-cloudflare
cd dyndns-cloudflare
go build ./cmd/dyndns/
```

## Usage

Before running the application, configure the `config.yml` with your Cloudflare API token and zone identifier. The
application can be run using the following command (add .exe for Windows):

```bash
./dyndns
```

## Configuration

Modify the `config.yml` file inside the root directory to include your Cloudflare token and zone identifier. The format
is as follows:

```yaml
token: your_cloudflare_api_token # Your Cloudflare API token for authentication.
zone-identifier: your_zone_identifier # The Zone ID of your domain on Cloudflare.
cloudflare-dns-record:
  # The required fields are the identifier, content, name, and type.
  identifier: your_dns_record_identifier # The Identifier for the specific DNS record you want to dnsclient.
  name: your_domain # The domain name for the DNS record.
  type: A # The type of DNS record (e.g., A, AAAA, CNAME, etc.).
  proxied: false # Whether the record is receiving the performance and security benefits of Cloudflare.
  comment: your_comment # Comments or notes about the DNS record. This field has no effect on DNS responses.
  tags: [ tag1, tag2 ] # Custom tags for the DNS record. This field has no effect on DNS responses.
  ttl: 1 # Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'. Value must be between 60 and 86400, with the minimum reduced to 30 for Enterprise zones.
```

## Run workflows locally using ACT

This project uses [act](https://github.com/nektos/act) to run GitHub Actions locally.
To run the workflows locally, use the following command:

```bash
  act -e event.json
```

You can also specify a specific workflow to run using the `-j` flag:

```bash
  act -j <workflow_name> -e event.json
```

or simulate an event like a pull_request or a push:

```bash
  act -e event.json pull_request
```

The `event.json` file is used to simulate the event payload sent by GitHub to the runner and skip some steps that we
only want to be executed in production environment.

## Contributing

Contributions to the project are welcome. Please follow standard Git workflows for contributions.

## License

This project is licensed under MIT license. Please refer to the `LICENSE` file for more details.
