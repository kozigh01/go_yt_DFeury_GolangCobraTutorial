package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var (
	persistRootFlag bool
	localRootFlag bool
	times int
	rootCmd = &cobra.Command{
		Use: "cobra-app",
		Short: "An example cobra program",
		Long: `This is a simple example of a cobra program.
It will have several subcommands and flags.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from the root command")
		},
	}
	echoCommand = &cobra.Command{
		Use: "echo [strings to echo]",
		Short: "prints given strings to stdout",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
			p, _ := rootCmd.Flags().GetBool("persistFlag")
			fmt.Printf("Echo: %v - persistentFlag: %v\n", strings.Join(args, " "), p)
		},
	}
	timesCommand = &cobra.Command{
		Use: "times [strings to echo]",
		Short: "prints given strings to stdout multiple times",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if times == 4 {
				// return fmt.Errorf("don't use times = 4")
				return errors.New("don't use times == 4")
			}
			for i := 0; i < times; i++ {
				fmt.Printf("Echo times: %v - persistent flag: %v\n", strings.Join(args, " "), persistRootFlag)
			}
			return nil
		},
	}
	helloCommand = &cobra.Command {
		Use: "hello",
		Short: "hello",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires an argument that is 'good'")
			}
			if args[0] == "good" {
				return nil
			}
			return fmt.Errorf("invalid argument - was not 'good'")
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, World: %v\n", args)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "a persistent root flag")
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", true, "a local root flag")
	rootCmd.AddCommand(echoCommand)

	timesCommand.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo to stdout")
	timesCommand.MarkFlagRequired("times")
	echoCommand.AddCommand(timesCommand)

	// custom argument validation
	rootCmd.AddCommand(helloCommand)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}