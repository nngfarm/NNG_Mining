package mining

import (
	"bufio"
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ssh/terminal"
	"nng_ming/utils"
	"os"
	"strings"
	"syscall"
)

const (
	NNG_VERSION = "1.0.0"
)

type KeystoreFile struct {
	Version    string `json:"version"`
	Account    string `json:"account"`
	Mnemonic   string `json:"mnemonic"`
	PrivateKey string `json:"privateKey"`
}

// AccountFromKeystore get the account
func AccountFromKeystore() (*ecdsa.PrivateKey, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Password Err: ", r)
		}
	}()
	// define privateKey
	var privateKey *ecdsa.PrivateKey
	// read private key from .keystore
	fileContent, err := os.ReadFile(".keystore")
	// no .keystore file found
	if err != nil || len(fileContent) == 0 {
		var (
			nextStep   = 0
			mnemonic   string
			privateHex string
		)
		// Print generate a
		fmt.Println("1 ==> Generate a new one.")
		fmt.Println("2 ==> Restore via private key.")
		fmt.Println("3 ==> Restore via mnemonic phrase.")
		fmt.Printf("The .keystore file was not found. Whether to generate a new mnemonic phrase [1,2,3]: ")
		fmt.Scan(&nextStep)
		// get mnemonic
		switch nextStep {
		case 1:
			mnemonic, err = generateMnemonic()
			if err != nil {
				fmt.Printf("Generating mnemonic failed: %s \n", err.Error())
				return nil, err
			}
		case 2:
			fmt.Printf("Please enter the private key: ")
			fmt.Scan(&privateHex)
			// has0x
			if strings.HasPrefix(privateHex, "0x") {
				privateHex = privateHex[2:]
			}
			// get private key
			privateKey, _ = crypto.HexToECDSA(privateHex)
		case 3:
			fmt.Printf("Please enter the mnemonic phrase: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				mnemonic = scanner.Text()
			}
			// Invalid mnemonic
			if len(strings.Split(mnemonic, " ")) != 12 {
				fmt.Printf("Mnemonic phrase should be 12 characters.")
				return nil, err
			}
		default:
			return nil, errors.New("Invalid input")
		}
		// read password
		password := readConsolePassword("Please set a password for the Mnemonic: ")
		// mnemonic to privateKey
		if privateKey == nil {
			// from mnemonic
			privateKey, err = mnemonicToPrivateKey(mnemonic)
			if err != nil {
				fmt.Println("Mnemonic to Seed Err: ", err.Error())
				return nil, err
			}
			// privateKey
			privateHex = hex.EncodeToString(crypto.FromECDSA(privateKey))
		}
		// to Address
		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		// console mnemonic, address
		fmt.Println()
		fmt.Println("==========================================================")
		fmt.Println("==========================================================")
		fmt.Println("VERSION:", NNG_VERSION)
		fmt.Println("address:", address)
		// has mnemonic
		if mnemonic != "" {
			fmt.Println("mnemonic:", mnemonic)
		}
		fmt.Println("privateKey:", "0x" + privateHex)
		fmt.Println("==========================================================")
		fmt.Println("==========================================================")
		fmt.Println("Important:", "Remember, anyone who obtains your mnemonic phrase can access your funds. ")
		fmt.Println("Please take all necessary protective measures.")
		// write to .keystore
		enPrivateKey := utils.AesEncrypt(privateHex, string(password))
		// set keystore content
		var k KeystoreFile
		k.Version = NNG_VERSION
		k.Account = address.Hex()
		k.PrivateKey = enPrivateKey
		// no mnemonic
		if mnemonic != "" {
			enMnemonic := utils.AesEncrypt(mnemonic, string(password))
			k.Mnemonic = enMnemonic
		}
		// 序列化数据为 JSON
		jsonData, err := json.MarshalIndent(k, "", "    ")
		if err != nil {
			fmt.Printf("Error occurred during marshaling. Error: %s", err.Error())
			return nil, err
		}
		// write to keystore
		os.WriteFile(".keystore", jsonData, 0644)
		//
		return privateKey, nil
	}
	// trim space
	fileContent = bytes.TrimSpace(fileContent)
	if len(fileContent) == 0 {
		return nil, fmt.Errorf("No keystore found")
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
		privateHex := utils.AesDecrypt(keystore.PrivateKey, string(password))
		// to private key
		privateKey, _ = crypto.HexToECDSA(privateHex)
		// pubkey to address
		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		// password failed
		if address.Hex() != keystore.Account {
			fmt.Printf("Password Err, Please Try Again")
			continue
		}
		return privateKey, nil
	}
	// To ECDSA
	return nil, errors.New("Keystore not found")
}

func readConsolePassword(title string) []byte {
	// read password
	for {
		fmt.Printf(title)
		// read Password
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		// password length must > 7
		if err != nil || len(password) < 8 {
			fmt.Println("Password should at least 8 characters")
			continue
		}
		return password
	}
}

func generateMnemonic() (mnemonic string, err error) {
	// Generate a random 128-bit entropy
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

func mnemonicToPrivateKey(mnemonic string) (*ecdsa.PrivateKey, error) {
	// get hd wallet
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	// derivation path
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0") // BIP44: Ethereum
	// Derive new account
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, err
	}
	// Get the account private Key
	return wallet.PrivateKey(account)
}
