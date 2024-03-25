package cmd

import (
	"fmt"
	agenda "github.com/rmndvngrpslhr/go-cli-calendar/cmd/agenda"
	events "github.com/rmndvngrpslhr/go-cli-calendar/cmd/events"
	"github.com/spf13/cobra"
	"os"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "calendar",
		Short:         "Your calendar CLI",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(
		events.EventsCmd,
		agenda.AgendaCmd,
	)

	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
