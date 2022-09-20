package auth

import "testing"

func TestEncrypt(t *testing.T) {
	soucre := "admind1234"
	str, err := Encrypt(soucre)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

func TestCompare(t *testing.T) {
	soucre := "admind1234"
	hashedPass := "$2a$10$7LiNRYlx839YxuLM.d51Bub8q2SlJV4N7OuK3RPVekHXUA07m5uNi"
	err := Compare(hashedPass, soucre)
	if err != nil {
		t.Fatal(err)
	}
}
