package rest

import (
	"fmt"
	"net/http"
)

// Role definse role type
type Role string

// Hard coded role for now
const (
	AdminRole  Role = "Admin"
	MemberRole Role = "Member"
)

// AdminOnly middleware protects admin only enpoints
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("x-role")
		fmt.Println(role)
		if role == "" || role != string(AdminRole) {
			ServeJSON(http.StatusForbidden, "Un authorized", nil, nil, w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
