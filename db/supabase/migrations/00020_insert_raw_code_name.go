package migrations

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"os/exec"
)

func init() {
	goose.AddMigrationContext(upInsertRawCodeNames, downInsertRawCodeNames)
}

func upInsertRawCodeNames(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	dbString := os.Getenv("SB_DB_LOCAL")

	filePath := "data/input/jscode20240208/KIKcd_H.20240208.csv"

	log.Printf("Inserting records from file: %s\n", filePath)
	psqlCmd := fmt.Sprintf(`\copy raw.code_names FROM '%s' WITH (FORMAT csv, DELIMITER ',', HEADER true, ENCODING 'UTF8');`, filePath)
	cmd := exec.Command("psql", "-v", "ON_ERROR_STOP=1", dbString, "-c", psqlCmd)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("file: %s\nfailed to execute psql command: %w\ngot stderr: %s", filePath, err, stderr.String())
	}
	log.Printf("psql command executed successfully for file: %s\nstdout: %s\n", filePath, out.String())

	log.Println("Successfully inserted records into raw.code_names")
	return nil
}

func downInsertRawCodeNames(ctx context.Context, tx *sql.Tx) error {
	//goland:noinspection SqlWithoutWhere
	_, err := tx.Exec("DELETE FROM raw.code_names")
	if err != nil {
		return fmt.Errorf("failed to delete records from raw.code_names: %w", err)
	}
	return nil
}
