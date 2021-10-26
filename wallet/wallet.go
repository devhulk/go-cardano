package wallet

import (
    "os/exec"
    "fmt"
)

type Wallet struct {
    Name string
    Location string
    PaymentAddr string
}

//cardano-cli address key-gen --verification-key-file payment.vkey --signing-key-file payment.skey
func CreateWallet(name string) *Wallet {
    CreateWalletKeys()
    w := CreatePaymentAddress(name)
    return w
}

func CreateWalletKeys(){
    cmd := exec.Command("cardano-cli", "address", "key-gen", "--verification-key-file", "payment.vkey", "--signing-key-file", "payment.skey")

    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Print(string(stdout))
}

// cardano-cli address build --payment-verification-key-file payment.vkey --out-file payment.addr --testnet-magic 1097911063
func CreatePaymentAddress(name string) *Wallet {
    arg1 := "address"
    arg2 := "build"
    arg3 := "--payment-verification-key-file"
    arg4 := "payment.vkey"
    arg5 := "--out-file"
    arg6 := "payment.addr"
    arg7 := "--testnet-magic"
    arg8 := "1097911063"

    cmd := exec.Command("cardano-cli", arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)

    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Print(string(stdout))

    showAddr := exec.Command("cat", "payment.addr")

    addr, err2 := showAddr.Output()

    if err2 != nil {
        fmt.Println(err2.Error())
    }

    // fmt.Println(string(stdout2))

    // printDir := exec.Command("pwd")

    // location, err3 := printDir.Output()

    // if err3 != nil {
    //     fmt.Println(err3.Error())
    // }

    // fmt.Print(string(location))

    // Mkdir with keys and payment addr

    walletDir := fmt.Sprintf("priv/%v", string(name))

    mkdir := exec.Command("mkdir", "-p", walletDir)

    walletDirOutput, err4 := mkdir.Output()

    fmt.Print(string(walletDirOutput))

    if err4 != nil {
        fmt.Println(err4.Error())
    }
    
    mv := exec.Command("/bin/sh", "-c", fmt.Sprintf("%v %v %v", "mv", "./payment.*",fmt.Sprintf("./%v", walletDir)))

    mvOutput, err5 := mv.Output()

    if err4 != nil {
        fmt.Println(err5.Error())
    }

    fmt.Print(string(mvOutput))

    return &Wallet {name, walletDir, string(addr)}


}

