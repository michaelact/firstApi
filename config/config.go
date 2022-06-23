package config

import (
    "log"

    "github.com/joeshaw/envdecode"

    "github.com/michaelact/firstApi/model/environment"
)

func NewConfig() *environment.Global {
    var c environment.Global
    if err := envdecode.StrictDecode(&c); err != nil {
        log.Fatalf("Failed to decode: %s", err)
    }

    return &c
}
