package utils

import (
	"fmt"
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	fmt.Println(HashAndSalt("123456"))
}

func TestComparePasswords(t *testing.T) {
	fmt.Println(ComparePasswords("$2a$04$vU6ZqEedbxHjU2yJdhooX.YYYjGXAUvXYGMzUk3lGmhlpSKqzrGwa", "123456"))
}
