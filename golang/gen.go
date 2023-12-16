package main

import (
	"encoding/base64"
	"encoding/csv"
	"log"
	"os"

	"github.com/blocto/solana-go-sdk/types"
)

const WALLETSAMOUNT int = 20

func main() {

	csvFile, err := os.Create("wallets.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	header := []string{"Wallet", "Private Key"}
	csvWriter.Write(header)

	for i := 0; i <= WALLETSAMOUNT; i++ {
		wallet := types.NewAccount()
		privateKey := base64.StdEncoding.EncodeToString(wallet.PrivateKey)

		data := []string{wallet.PublicKey.ToBase58(), privateKey}
		csvWriter.Write(data)
	}

}
