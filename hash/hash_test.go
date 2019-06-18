package hash

import "testing"

func TestSHA1(t *testing.T) {
	t.Log(SHA1("0"))
}

func TestSHA512(t *testing.T) {
	t.Log(SHA512("0"))
}

func TestMD5(t *testing.T) {
	t.Log(MD5("0"), MD5("1"), MD5("2"), MD5("3"))
}

func TestHMAC(t *testing.T) {
	t.Log(HMAC("123456", "0"))
}

func TestBase64Encode(t *testing.T) {
	str := "hello world"
	t.Logf("old str = %s, new str = %s", str, Base64Encode([]byte(str)))
}

func TestBase64Decode(t *testing.T) {
	str := "aGVsbG8gd29ybGQ="
	data, _ := Base64Decode(str)
	t.Logf("old str = %s, new str = %s", str, string(data))
}