package middleware

import (
	"fmt"
	"net/http"
)

type CheckLoginMiddleware struct {
}

func NewCheckLoginMiddleware() *CheckLoginMiddleware {
	return &CheckLoginMiddleware{}

}

func (m *CheckLoginMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("走进来了没")
		next.ServeHTTP(w, r)
	}
}
