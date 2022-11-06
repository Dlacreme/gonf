// Provide a quick & easy way to load & use your configuration
package gonf

import "fmt"

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
	root := fetchFromFile(buildRootFilename(), []string{})
	env := fetchFromFile(buildFilenameForEnv(), []string{})
	printCompare(root, env)
	return []string{""}
}

func printCompare(root []Set, env []Set) {
	fmt.Println("ROOT >")
	printArraySet(root)
	fmt.Println("ENV >")
	printArraySet(env)

}

func printArraySet(a []Set) {
	for v, k := range a {
		fmt.Printf("KEY[%s] VALUE[%s]\n", v, k)
	}
}
