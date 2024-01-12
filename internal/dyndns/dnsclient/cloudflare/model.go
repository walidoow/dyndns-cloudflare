package cloudflare

type DNSRecordPayload struct {
	Content string   `json:"content"`
	Name    string   `json:"name"`
	Proxied bool     `json:"proxied,omitempty"`
	Type    string   `json:"type"`
	Comment string   `json:"comment,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Ttl     int      `json:"ttl,omitempty"`
}

type ApiGetRecordResponse struct {
	Result   Result        `json:"result"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}

type Result struct {
	ID        string `json:"id"`
	ZoneID    string `json:"zone_id"`
	ZoneName  string `json:"zone_name"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Proxiable bool   `json:"proxiable"`
	Proxied   bool   `json:"proxied"`
	Ttl       int    `json:"ttl"`
	Locked    bool   `json:"locked"`
	Meta      struct {
		AutoAdded           bool `json:"auto_added"`
		ManagedByApps       bool `json:"managed_by_apps"`
		ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
	} `json:"meta"`
	Comment    string   `json:"comment"`
	Tags       []string `json:"tags"`
	CreatedOn  string   `json:"created_on"`
	ModifiedOn string   `json:"modified_on"`
}
