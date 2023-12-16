package main

import "fmt"

func main() {
    var sec, ore, min, giorni, anni int
    fmt.Scan(&sec)
    min = sec / 60
    sec = sec % 60
    ore = min / 60
    min = min % 60
    giorni = ore / 24
    ore = ore % 24
    anni = giorni / 365
    giorni = giorni % 365
    fmt.Println(anni,"anni",giorni,"gg",ore,":",min,":",sec)
}