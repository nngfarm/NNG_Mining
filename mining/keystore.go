package mining

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"nng_ming/utils"
	"os"
	"strings"
	"syscall"
)

const (
	NNG_VERSION = "0.1.0"
)

type KeystoreFile struct {
	Version  string `json:"version"`
	Account  string `json:"account"`
	Mnemonic string `json:"mnemonic"`
}

// AccountFromKeystore get the account
func AccountFromKeystore() (*ecdsa.PrivateKey, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Password Err: ", r)
		}
	}()
	// read private key from .keystore
	fileContent, err := os.ReadFile(".keystore")
	if err != nil || len(fileContent) == 0 {
		var (
			isContinue = "n"
		)
		// Print no
		fmt.Printf("The .keystore file was not found. Whether to generate a new mnemonic phrase? y/n : ")
		fmt.Scan(&isContinue)
		// if not Exit.
		if isContinue != "y" && isContinue != "Y" {
			log.Fatal("Exit.")
		}
		// read password
		password := readConsolePassword("Please set a password for the Mnemonic: ")
		// Generate a random 128-bit entropy
		entropy, err := bip39.NewEntropy(128)
		if err != nil {
			log.Fatalf("Error generating entropy: %s", err.Error())
		}
		// Generate a mnemonic from the entropy
		if mnemonic, err := bip39.NewMnemonic(entropy); err != nil {
			log.Fatalf("Error generating mnemonic: %s", err.Error())
		} else {
			// mnemonic to privateKey
			privateKey, err := mnemonicToPrivateKey(mnemonic)
			if err != nil {
				return nil, err
			}
			address := crypto.PubkeyToAddress(privateKey.PublicKey)
			// console mnemonic, address
			fmt.Println()
			fmt.Println("==========================================================")
			fmt.Println("==========================================================")
			fmt.Println("VERSION:", NNG_VERSION)
			fmt.Println("address:", address)
			fmt.Println("mnemonic:", mnemonic)
			fmt.Println("==========================================================")
			fmt.Println("==========================================================")
			fmt.Println("Important:", "Remember, anyone who obtains your mnemonic phrase can access your funds. ")
			fmt.Println("Please take all necessary protective measures.")
			// write to .keystore
			enMnemonic := utils.AesEncrypt(mnemonic, string(password))
			// set keystore content
			var k KeystoreFile
			k.Version = NNG_VERSION
			k.Mnemonic = enMnemonic
			k.Account = address.Hex()
			// 序列化数据为 JSON
			jsonData, err := json.MarshalIndent(k, "", "    ")
			if err != nil {
				log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
			}
			// write to keystore
			os.WriteFile(".keystore", jsonData, 0644)
			//
			return privateKey, nil
		}
	}
	// trim space
	fileContent = bytes.TrimSpace(fileContent)
	if len(fileContent) == 0 {
		log.Fatalf("No keystore found")
	}
	// foreach to input keystore
	for {
		// read password
		password := readConsolePassword("Please input password to read keystore: ")
		fmt.Println()
		// key store
		var keystore KeystoreFile
		json.Unmarshal(fileContent, &keystore)
		// decrypt keystore to private key
		mnemonic := utils.AesDecrypt(keystore.Mnemonic, string(password))
		// Split mnemonic
		if len(strings.Split(mnemonic, " ")) != 12 {
			fmt.Println("Password Err, Please Try Again")
			continue
		}
		// Check Address
		privKey, err := mnemonicToPrivateKey(mnemonic)
		if err != nil {
			fmt.Printf("Error occurred during mnemonic decryption. Error: %s", err.Error())
			continue
		}
		// pubkey to address
		address := crypto.PubkeyToAddress(privKey.PublicKey)
		// password failed
		if address.Hex() != keystore.Account {
			fmt.Printf("Password Err, Please Try Again")
			continue
		}
		return privKey, nil
	}
	// To ECDSA
	return nil, errors.New("Keystore not found")
}

func readConsolePassword(title string) []byte {
	// read password
	for {
		fmt.Printf(title)
		// read Password
		password, err := terminal.ReadPassword(syscall.Stdin)
		// password length must > 7
		if err != nil || len(password) < 8 {
			fmt.Println("Password should at least 8 characters")
			continue
		}
		return password
	}
}

func mnemonicToPrivateKey(mnemonic string) (*ecdsa.PrivateKey, error) {
	// New seed
	seed := bip39.NewSeed(mnemonic, "")
	//
	return crypto.ToECDSA(seed[:32])
}
