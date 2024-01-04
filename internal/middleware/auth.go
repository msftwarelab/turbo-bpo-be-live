package middleware

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"strings"

	err "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	pkg_strings "github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/token"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var UserCtxKey = &contextKey{"userId"}
var ipAddressKey = &contextKey{"ipAddress"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			if len(splitToken) != 2 {
				next.ServeHTTP(w, r)
				return
			}
			reqToken = splitToken[1]
			claims, err := token.Decode(reqToken)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			// put it in context
			ctx := context.WithValue(r.Context(), UserCtxKey, claims)
			//and call the next with our new context
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func ReadUserIP(req *http.Request) (*string, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, err
	}
	ip = strings.TrimSpace(ip)

	return pkg_strings.ToObject(ip), nil
}

func LoginLog() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userIPAddress, _ := ReadUserIP(r)

			// if err != nil {
			// 	next.ServeHTTP(w, r)
			// 	return
			// }
			// put it in context
			ctx := context.WithValue(r.Context(), ipAddressKey, userIPAddress)
			//and call the next with our new context
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func GetUserCtx(ctx context.Context) (*token.Claims, error) {
	raw, ok := ctx.Value(UserCtxKey).(*token.Claims)
	if ok {
		return raw, nil
	}
	return nil, err.InvalidToken

}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func GetUserIPAddress(ctx context.Context) string {
	ipaddressRaw := ctx.Value(ipAddressKey)
	return PrettyPrint(ipaddressRaw)

}
