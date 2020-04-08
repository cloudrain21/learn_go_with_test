package hello

import "fmt"

const englishhelloPrefix = "Hello, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case "Spanish":
        prefix = "Hola, "
    case "Japan":
        prefix = "Gonnichiwa, "
    default:
        prefix = "Hello, "
    }
    return
}

func main() {
    fmt.Println(Hello("dplee", "xxx"))
}
