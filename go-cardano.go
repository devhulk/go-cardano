package main

import (
    "devhulk/wallet"
    "fmt"
)


func main() {

    w := wallet.CreateWallet("GeraldWallet")
    w2 := wallet.CreateWallet("GeraldWallet2")
    fmt.Println(w.Name)
    fmt.Println(w.PaymentAddr)
    fmt.Println(w.Location)

    fmt.Println(w2.Name)
    fmt.Println(w2.PaymentAddr)
    fmt.Println(w2.Location)
}
