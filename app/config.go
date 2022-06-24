package app

import (
    "log"

    "github.com/joeshaw/envdecode"

    "github.com/michaelact/firstApi/model/config"
)

func NewConfig() *config.Global {
    var c config.Global
    if err := envdecode.StrictDecode(&c); err != nil {
        log.Fatalf("Failed to decode: %s", err)
    }

    return &c
}
