package migrations

import (
	"context"
	"fmt"

	"github.com/bloxapp/ssv/storage/basedb"
	"go.uber.org/zap"
)

var migration_4_exporter_delete_hashed_key = Migration{
	Name: "migration_4_exporter_delete_hashed_key",
	Run: func(ctx context.Context, logger *zap.Logger, opt Options, key []byte, completed CompletedFunc) error {
		return opt.Db.Update(func(txn basedb.Txn) error {
			err := txn.Delete([]byte("operator/"), []byte("hashed-private-key"))
			if err != nil {
				return fmt.Errorf("failed to delete hashed private key: %w", err)
			}
			return completed(txn)
		})
	},
}
