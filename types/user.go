package types

type User struct {
	ID       int    `bson:"_id", json:"id"`
	Username string `bson: "username", json:"username"`
	Email    string `bson: "email", json:"email"`
}
