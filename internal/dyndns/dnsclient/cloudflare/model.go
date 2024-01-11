package cloudflare

type DNSRecord struct {
	Content string   `json:"content"`
	Name    string   `json:"name"`
	Proxied bool     `json:"proxied,omitempty"`
	Type    string   `json:"type"`
	Comment string   `json:"comment,omitempty"`
	Tag     []string `json:"tag,omitempty"`
	Ttl     int      `json:"ttl,omitempty"`
}
