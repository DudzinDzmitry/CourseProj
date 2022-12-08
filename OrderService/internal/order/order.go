package order

// Item : struct for medicine
type Item struct {
	Id    string `bson,json:"id"`
	Name  string `bson,json:"name"`
	Count int32  `bson,json:"count"`
	Price int32  `bson,json:"price"`
}

type Config struct {
	CurrentDB     string `env:"CURRENT_DB" envDefault:"postgres"`
	PostgresDBURL string `env:"POSTGRES_DB_URL" envDefault:"postgresql://postgres:catalog@localhost:5436/catalog?sslmode=disable"`
	JwtKey        []byte `env:"JWT-KEY" `
}
