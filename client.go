package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

func GenTest(str string) *CertificateResponse {
    client := &http.Client{Timeout: 10 * time.Second}

    resp, err := client.Get("http://127.0.0.1:9000/ApplyForABSCertificate?uid=" + str)
    if err != nil {
        return nil
    }
    defer resp.Body.Close()

    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil
    }

    var cer *CertificateResponse
    if err := json.Unmarshal(content, &cer); err != nil {
        return nil
    }

    return cer
}

func VerifyTest(cer *CertificateResponse) {
    body, err := json.Marshal(cer)
    if err != nil {
        return
    }

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Post("http://127.0.0.1:9000/VerifyABSCertificate", "application/json", bytes.NewReader(body))
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer resp.Body.Close()


    //content, err := ioutil.ReadAll(resp.Body)
    //if err != nil {
    //    return
    //}

    //fmt.Println(string(content))
}

func abs_test() {
    fmt.Println("ABS test ---------------------")
    fmt.Print("ABS gen: ")
    start := time.Now().UnixNano()
    for i := 0; i < 100; i += 1 {
        GenTest("123")
    }
    end := time.Now().UnixNano()
    fmt.Println(float64(end - start) / 1e9)

    fmt.Print("ABS gen & verify: ")
    start = time.Now().UnixNano()
    for i := 0; i < 100; i += 1 {
        sign := GenTest("123")
        VerifyTest(sign)
    }
    end = time.Now().UnixNano()
    fmt.Println(float64(end - start) / 1e9)
}

func main() {
    c := GenTest("123")
    fmt.Println(c)

    for true {
        VerifyTest(c)
        time.Sleep(time.Duration(5)*time.Second)
    }
    //abs_test()
    //rsa_test()
    //ecdsa_test()
}
