package user

// AccountIfo : struct for user
type AccountIfo struct {
	ID       string `bson,json:"id"`
	UserName string `bson,json:"userName"`
	Password string `bson,json:"password"`
}

type Config struct {
	CurrentDB     string `env:"CURRENT_DB" envDefault:"postgres"`
	PostgresDBURL string `env:"POSTGRES_DB_URL" envDefault:"postgres://postgres:user@localhost:5436/user?sslmode=disable"`
	MongoDBURL    string `env:"MONGO_DB_URL"`
	JwtKey        []byte `env:"JWT-KEY" `
}
