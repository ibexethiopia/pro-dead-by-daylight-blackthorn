package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.8" )

func 5Z1cFgI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MrTgZSnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0xS7EBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7E488mCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEHYCbcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJnax5c3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCdAwG1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cd2gUNGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOCq8vprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QwdykYI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FolrZO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0lMuc8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFHwvvwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHN8FSZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ti1IM1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BqhkIWwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MpWXCHMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRN4G0EKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TM7xylovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIOrbamaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pz3yQxxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSfXzo50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HaFwQkTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1o9er3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJxte9plWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yBKSQZEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rGcrWylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbR02WuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Le4A3PbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnVjOcciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CucgFwkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bEqmDyKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEExMLOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQtZjIilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1JZ2FYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0iP3gNImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qVkw6SjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MboDEXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Etf8Mj1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUQE9VuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vcISZg9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iohmaLq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0kaoMPxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8JTjmHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRQxJ8wBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shL1e3wVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dc7X9L7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3Ski0bBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZhYZiGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAoxoeR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHhHKCHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKVoOQLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DGV49lQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqiANwmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmIwp7brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cBPkCymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6ZYPBRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNNNQ2dqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oz6insZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zXX42w0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KmdUXxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcJ8UZiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UBZzcsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLBrfHQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyVBk3SkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ui9BQA9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYKUWwOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTVtJo0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLSUkoCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mA5GuojOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func anTWTUx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qox9XkYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONjFsOt5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Q4wBVuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sF6k7Tu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckltoaN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uczVSdNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aRWjX9qlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N5g6n33oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UxJADdPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGi877WKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7MDkNKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qKhTeqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojpQscnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOh0WUYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SAufpDE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iA1Q7oXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jb0zwXY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIBVsqlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwWJbqjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgWWz95iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RkYNY0LsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ES9duGr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTqBpPeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8etSnjjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IuWNsnUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fmhe43RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdBTvqpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2a4pWSbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qm7qAkKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rthoPglhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzgpqjwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ib2qbtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5F7wKwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8NR7TuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nh6XBjwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BI5ZoigEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHQLZ3TKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A10nJ6XBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7TAy8E9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func olZnjicFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwyE8WO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0137UBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEtLLkV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgiCq44EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEk8SETsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdgv0nUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtRJ8I7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fXMkLz2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNbVkSVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HkthyBuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8Yvd0QqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMndAWZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXYbYnlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wc3ypBAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAdXZDEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ukt5rxhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDNpmsYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jrdPducWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Kl38iCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkJUELUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5ehAAGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIUU28iAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FeUe19eAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIlBYE57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suwAX8oyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIrRh3yNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnvPgi4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyeiBGTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIZ9nsQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzPrVmaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYuNqGEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQtfdEq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P54GttJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Guq4rx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boEN2W5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8sZ8CIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lGFCj1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8c5ZkUh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nh2cCquPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpdzi4gBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YX59G7X3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87E7UvW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThSoPV3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKauoGvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dv7Z3q6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWPeIRz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aRLxJ5aOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVCXPf8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 58djJyZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64gLop6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OgFSTcfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoFCoFHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHVft79TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qcri1E6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7h7wtiMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LmbYyHKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xM6vzXeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQr4lLyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGRe2t65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLRoHOjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35Dwk7QuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VuZD4a9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebnXcKRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpyUSrvrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03ScnBerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndzBn9lIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vl8LdAW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYrZwUZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdGlaczsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZZD83ZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0y9NC9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpxXmMxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GoY8ur1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6j6XXFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEuNEykPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNDbcab1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYZlIhyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTVyV7j8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqESQPgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7POSJKUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0p7sarP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KD0nECEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWYkEQ23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mH8NrSdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TllbmLPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Pl5rliVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iBUGm7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDpzXj03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZ2wb60pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdwBT6fKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NM583R3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbPU25gwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qcdGDgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sSnSW5fQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAqLIDxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pAdt3GDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48NzfLFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOaVOzo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwrsxqfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hg4UK4o9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Lan9BV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpDl8xYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yWTI73q3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTKsMuj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XY2GtcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Zb7ggerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2H3MEzCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjyndOhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vKmHjQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func do6hXNP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKRxvPSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMtYJayOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dD1lh7xUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99EEwmidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ERJwmacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnjE38fNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9naUyBphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAJHGXEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzR00DjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qyU2URHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFEIUBcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B3WjepyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z45ru9v7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Q3bWW4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M0BE05jyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHqugKunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXv3LGIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1RrNDhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a32j54z5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zLxCsaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9Erp4CHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdznVc6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLxw5n0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZKTL3CRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjhTIlnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21xmHGJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wiga4mfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ht75hMMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AU2sx1XdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1Kdb6DcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGc70Qc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AQgZBPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evr0QMJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGk3m6SlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHe74lr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mMfHI8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWYraHEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBlq8QIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lndyKtMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3VwmdIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VNypFoS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjKno0T7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6cA5jUmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3X7ru0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9N2LM1NoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IGvhkcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmm4ALLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 230CTzZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqjS2thNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ckls5yMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yofhbTmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNmsN3UgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSQBUnN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtiMeNMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kb6tGALrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1psG0qRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D9rrNkFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmrsdcSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozA1Chu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6p7AuDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMtdnd6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func za1QHB7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xpe31OpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzPQrnEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nC8BO77oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HypksnXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1zIaMt2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gPuzDolWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Es9Mr3pzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERF4tjG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtMPm2slWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZ8135zlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6tIOTOjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3hBXW71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uLKPDnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvJJm9voWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1K157yGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgA9P07BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4mAJAdRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4qUeGUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmXG4ncmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7QjDpxcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXxLluPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qiCVtOoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMeAa8M6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y85vdpQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XNq3q4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8HpJNdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YN2xGzDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkgUgG1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RmPYx6myWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzNoYKl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6QGqopwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBvVvxEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64qFoDz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ME330fOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sd6kwPPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MpbumTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNBWdoGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6hfdNikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OcEHKWZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WaVjobD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RM9fET6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyBPO8RLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNNxSLoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIpd49jvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTg2WjsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G22t6c2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjwXU4UWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEE5JFpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1MEFdiCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxhapVLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyPCFBjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agxhQcXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHCAjk5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3tmvZzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgmrkKdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GuGg1lRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llITt9W6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bg5Oi2VtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTr6f1JMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9rkqs4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wa0FARwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dl6FoMalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkr5Ju2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uN095c6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NbzydcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHqWHtgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTKPyOyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwXLWEr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yH948R48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3IDSpArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bbt7EyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0NrM3kHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8p9Cp0NNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1fM0pNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCE9w5UFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZF9zYu2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RaN1dwM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpUkRSlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bD1LcX3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5IuiRx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJ0fpwW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLQiYUICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WZCKcKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yi1njqLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKNdIJtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVRyi1oLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlkdirVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzOQogjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1l6BNrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEgq53g4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57T4dGPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhAPjgs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y76RMcNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRcM3nnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B71NW4JQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zR1Up7EqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3QjEBwOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bL3w13mTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTaxBQuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R45UyP47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aCOcIpHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8iximfHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCSrEna2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MzOof5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihgmC2CNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQlgLKTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkuIlp2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FcBciaXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bgOJcUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hT53nMDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OImuIVp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGnddR7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0Vcl1N2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iw40mhbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2chs8e1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbYEBBq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGLRW3aRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5zsNdSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INNlgnjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aio1VJtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYNsOK0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZxl2xbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CpjPW8NFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uYPnhu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52KXcK1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQR7gHf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJcwTInrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ixlq6bqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJtDhRp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrVWHTaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dvr4jJsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqpiyNLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AEdXCNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EotCtPavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WZCQM7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GA4tojQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05j0m3L4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0aZleMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvYjhHcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hygQdSkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlGZFRrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2itkuqsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 432aIeKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoIKRXMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhhizqMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8siiNU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVKu685CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WbIZrVUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXIlBysnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FBYt14LRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhHrKwHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6OhMUYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7aXNnj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yv2bmH7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNjV0cE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARVw3jIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHgZThshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xx6ebJetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9tn8BxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DceA5jleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXX7rvWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ce5I0T0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Us01lRDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yZBjy4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ll2r3LXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZiGs5aLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IQxWe0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LegCNza6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSVBs01oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7jWFM0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdCoqcMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bESOmORlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0BUmTQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YV57GiQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7M9QcDFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0qEUYAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1hxLX7qqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ca4OI8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YcHCXWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sd81p4I7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhedqboHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ELGLyhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nxtyCU9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khwk5i0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFIFtWvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pB2lS7XJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIvJJ2OWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33zyNDarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGXYVIWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXEGpalVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMFZndNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6E3Y9EjBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3iSpJJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5k6d6zXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVBBx08eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OE4TeiTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bb0rc8TjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3KXMkZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWmsZRTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21swBce3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XBjnbvucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdxNqzSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YfuYU7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pJjbQxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vl9yWjZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PkZqMxM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLVZqXRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1CthxOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgigPUl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTQjCGKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxvyMhwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xk8H7D0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrYep2uZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNGk6DR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdsZ0azPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzPeIZe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3mmTUqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QM6iBAKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04qlouYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfev3LqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSigV9PZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0er40dSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xSuCUICTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICdt8uV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fqoqdKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqTwHy7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dBNZqpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OI0rtNsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lG2TXqmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tu4C49gOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzdaoeqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4B80vwP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P43gWlSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtkRA0QJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3nWQgBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpEQCzycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DADPq7j4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vdeGFwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1V6TILfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTCg8KJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUkethIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfEiEjQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOlq0QU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hRXnpwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQRr7Gr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 11tSB6ibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h4tZOhiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxIGVNzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZdR5CfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8HtqTr6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZaCQ68MSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gx3t4LsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WIf3MFeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1CMRk2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8a9C1DLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZYZIKQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjV77qeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQygDPokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1DreGgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAjlu8RCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6d2JFQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7eRpcAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyXgVhrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcf3nmAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxoCLU2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRTSWTgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AAT73dA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmkW8MCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3lq0Zq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PX49P1tEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nc47efuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5F3NgucXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func As0RYY0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6LGaH2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jYPnWuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XH8Goy6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWATEo7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Up94LJ1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlPnrd5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNQnug9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeFXP4HDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func brCIHmlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocJYHATvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIfw9VMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVlQVLDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSgVVQVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5xWAazKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y1QGqNR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xutdyDxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gE6WhOOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yE15Xdz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfnMD5KDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eW0VvfilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3WlN9sIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udFkXkjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EvUum3hQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RX3nVGY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0KQm8dmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXHh9hnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qp3aDhnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AqVeCLtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5cjfhT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22jUQYthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WeDadLnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5uo95NgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6HlIFNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7EfacCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPHXuySuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AP2d8G57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJusvAg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJP9MeKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETZXrztNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7h8KFQiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7zX8d0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogFs6WHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1j5TdsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0P53IRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ub8R4mM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96aB8OjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76z3oGQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWRbXUOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9urCukGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XhUmFf38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTnZf5swWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ku0Hs4qPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0kooLOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hSji9KzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Yeop9q8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3hWOxAo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwzpBiukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlC3QAr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTwpfHqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWZcvf4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUfr3DLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKdA4hiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSG9gNT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDVzE5qoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVRCYYgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbdZ5VtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKjTbmXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KfIi7rE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6DqWtDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSuIwFhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URGNeoWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54NfhFOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7bF1puYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrSOSUQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQSv5tlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmbRyXQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMOq76KqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obpQLRKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiWJAouyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvasFbbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Pq1P3XPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3qwpvTJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjNhJHmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haQtCacpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spQUPKj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYZ27b6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jITLe5fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYGSUwkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ayl8sVpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ud7uWt75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfdTPOAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCLL5ildWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEuAANnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qoV5m4eEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KO7cR1UyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEZ7BhN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4X8BjWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGjs88GGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vp7msTNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMMakGHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImJlNriJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1MDWjdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zelA0oUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lty8CzkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDmBDMFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHvosRjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oT2cEAF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3XNPlpssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hVUCKDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vtZ5BlvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AFbpDTrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dZg1GWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bolDJ2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ci3fKkypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tDm1euhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6XO2x5HGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giW3orGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WA5pKmmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ngz78O9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6g6yU5uUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKXhFaoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dw46G8LRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Yb2h6k4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mOAf9DjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XScoZYqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1CNC8J0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srVD3GNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V66lzX71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hydDGrcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DaUSUobZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjS4juHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJ3iiU00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55qiNf7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tKhDq0lZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KddDZgYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oEZF3GvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L76GeEylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTStuossWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrbXVES8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwZQQIiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRNIkF4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYWmE26zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seaaaIiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QV3hLCo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZHtJ9ASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mBsbWFiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func roKYoVOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UlvncIaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTO3gdIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVosILpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GfDm6NPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOiqbVNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdPgW1ZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LAKoDCvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cz8tqzXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gd9FXXAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 00p0Od3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxzIduDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52gpwRAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Npg9PvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8JLU6VR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cjovLIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YstccrgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGNSDt0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIs49mO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yuiflWt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSKXrYuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qauGAiDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiNn7AVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvuJhPEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LD6zBTrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMAj2QrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dm5mR0s5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qgs3QqnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2rBADefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxltKUPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKqomkDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XSmeOn7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yp2FP67kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQyQHQrvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cyEqCVneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AsEBDjETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRC8uy2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkUCzXSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rwQe2uGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7e5CNWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Rzle22dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTSvl7cbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lr8aDSkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2fb6C9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RE0TGrH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAAvDz7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1KX6h5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2wX9uC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bjans9V6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPDn3oGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T672uMb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYdwS41FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hr6aVQiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKoYBA3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwStN8h0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 11VN4lChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DwoUxI6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0969WnHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcY2YFAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MjvWo8jtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpaE0IdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkPhWsA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3BpN41CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4Yg2BKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KU2rDiqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYCwSMQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Buuum4wcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSrC2ie2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmNP8C8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SpIP9K8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmF0K6v5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJWEKhS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CgCMcZIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXUew2rRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pj5emMHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6T18FCdAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uYJcEQQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfastl89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mj05OsP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ondE5yYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJhl9nzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSymPBnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjgJ5AjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVenziOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsRnpsWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RIm84IFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jiM0LtMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UhmzLTL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J68dxxiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQl79Dz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lyqg0loLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLLVdCkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGaps8yqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1VfToeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mk5f8mk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1PlFmAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5COkxIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DySXCgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXDYASEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqmugBEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Irl0Z6BAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I41Sk7Q0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5qdPNPEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kycEghLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLzYxsAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hN5Umo2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hgo967w8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQJOWQFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gE6XkTASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zO2UUcaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cV8UKqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACwqzSP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGAo445FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofllXGnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hoLoCZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4fGkng6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqNWfwxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12VIcIJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXZ6SV9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fst7hWmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zu1R4bElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Ul6nq5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqL6NzDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4IAEUyAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZOJFwrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46O0Rm4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmE7xy1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8N0slQdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIdsGOz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8F5o1WVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AyLt0ljCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZ3Kai5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 613pDxXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s3KrRX6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twz42HhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOcVRicJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2OegNPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohkF5CnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5aEFGxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvmSyktQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIIUFIemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func byoybHCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gq4VLSXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Yv03ufzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R97z58lIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pi4qhU3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3q0XaFUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QM8XF6EvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZ4BPRkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cm089QgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UM9ZzWzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNMQ37hnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47ak7PHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oTTQDZmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXxVjpZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONDVK5SoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3gyprEeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDDVT6KZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hmxV1KZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oB99FRpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SVu5bQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWLYWhw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wrq80MwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQA12RqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDblXeMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebcNdg9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0kT9T12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EyCJ1nwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2jWlwr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZDaIoQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func caEzNLLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGsT385MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 163KZfSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4mGrjGaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZV86Bs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flp1kIqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DevSh6IjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4rrciuaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwFZlfYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIVplmdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qycaStAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3vF86PoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSw2PkojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHcBHhmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QvmiO4eqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xCEXPDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oX52Q3ZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IouYsKduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1KxoHoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZ3OiqqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ND7LYXhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDdtsARaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaLDUndCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyUrLbrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQOVdljMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUlrHO8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpu0Jj9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8EOYXMPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yeZZS2tjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 49WkZXYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCvC1XANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHJCYbliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlCyhfO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3I5Ie82dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Hix4i2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oop6t8IUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdvNooXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyBi94LnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qf3eQkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cahpgv83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ap7Lhy6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsIeJy9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ua4fqrK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmtdpMK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2H23jMnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUd8UZ8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZbkDiFvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRjuA5H6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 49ZrmBKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNzb7Pc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyzTEiIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0DFpRHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbG5cbVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVR1DqTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func laQWbPVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66NhAURPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OVLNms4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4RgipP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlwccAQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3tUojvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOxzt4IEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgFIsbKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wzo0ML9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 02HfJYPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0vF3i60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xov4oFNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcVkIKOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNls1cu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yEQuiZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVhycXS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HL1GZYKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vl88wdOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXW1nj2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAPnFuN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dfi577MEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0nkHI8QbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gm9tWtY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uy4FQBX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMbGqeWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qien6da2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsEJNJgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func guoTR9jMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C49NzzEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kih3XdabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqmX939TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4bjq7SXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gkWhtB6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYBbi0OfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRUHhHcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPxjfPa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtLsy3NyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APcSATv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzBWUKxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CJBdvjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VF7weBqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78bEoGGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7Lf37H3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ad3S9oEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9eYZoIKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agqBvzbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQyLxd5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckLprKMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BEheRRjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odZouVWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xD5W7JxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mb8Z7qX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RFD83FZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVUFMQNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFXFqwxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiLfVtTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E68yYxEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GibUQKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAg9ge3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GC9XA3ecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5xLPvJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYqMu0vbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6d8cP2zEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbC5O5ijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zgoriBYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1cYot5tOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vh8vVZPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6M1h6bK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7mtUNGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhIYoXGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U19LoQTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rozMddS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOu5txm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GydzHl1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEv56J8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYItqw79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55VR4LYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBKOp8juWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npw23XHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuncgHynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func foPGvZS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 75qqhldRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4ZzxklQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2q2OJ5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AuMmEpEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UGEkPX17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7wVBUYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3usjFNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y459RvwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJLLuHOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Js9sEqVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGGi5VQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ScRfeNGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HOzsbssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIRzIDg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InOwmtacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2Pw693aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUCcIMjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0KRrRsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ua82CN31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6ONE7omWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKcixMEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmctuwGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3tVSjeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRWduW5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPD5YvODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwKLKLd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZeCZBnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AddavEFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eX3DekupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EIdRomr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWYpnNq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxHug6P0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HlcFJk3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAC7p5tjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qcT61WpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cehh9467Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTd4X9QHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBJCbRNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAlwYrWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzU1SCAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Timkma58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5KGR1r85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWvbPDuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AR9u5nyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1eSv7a2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yl817xzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rx23OwCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1bKXkFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqSYO4BwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func naZMBffjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQP42t8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Irf1lyxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKBi9PiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GpgdAr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZ8TeoHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2O8OcvAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jl4C5yCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0JsPTndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JjBc43U4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cl154zZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AADdxkIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GiuC2nEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Awnb2I7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNWnvXqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MS7wOXOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yFglCraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y1rM1tqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ibCl6UIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xY2t5mccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8FGaR9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PjTJ3HtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNm56pbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBFokXwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cYgXpTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AccM8j3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G16ah0eoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TsP8rlTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NajuRqB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7Nhy3b4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8IILPmuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PQ2mPioPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hEo4wWbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UvxRaGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSRB5oQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kghntmTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBpPcyqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8xA1EU78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15EeS5N4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wy0HIztOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SB4g8WzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

