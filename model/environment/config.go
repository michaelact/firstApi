package environment

type Global struct {
    Api      Api
    Database Database
    Server   Server
}

type Api struct {
    Key string `env:"API_KEY,required"`
}

type Server struct {
    Host   string `env:"SERVER_HOST,required"`
    Port   string `env:"SERVER_PORT,required"`
}

type Database struct {
    User     string `env:"DATABASE_USER,required"`
    Password string `env:"DATABASE_PASSWORD,required"`
    Host     string `env:"DATABASE_HOST,required"`
    Port     int    `env:"DATABASE_PORT,required"`
    Name     string `env:"DATABASE_NAME,required"`
}
