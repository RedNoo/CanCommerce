package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Payload struct {
	Sub  string
	Name string
	Exp  time.Time
}

func Base64Encode(src string) string {
	return base64.URLEncoding.EncodeToString([]byte(src))

}

func Base64Decode(src string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(src)
	if err != nil {
		errMsg := fmt.Errorf("Decoding Error %s", err)
		return "", errMsg
	}
	return string(decoded), nil
}

func Hash(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func isValidHash(value string, hash string, secret string) bool {
	return hash == Hash(value, secret)
}

func Encode(payload Payload, secret string) string {
	type Header struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}

	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	strHeader, _ := json.Marshal(header)
	base64Header := Base64Encode(string(strHeader))
	strPayload, _ := json.Marshal(payload)
	base64Payload := Base64Encode(string(strPayload))
	signatureValue := base64Header + "." + base64Payload
	return signatureValue + "." + Hash(signatureValue, secret)
}

func Decode(jwt string, secret string) (interface{}, error) {
	token := strings.Split(jwt, ".")

	if len(token) != 3 {
		splitErr := errors.New("Invalid token")
		return nil, splitErr
	}

	decodedPayload, PayloadErr := Base64Decode(token[1])
	if PayloadErr != nil {
		return nil, fmt.Errorf("Invalid payload: %s", PayloadErr.Error())
	}

	payload := Payload{}

	ParseErr := json.Unmarshal([]byte(decodedPayload), &payload)
	if ParseErr != nil {
		return nil, fmt.Errorf("Invalid payload: %s", ParseErr.Error())
	}

	if time.Now().Unix() > payload.Exp.Unix() {
		return nil, errors.New("Expired token")
	}
	signatureValue := token[0] + "." + token[1]

	if isValidHash(signatureValue, token[2], secret) == false {
		return nil, errors.New("Invalid token")
	}
	return payload, nil
}
