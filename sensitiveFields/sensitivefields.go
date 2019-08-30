package sensitiveFields

type User struct {
	Username  string `json:"username"`
	Password  string `json:"-"` // Note that '-' will make json marshaler completely ignore this field.
	Email     string `json:"email"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
}

