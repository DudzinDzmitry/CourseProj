package user

// User : struct for user
type User struct {
	ID           string `bson,json:"id"`
	Name         string `bson,json:"name"`
	Position     string `bson,json:"position"`
	Password     string `bson,json:"password"`
	RefreshToken string `bson,json:"refreshToken"`
}

type Config struct {
	CurrentDB     string `env:"CURRENT_DB" envDefault:"postgres"`
	PostgresDBURL string `env:"POSTGRES_DB_URL" envDefault:"postgres://postgres:user@localhost:5436/user?sslmode=disable"`
	MongoDBURL    string `env:"MONGO_DB_URL"`
	JwtKey        []byte `env:"JWT-KEY" `
}
