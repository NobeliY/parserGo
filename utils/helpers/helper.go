package helpers

import "fmt"

func SprintF(module string, err error) string {
	return fmt.Sprintf("Module %s | err %s", module, err)
}