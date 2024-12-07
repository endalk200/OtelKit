/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the provided json file against a schema",
	Run: func(cmd *cobra.Command, args []string) {
		schemaFile, _ := cmd.Flags().GetString("schema")
		jsonFile, _ := cmd.Flags().GetString("target")

		if schemaFile == "" || jsonFile == "" {
			fmt.Println("Error: both --schema and --target flags must be specified")
			os.Exit(1)
		}

		err := validateJSON(schemaFile, jsonFile)
		if err != nil {
			fmt.Println("Validation failed:", err)
			os.Exit(1)
		}
		fmt.Println("Validation successful!")
	},
}

func validateJSON(schemaFile, jsonFile string) error {
	// Convert schema and JSON file paths to absolute paths
	absSchemaFile, err := filepath.Abs(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to resolve schema file path: %w", err)
	}

	absJSONFile, err := filepath.Abs(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to resolve JSON file path: %w", err)
	}

	// Load schema and JSON files
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + absSchemaFile)
	jsonLoader := gojsonschema.NewReferenceLoader("file://" + absJSONFile)

	// Perform validation
	result, err := gojsonschema.Validate(schemaLoader, jsonLoader)
	if err != nil {
		return fmt.Errorf("error during validation: %w", err)
	}

	// Check if validation passed
	if !result.Valid() {
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		return fmt.Errorf("JSON validation failed")
	}

	return nil
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Add flags for the validate command
	validateCmd.Flags().StringP("schema", "s", "", "Path to the JSON schema file (required)")
	validateCmd.Flags().StringP("target", "t", "", "Path to the JSON file to validate (required)")

	// Mark flags as required
	validateCmd.MarkFlagRequired("schema")
	validateCmd.MarkFlagRequired("target")
}
