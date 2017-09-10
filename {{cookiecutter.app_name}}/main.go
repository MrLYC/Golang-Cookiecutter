package main

import (
    "fmt"
    "{{cookiecutter.app_name}}/config"
)

func main()  {
    fmt.Printf("Mode: %v, Version: %v\n", config.Mode, config.Version)
}
