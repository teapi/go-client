package teapi

// Configuration of the client
type Configuration struct {
	host   string
	key    string
	secret []byte
}

// Create a configuration with the specified host, key and secret
// Get these three values from Teapi's management console
func Configure(host, key, secret string) *Configuration {
	if l := len(host); l > 4 {
		if host[:4] != "http" {
			host = "https://" + host
			l = len(host)
		}
		if host[l-1] == '/' {
			host = host[:l-1]
		}
	}
	return &Configuration{
		host:   host,
		key:    key,
		secret: []byte(secret),
	}
}
