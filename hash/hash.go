package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

func SHA1(text string) string {
	ctx := sha1.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func SHA256(text string) string {
	ctx := sha256.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func SHA512(text string) string {
	ctx := sha512.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func HMAC(algo, key, data string) string {
	var ctx hash.Hash

	switch algo {
	case "md5":
		ctx = hmac.New(md5.New, []byte(key))
	case "sha1":
		ctx = hmac.New(sha1.New, []byte(key))
	case "sha256":
		ctx = hmac.New(sha256.New, []byte(key))
	case "sha512":
		ctx = hmac.New(sha512.New, []byte(key))
	}

	ctx.Write([]byte(data))
	return hex.EncodeToString(ctx.Sum([]byte(nil)))
}

func Base64Encode(text []byte) string {
	encodeString := base64.StdEncoding.EncodeToString(text)
	return encodeString
}

func Base64Decode(text string) ([]byte, error) {
	if encodeString, err := base64.StdEncoding.DecodeString(text); err != nil {
		return nil, err
	} else {
		return encodeString, nil
	}
}
