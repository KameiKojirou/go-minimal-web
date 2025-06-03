package utils

import (
	"net/http"
	"time"
)


func CreateCookie(w http.ResponseWriter)    {
	cookie := http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = "ro8BS6Hiivgzy8Xuu09JDjlNLnSLldY5"
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"
	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request) string {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		return ""
	}
	return cookie.Value
}