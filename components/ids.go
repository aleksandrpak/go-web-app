package components

import (
	"fmt"
	"strings"
)

const (
	counterID    = "counter"
	errorToastID = "error-toast"
)

func idTarget(id string, mods ...string) string {
	return fmt.Sprintf("#%s", strings.Join(append([]string{id}, mods...), " "))
}
