package credentials

const (
	mongoHost = "localhost:27017"
)

var credentials = map[string]string{
	"mongoUri": "mongodb://" + mongoHost + "/friends",
}

// GetCredential function to retrieve credentials
func GetCredential(credentialKey string) string {
	return credentials[credentialKey]
}
