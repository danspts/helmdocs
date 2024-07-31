package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/danspts/helmdocs/pkg/generate/readme"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'generate' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
		generateCmd.Usage = func() {
			fmt.Fprintf(generateCmd.Output(), "Usage of %s:\n", "generate")
			fmt.Println("  readme      Generates README")
			fmt.Println("  values      Generates values")
			generateCmd.PrintDefaults()
		}

		if len(os.Args) < 3 {
			fmt.Println("expected 'readme' or 'values' subcommand")
			generateCmd.Usage()
			os.Exit(0)
		}

		command := "\033[1mhelmdocs\033[0m "
		command += strings.Join(os.Args[1:3], " ")
		switch os.Args[2] {
		case "readme":
			handleGenerateReadmeCommand(command, os.Args[3:])
		case "values":
			handleGenerateValuesCommand(command, os.Args[3:])
		default:
			fmt.Println("expected 'readme' or 'values' subcommand")
			generateCmd.Usage()
			os.Exit(1)
		}
	default:
		fmt.Println("expected 'generate' subcommand")
		os.Exit(1)
	}
}

func usage(command string, f *flag.FlagSet) func() {
	return func() {
		fmt.Fprintf(f.Output(), "Usage of %s:\n", command)
		f.PrintDefaults()
	}
}

func handleGenerateReadmeCommand(command string, args []string) {
	generateReadmeCmd := flag.NewFlagSet("readme", flag.ExitOnError)
	schemaPath := generateReadmeCmd.String("schema-path", "values.schema.json", "Path to the schema file")

	generateReadmeCmd.Usage = usage(command, generateReadmeCmd)

	generateReadmeCmd.Parse(args)
	readme.GenerateReadme(*schemaPath)
}

func handleGenerateValuesCommand(command string, args []string) {
	generateValuesCmd := flag.NewFlagSet("values", flag.ExitOnError)
	schemaPath := generateValuesCmd.String("schema-path", "values.schema.json", "Path to the schema file")
	omitDefault := generateValuesCmd.Bool("omit-default", false, "Omit default values")

	generateValuesCmd.Usage = usage(command, generateValuesCmd)

	generateValuesCmd.Parse(args)
	generateValues(*schemaPath, *omitDefault)
}

func generateValues(schemaPath string, omitDefault bool) {
	fmt.Printf("Generating values with schema path: %s, omit default: %t\n", schemaPath, omitDefault)
	// Add logic to generate values based on the schemaPath and omitDefault
}
