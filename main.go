package main
import (
    "fmt"
    "github.com/rapid7/go-get-proxied/proxy"
)
func main() {
    p := proxy.NewProvider("").GetProxy("http", "http://whitehatsecurit.com")
    if p != nil {
        fmt.Printf("Found proxy: %s\n", p)
    }
}
