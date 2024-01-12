# DynDNS-Cloudflare

## Overview
DynDNS-Cloudflare is a dynamic DNS update tool designed to work with Cloudflare's API. This tool helps in automatically updating DNS records based on the current IP address of your system. It's particularly useful for systems with dynamic IP addresses that need to maintain a consistent domain name.

## Installation
Ensure you have Go installed on your system. You can then clone this repository and build the project using:

```bash
git clone https://github.com/walidoow/dyndns-cloudflare
cd dyndns-cloudflare/cmd/dyndns
go build -o ./../../dyndns-cloudflare && cd ../../
```

## Usage
Before running the application, configure the `config.yml` with your Cloudflare API token and zone identifier. The application can be run using the following command:

```bash
./dyndns-cloudflare
```

## Configuration
Modify the `config.yml` file inside the root directory to include your Cloudflare token and zone identifier. The format is as follows:

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

## Contributing
Contributions to the project are welcome. Please follow standard Git workflows for contributions.

## License
This project is licensed under MIT license. Please refer to the `LICENSE` file for more details.
