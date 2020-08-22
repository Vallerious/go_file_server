package stringutils

import (
    "fmt"
)

func ConstructPortString(port uint16) string {
    return fmt.Sprintf(":%d", port)
}

