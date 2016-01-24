package main

import (
	cryptorand "crypto/rand"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gorilla/securecookie"
)

// I was going to call it cookieData, but data about a cookie is just the
// nutrition facts #HowDidIGetAProfessionalCodingJob
type nutritionFacts struct {
	MagicRandomness string
	AccessToken     string
}

func withLogin(handler func(c context)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := newContext(w, r)

		// Do some best-effort context-filling
		if nutFact, err := loadCookie(r); err == nil {
			c.magicToken = nutFact.MagicRandomness
			if nutFact.AccessToken != "" {
				if p, err := db.loadUser(userID(nutFact.AccessToken)); err == nil {
					c.p = p
				}
			}
		} else {
			// They don't have a cookie from us yet. Let's fix that
			r := genName(128)
			val := nutritionFacts{
				MagicRandomness: r,
			}
			if encoded, err := s.Encode("info", val); err == nil {
				cookie := &http.Cookie{
					Name:  "info",
					Value: encoded,
					Path:  "/",
				}
				http.SetCookie(w, cookie)
				c.magicToken = r
			}
		}
		handler(c)
	}
}

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genName(n int) string {
	b := make([]byte, n)
	r := rand.New(cryptoRandSource{})
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func initSecretz() (string, error) {
	if dat, err := ioutil.ReadFile("secretz"); err == nil {
		return strings.TrimSpace(string(dat)), nil
	} else {
		return "", err
	}
}

func initKeys() (*securecookie.SecureCookie, error) {
	var hashKey []byte
	var blockKey []byte

	if dat, err := loadOrGenKey("hashKey"); err != nil {
		return nil, err
	} else {
		hashKey = dat
	}

	if dat, err := loadOrGenKey("blockKey"); err != nil {
		return nil, err
	} else {
		blockKey = dat
	}

	return securecookie.New(hashKey, blockKey), nil
}

func loadOrGenKey(name string) ([]byte, error) {
	if f, err := ioutil.ReadFile(name); err != nil {
		if dat := securecookie.GenerateRandomKey(32); dat != nil {
			if err := ioutil.WriteFile(name, dat, 0777); err == nil {
				return dat, nil
			}
			return nil, errors.New("Error writing file")
		}
		return nil, errors.New("Failed to generate key")
	} else {
		return f, nil
	}
}

func loadCookie(r *http.Request) (nutritionFacts, error) {
	if cookie, err := r.Cookie("info"); err == nil {
		value := nutritionFacts{}
		if err = s.Decode("info", cookie.Value, &value); err != nil {
			return nutritionFacts{}, err
		}
		return value, nil
	} else {
		return nutritionFacts{}, err
	}
}

type cryptoRandSource struct{}

func (cryptoRandSource) Int63() int64 {
	var buf [8]byte
	_, err := cryptorand.Read(buf[:])
	if err != nil {
		panic(err)
	}
	return int64(buf[0]) |
		int64(buf[1])<<8 |
		int64(buf[2])<<16 |
		int64(buf[3])<<24 |
		int64(buf[4])<<32 |
		int64(buf[5])<<40 |
		int64(buf[6])<<48 |
		int64(buf[7]&0x7f)<<56
}

func (cryptoRandSource) Seed(int64) {}
