package auth

//UserAuthentication export
type UserAuthentication struct {
	ID       string `json:"id" bson:"_id"`
	UserName string `json:"userName" bson:"userName"`
	Pass     string `json:"pass" bson:"pass"`
}

//Response export
type Response struct {
	Token string `json:"token" `
	IsLog bool   `json:"islog" `
}
