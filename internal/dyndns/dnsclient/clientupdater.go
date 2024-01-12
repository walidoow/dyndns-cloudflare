package dnsclient

type DNSUpdater interface {
	Update(ip string) (string, error)
	DNSRecordIp() (string, error)
}
