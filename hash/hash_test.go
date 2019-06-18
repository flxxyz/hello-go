package hash

import (
	"github.com/flxxyz/hello-go/hash"
	"testing"
)

func TestSHA1(t *testing.T) {
	t.Log(SHA1("0"))
}

func TestSHA512(t *testing.T) {
	t.Log(SHA512("0"))
}

func TestMD5(t *testing.T) {
	t.Log(MD5("0"))
}

func TestHMAC(t *testing.T) {
	t.Log(HMAC("md5", "123456", "0"))
	t.Log(HMAC("sha1", "123456", "0"))
	t.Log(HMAC("sha256", "123456", "0"))
	t.Log(HMAC("sha512", "123456", "0"))
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
