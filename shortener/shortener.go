package shortener

import (
    "strconv"
    "math/rand/v2"
    "errors"
)

const (
    maximumLength = 13
    minimumLength = 7
    numericProb = 0.3
    alphabetLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
    errNotFound = errors.New("Given token not present in the database")
)

var db = make(map[string]string, 0)

func Get(token string) (string, error) {
    url, ok := db[token]
    if ok {
        return url, nil
    }

    return "", errNotFound
}

func Add(url string) string {
    for t, u := range db {
        if u == url {
            return t
        }
    }

    shorten := token()
    db[shorten] = url 
    return shorten
}

func randomDigit() string {
    offset := rand.IntN(9)
    return strconv.Itoa(offset)
}

func randomChar() string {
    offset := rand.IntN(len(alphabetLetters) - 1)
    return string(alphabetLetters[offset])
}

func token() string {
    size := rand.IntN(maximumLength - minimumLength) + minimumLength
    res := ""
    for i := 0; i < size; i++ {
        if rand.Float64() < numericProb {
            res += randomDigit()
        } else {
            res += randomChar()
        }
    }
    return res
}
