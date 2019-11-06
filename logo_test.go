package kara

import (
	"fmt"
	"testing"
)

func TestLogo(t *testing.T) {
	fmt.Println(GetLogo(":8080", ":8081"))
}
