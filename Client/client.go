package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
    "sync"
    "time"
)

func GenTest(uid string) string {
    client := &http.Client{Timeout: 10 * time.Second}

    resp, err := client.Get("http://127.0.0.1:8001/ApplyForABSCertificate?uid=" + uid)
    if err != nil {
        return err.Error()
    }
    defer resp.Body.Close()

    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err.Error()
    }

    cer := string(content)
    return cer
}

func VerifyTest(no string) bool {
    client := &http.Client{Timeout: 10 * time.Second}

    resp, err := client.Get("http://127.0.0.1:8001/VerifyABSCertificate?no=" + no)
    if err != nil {
        return false
    }
    defer resp.Body.Close()

    return true
}

func abs_test(num int) {
    fmt.Println("ABS test ---------------------")
    //fmt.Print("ABS gen: ")
    //start := time.Now().UnixNano()
    //for i := 0; i < 100; i += 1 {
    //    GenTest("123")
    //}
    //end := time.Now().UnixNano()
    //fmt.Println(float64(end - start) / 1e9)

    fmt.Print("ABS gen & verify: ")
    start := time.Now().UnixNano()
    var wg sync.WaitGroup

    for j := 0; j < num/100; j += 1 {
        for i := 0; i < 100; i += 1 {
            wg.Add(1)

            go func(uid string) {
                defer wg.Done()

                sign := GenTest(uid)
                var cer CertificateResponse
                if err := json.Unmarshal([]byte(sign), &cer); err != nil {
                    return
                }
                VerifyTest(cer.CertificateContent.SerialNumber)
            }(strconv.Itoa(i))

        }
        wg.Wait()
    }

    end := time.Now().UnixNano()
    fmt.Println(float64(end - start) / 1e9)
}

func main() {
    //c := GenTest("123")
    //fmt.Println(c)
    //
    //for true {
    //    VerifyTest(c)
    //    time.Sleep(time.Duration(5)*time.Second)
    //}
    num := flag.Int("n", 1000, "number of test.")
    flag.Parse()
    abs_test(*num)
    //rsa_test()
    //ecdsa_test()
}
