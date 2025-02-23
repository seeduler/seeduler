package utils

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

func DecodeRequestBody(r *http.Request, v interface{}) error {
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(v)
    if err != nil {
        return err
    }
    return nil
}

func ParseQueryParams(r *http.Request, key string) ([]int, error) {
    queryValues := r.URL.Query().Get(key)
    if queryValues == "" {
        return nil, nil
    }

    values := strings.Split(queryValues, ",")
    intValues := make([]int, len(values))
    for i, v := range values {
        intValue, err := strconv.Atoi(v)
        if err != nil {
            return nil, err
        }
        intValues[i] = intValue
    }
    return intValues, nil
}