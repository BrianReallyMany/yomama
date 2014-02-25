package jokes

import (
    "strings"
    "io/ioutil"
    "math/rand"
    "fmt"
    "time"
)

func RandomJoke() string {
    rand.Seed(time.Now().UnixNano())
    // Read in jokes file
    text, err := ioutil.ReadFile("jokes/jokes.data")
    if err != nil {
        fmt.Println(err)
        
        return "I/O error --> no joke for you :("
    }
    
    // Choose a joke randomly and return it
    jokes := strings.Split(string(text), "\n")
    jokenumber := rand.Intn(len(jokes))
    return jokes[jokenumber]+"\n"
}

