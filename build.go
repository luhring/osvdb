package osvdb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/luhring/osvdb/model"
	"github.com/luhring/osvdb/osv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"os"
)

// Build specifies the configuration for building a database.
type Build struct {
	// OutputDatabaseLocation is the location of the SQLite database to write
	OutputDatabaseLocation string

	// OverwriteDatabase specifies whether to overwrite the database if it already exists
	OverwriteDatabase bool
}

// Do builds the database.
func (cfg Build) Do(ctx context.Context, inputs ...Input) error {
	if len(inputs) == 0 {
		return fmt.Errorf("no build inputs")
	}

	if cfg.OutputDatabaseLocation == "" {
		return fmt.Errorf("build output database location was empty")
	}

	_, err := os.Stat(cfg.OutputDatabaseLocation)
	outputDatabaseExistsAlready := err == nil
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to stat output database location: %w", err)
		}

		outputDatabaseExistsAlready = false
	}

	if outputDatabaseExistsAlready {
		if !cfg.OverwriteDatabase {
			return fmt.Errorf("specified file already exists: %s", cfg.OutputDatabaseLocation)
		}

		err = os.Remove(cfg.OutputDatabaseLocation)
		if err != nil {
			return fmt.Errorf("failed to remove existing database in preparation for new database build: %w", err)
		}
	}

	db, err := gorm.Open(sqlite.Open(cfg.OutputDatabaseLocation), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	err = db.AutoMigrate(
		&model.Vulnerability{},
		&model.Alias{},
		&model.Related{},
		&model.Severity{},
		&model.Affected{},
		&model.Package{},
		&model.Range{},
		&model.Version{},
		&model.Event{},
		&model.Reference{},
		&model.Credit{},
		&model.Contact{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate schema: %w", err)
	}

	// TODO: Consider running these concurrently.
	for _, input := range inputs {
		err := cfg.addInputToDatabase(ctx, db, input)
		if err != nil {
			return fmt.Errorf("failed to add input to database: %w", err)
		}
	}

	return nil
}

func (cfg Build) addInputToDatabase(ctx context.Context, db *gorm.DB, input Input) error {
	if input == nil {
		return nil
	}

	reader, cleanup, err := input(ctx)
	if cleanup != nil {
		defer cleanup()
	}
	if err != nil {
		return fmt.Errorf("failed to get input: %w", err)
	}
	dec := json.NewDecoder(reader)

	for {
		var document osv.Document
		if err := dec.Decode(&document); err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("failed to decode JSON: %w", err)
		}

		vulnerability := documentToModel(document)

		// TODO: Consider using a transaction here and/or batching these up.
		db.WithContext(ctx).Create(&vulnerability)
	}

	return nil
}
