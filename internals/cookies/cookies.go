package cookies

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
)

// should base64 encode the cookies so that non ASCII characters don't get cut
func WriteCookies(Cookie http.Cookie, w http.ResponseWriter) error {
	Cookie.Value = base64.URLEncoding.EncodeToString([]byte(Cookie.Value))

	if len(Cookie.String()) > 4096 {
		return errors.New("cookie value too long")
	}

	http.SetCookie(w, &Cookie)

	return nil
}

func ReadCookies(r *http.Request, CookieName string) (string, error) {
	Cookie, err := r.Cookie(CookieName)

	if err != nil {
		return "", err
	}

	fmt.Printf("Cookie name: %s, Cookie value: %s\n", CookieName, Cookie.Value)

	// Decode the cookie value:
	Value, err := base64.URLEncoding.DecodeString(Cookie.Value)

	if err != nil {
		return "", errors.New("couldn't decode the cookie back")
	}

	return string(Value), nil
}
