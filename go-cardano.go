package main

import (
    "devhulk/wallet"
    "fmt"
)

// type Wallet {}

// type Transaction {}

func main() {

    w := wallet.CreateWallet("GeraldWallet")
    fmt.Println(w.Name)
    fmt.Println(w.PaymentAddr)
    fmt.Println(w.Location)
}
