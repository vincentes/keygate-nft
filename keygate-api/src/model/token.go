package model

import (
	"context"
	"database/sql"
	"keygate/api/network"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/program/token"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/labstack/gommon/log"
)

type Token struct {
	ID string `json:"id"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Address string `json:"address"`
}

func NewToken (name string, description string) Token {
	return Token{
		Name: name,
		Description: description,
	}
}

func CreateToken (tx *sql.Tx, tkn *Token) error {
	context := context.Background();
	
	// Token account
	c := network.GetSolanaClient();

	account := types.NewAccount();

	log.Printf("Creating token account: %s", account.PublicKey);

	tkn.Address = account.PublicKey.String();

	rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(context, token.MintAccountSize);

	if err != nil {
		log.Fatalf("Failed to get minimum balance for rent exemption: %v", err);
		return err;
	}

	// log the rent exemption balance
	log.Printf("Rent exemption balance: %d", rentExemptionBalance);

	res, err := c.GetLatestBlockhash(context)
	if err != nil {
		log.Fatalf("Failed to get latest blockhash: %v", err);
		return err;
	}

	// log the latest blockhash
	log.Printf("Latest blockhash: %s", res.Blockhash);

	systemAccount := network.GetSystemAccount();

	solTx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        systemAccount.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				system.CreateAccount(system.CreateAccountParam{
					From:     systemAccount.PublicKey,
					New:      account.PublicKey,
					Owner:    common.TokenProgramID,
					Lamports: rentExemptionBalance,
					Space:    token.MintAccountSize,
				}),
				token.InitializeMint(token.InitializeMintParam{
					Decimals:   8,
					Mint:       account.PublicKey,
					MintAuth:   systemAccount.PublicKey,
					FreezeAuth: nil,
				}),
			},
		}),
		Signers: []types.Account{systemAccount, account},
	})

	if err != nil {
		log.Fatalf("Failed to create transaction: %v", err);
		return err;
	}

	txSig, err := c.SendTransaction(context, solTx);
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err);
		return err;
	}

	log.Printf("Transaction sent: %s", txSig);

	tkn.Address = account.PublicKey.String();

	_, err = tx.Exec("INSERT INTO Tokens (ID, Name, Description, Address) VALUES (?, ?, ?, ?)", tkn.ID, tkn.Name, tkn.Description, tkn.Address);

	return nil
}

func GetTokens (tx *sql.Tx) ([]Token, error) {
	rows, err := tx.Query("SELECT ID, Name, Description, Address FROM Tokens");
	if err != nil {
		return nil, err;
	}
	defer rows.Close();

	tokens := []Token{};
	for rows.Next() {
		var tkn Token;
		err = rows.Scan(&tkn.ID, &tkn.Name, &tkn.Description, &tkn.Address);
		if err != nil {
			return nil, err;
		}
		tokens = append(tokens, tkn);
	}
	return tokens, nil;
}