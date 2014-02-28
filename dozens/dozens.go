package dozens

import (
    "strings"
    "io/ioutil"
    "math/rand"
    "fmt"
    "time"
)

func RandomDozens() string {
    rand.Seed(time.Now().UnixNano())
    // Read in dozens file
    text, err := ioutil.ReadFile("dozens/dozens.data")
    if err != nil {
        fmt.Println(err)
        
        return "I/O error --> no joke for you :("
    }
    
    // Choose a joke randomly and return it
    dozens := strings.Split(string(text), "\n")
    dozensnumber := rand.Intn(len(dozens)-1)
    return dozens[dozensnumber]+"\n"
}

