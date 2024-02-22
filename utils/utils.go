package utils

import (
	"crypto/rand"
	"fmt"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"

	"routekey/config"

	"github.com/google/uuid"
	qrcode "github.com/skip2/go-qrcode"
)

func RandomChars(length int) string {
	var chars []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("wrong charset length")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4))
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

func GenerateShortURL() string {
	cfg := config.ReadConfig()
	mrand.Seed(time.Now().Unix())

	return cfg.Words[mrand.Intn(len(cfg.Words))]
}

func IsValidURL(url string) bool {
	return true
}

func GenerateQRCode(content string) ([]byte, error) {
	var qrCode []byte
	var err error
	qrCode, err = qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return []byte{}, err
	}
	return qrCode, err
}

func GenerateUUID() string {
	uid := uuid.New()
	return uid.String()
}

func GenerateUUIDWithoutDashes() string {
	id := uuid.New()
	uuid := strings.Replace(id.String(), "-", "", -1)
	return uuid
}

func NewUUID() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

func GetExpireAt(expireIn string) time.Time {
	var expireAt time.Time

	tokens := strings.Split(expireIn, " ")
	if len(tokens) != 2 {
		return expireAt
	}
	value, err := strconv.Atoi(tokens[0])
	if err != nil {
		fmt.Println(err)
		return expireAt
	}
	unit := tokens[1]
	if value == 1 {
		switch unit {
		case "day":
			expireAt = time.Now().Add(time.Hour * 24)
		case "week":
			expireAt = time.Now().Add(time.Hour * 24 * 7)
		case "month":
			expireAt = time.Now().Add(time.Hour * 24 * 30)
		case "year":
			expireAt = time.Now().Add(time.Hour * 24 * 365)
		}
	}
	switch unit {
	case "seconds":
		expireAt = time.Now().Add(time.Duration(value) * time.Second)
	case "minutes":
		expireAt = time.Now().Add(time.Duration(value) * time.Minute)
	case "hours":
		expireAt = time.Now().Add(time.Duration(value) * time.Hour)
	case "days":
		expireAt = time.Now().Add(time.Duration(value) * 24 * time.Hour)
	case "weeks":
		expireAt = time.Now().Add(time.Duration(value) * 7 * 24 * time.Hour)
	case "months":
		expireAt = time.Now().Add(time.Duration(value) * 30 * 24 * time.Hour)
	case "years":
		expireAt = time.Now().Add(time.Duration(value) * 365 * 24 * time.Hour)
	}
	return expireAt
}

func GenerateTrackerImage() []byte {
	var image []byte = []byte{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x08, 0x06, 0x00, 0x00, 0x00, 0x1f, 0x15, 0xc4,
		0x89, 0x00, 0x00, 0x00, 0x0c, 0x49, 0x44, 0x41, 0x54, 0x78, 0x9c, 0x63, 0x00, 0x01, 0x00, 0x00,
		0x05, 0x00, 0x01, 0x0d, 0x0a, 0x2d, 0xb4, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae,
	}
	return image
}
