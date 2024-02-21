package migrations

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"log"
	"os"
	"os/exec"
)

func init() {
	goose.AddMigrationContext(upInsertRawLocationsSummary, downInsertRawLocationsSummary)
}

func upInsertRawLocationsSummary(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	dbString := os.Getenv("SB_DB_LOCAL")

	dirPath := "data/input/location_202401"

	filePaths, err := fileprocessor.GetFilesWithExt(dirPath, ".txt")
	if err != nil {
		return err
	}

	for _, filePath := range filePaths {
		log.Printf("Inserting records from file: %s\n", filePath)
		psqlCmd := fmt.Sprintf(`\copy raw.locations_summary FROM '%s' WITH (FORMAT csv, DELIMITER '|', HEADER false, ENCODING 'UHC');`, filePath)
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
	}

	log.Println("Successfully inserted records into raw.locations_summary")
	return nil
}

func downInsertRawLocationsSummary(ctx context.Context, tx *sql.Tx) error {
	//goland:noinspection SqlWithoutWhere
	_, err := tx.Exec("DELETE FROM raw.locations_summary")
	if err != nil {
		return fmt.Errorf("failed to delete records from raw.locations_summary: %w", err)
	}
	return nil
}
