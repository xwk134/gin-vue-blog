package utils

import (
	"fmt"
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	fmt.Println(HashAndSalt("123456"))
}

func TestComparePasswords(t *testing.T) {
	fmt.Println(ComparePasswords("$2a$04$gPnTUpke58GCaX692v2pae2ccXh4evycKGgvSCZIJDBCmzL79M3UO", "123456"))
}
