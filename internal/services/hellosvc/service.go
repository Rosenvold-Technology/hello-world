
package hellosvc

import "fmt"

func GetGreeting(name string) string {
	if name == "" {
		name = "Kubernetes Adventurer"
	}
	return fmt.Sprintf("Hello, %s! 👋", name)
}
