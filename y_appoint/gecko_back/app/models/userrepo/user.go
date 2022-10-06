package userrepo

type User struct {
	ID       int64  `json:"id,omitempty" bun:"id"`
	Email    string `json:"email" bun:"email"`
	PassHash string `json:"pass_hash,omitempty" bun:"pass_hash"`
	UserRole int64  `json:"user_role" bun:"user_role"`
}
