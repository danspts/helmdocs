package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/danspts/helmdocs/pkg/generate/readme"
	"github.com/danspts/helmdocs/pkg/generate/values"
	"github.com/danspts/helmdocs/pkg/types"
)

var (
	version = "dev"
	commit  = ""
)

func helmdocs(args []string) {
	if len(args) < 2 {
		fmt.Printf("version: %s, commit:%s\n", version, commit)
		fmt.Println("expected 'generate' subcommand")
		panic("")
	}

	switch args[1] {
	case "generate":
		generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
		generateCmd.Usage = func() {
			fmt.Fprintf(generateCmd.Output(), "Usage of %s:\n", "generate")
			fmt.Println("  readme      Generates README")
			fmt.Println("  values      Generates values")
			generateCmd.PrintDefaults()
		}

		if len(args) < 3 {
			fmt.Println("expected 'readme' or 'values' subcommand")
			generateCmd.Usage()
			os.Exit(0)
		}

		command := "\033[1mhelmdocs\033[0m "
		command += strings.Join(args[1:3], " ")
		switch args[2] {
		case "readme":
			handleGenerateReadmeCommand(command, args[3:])
		case "values":
			handleGenerateValuesCommand(command, args[3:])
		default:
			fmt.Println("expected 'readme' or 'values' subcommand")
			generateCmd.Usage()
			panic("")
		}
	default:
		fmt.Println("expected 'generate' subcommand")
		panic("")
	}
}

func usage(command string, f *flag.FlagSet) func() {
	return func() {
		fmt.Fprintf(f.Output(), "Usage of %s:\n", command)
		f.PrintDefaults()
	}
}

type Configer interface {
	readSchema(filename string) (types.Schema, error)
	writeOutput(filename, content string)
}

type Config struct {
	filename, output string
}

func readSchema(filename string) (types.Schema, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return types.Schema{}, err
	}
	var schema types.Schema
	if err := json.Unmarshal(file, &schema); err != nil {
		return types.Schema{}, err
	}
	return schema, nil
}

func writeOutput(filename, content string) error {
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}

func handleGenerateReadmeCommand(command string, args []string) {
	generateReadmeCmd := flag.NewFlagSet("readme", flag.ExitOnError)
	schemaPath := generateReadmeCmd.String("schema-path", "values.schema.json", "Path to the schema file")
	output := generateReadmeCmd.String("output", "./README.md", "Path to the generatedValue")
	generateReadmeCmd.Usage = usage(command, generateReadmeCmd)

	generateReadmeCmd.Parse(args)
	schema, err := readSchema(*schemaPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	readme := readme.GenerateReadme(schema)
	err = writeOutput(*output, readme)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(*output + " generated successfully.")
}

func handleGenerateValuesCommand(command string, args []string) {
	generateValuesCmd := flag.NewFlagSet("values", flag.ExitOnError)
	schemaPath := generateValuesCmd.String("schema-path", "values.schema.json", "Path to the schema file")
	output := generateValuesCmd.String("output", "./values.yaml", "Path to the generatedValue")
	omitDefault := generateValuesCmd.Bool("omit-default", true, "Omit default values")

	generateValuesCmd.Usage = usage(command, generateValuesCmd)

	generateValuesCmd.Parse(args)
	schema, err := readSchema(*schemaPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	values := values.GenerateValues(schema, *omitDefault)
	err = writeOutput(*output, values)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(*output + " generated successfully.")
}

func main() {
	helmdocs(os.Args)
}
