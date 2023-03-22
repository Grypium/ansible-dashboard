package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-ldap/ldap/v3"
)

const (
	ldapServer   = "ldap.example.com"
	ldapPort     = 389
	ldapUserDN   = "ou=users,dc=example,dc=com"
	ldapGroupDN  = "ou=groups,dc=example,dc=com"
	ldapUserAttr = "uid"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || !authenticateLDAP(username, password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please provide your LDAP credentials"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func authenticateLDAP(username, password string) bool {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
	if err != nil {
		return false
	}
	defer l.Close()

	err = l.Bind(fmt.Sprintf("%s=%s,%s", ldapUserAttr, username, ldapUserDN), password)
	if err != nil {
		return false
	}

	return true
}
