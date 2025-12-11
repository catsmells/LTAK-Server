package config

type Config struct {
    ListenAddr string
    DBPath     string
    TileRoot   string
    JwtSecret  string
}

func LoadOrDefault() *Config {
    return &Config{
        ListenAddr: ":8080",
        DBPath:     "ltak.db",
        TileRoot:   "assets/tiles",
        JwtSecret:  "enter secret here upon install",
    }
}
