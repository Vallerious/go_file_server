package stringutils

import "testing"

func ConstructPortStringShouldReturnThePortWithDoubleDotsInfront(t *testing.T) {
    result := ConstructPortString(9999)

    if result != ":9999" {
        t.Errorf("executing ConstructPortString(9999) failed expected %v, but got %v", ":9999", result)
    }

}

