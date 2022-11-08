// Provide a quick & easy way to load & use your configuration
package gonf

// Retrieve a single value
func Get(key string) string {
	return GetMany([]string{key})[0]
}

// Retrieve many values
func GetMany(key []string) []string {
	return []string{""}
}

// Retrieve all values
func GetAll() []string {
	//base := fetchFromFile(buildBaseFilename(), []string{})
	//env := fetchFromFile(buildEnvFilename(), []string{})
	return []string{""}
}
