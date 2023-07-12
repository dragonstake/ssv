package cli

import (
	"encoding/base64"
	"encoding/json"
	"github.com/bloxapp/eth2-key-manager/encryptor/keystorev4"
	"log"
	"os"

	"github.com/bloxapp/ssv/logging"
	"github.com/bloxapp/ssv/utils/rsaencryption"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// generateOperatorKeysCmd is the command to generate operator private/public keys
var generateOperatorKeysCmd = &cobra.Command{
	Use:   "generate-operator-keys",
	Short: "generates ssv operator keys",
	Run: func(cmd *cobra.Command, args []string) {
		encrypted, _ := cmd.Flags().GetBool("encrypted")
		password, _ := cmd.Flags().GetString("password")

		if (encrypted && password == "") || (!encrypted && password != "") {
			log.Fatal("Both encrypted and password flags must be provided together")
		}

		if err := logging.SetGlobalLogger("debug", "capital", "console", ""); err != nil {
			log.Fatal(err)
		}
		logger := zap.L().Named(RootCmd.Short)

		pk, sk, err := rsaencryption.GenerateKeys()
		if err != nil {
			logger.Fatal("Failed to generate operator keys", zap.Error(err))
		}
		logger.Info("generated public key (base64)", zap.String("pk", base64.StdEncoding.EncodeToString(pk)))

		if encrypted {
			encryptedData, err := keystorev4.New().Encrypt(sk, password)
			if err != nil {
				logger.Fatal("Failed to encrypt private key", zap.Error(err))
			}

			encryptedJSON, err := json.Marshal(encryptedData)
			if err != nil {
				logger.Fatal("Failed to marshal encrypted data to JSON", zap.Error(err))
			}

			err = os.WriteFile("encrypted_private_key.json", encryptedJSON, 0600)
			if err != nil {
				logger.Fatal("Failed to write encrypted private key to file", zap.Error(err))
			}

			logger.Info("private key encrypted and stored in encrypted_private_key.json")
		} else {
			logger.Info("generated public key (base64)", zap.String("pk", base64.StdEncoding.EncodeToString(pk)))
			logger.Info("generated private key (base64)", zap.String("sk", base64.StdEncoding.EncodeToString(sk)))
		}
	},
}

func init() {
	generateOperatorKeysCmd.Flags().BoolP("encrypted", "e", false, "Encrypt the private key")
	generateOperatorKeysCmd.Flags().StringP("password", "p", "", "Password for the encrypted private key")
	RootCmd.AddCommand(generateOperatorKeysCmd)
}
