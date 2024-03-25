package cmd

import (
	"log"

	"github.com/rmndvngrpslhr/go-cli-calendar/internal/calendar"
	"github.com/spf13/cobra"
)

var AgendaCmd = &cobra.Command{
	Use:   "agenda",
	Short: "List all your agendas",
	Long:  "Check all your agendas",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.InsertAgenda(args[0])

		if err != nil {
			log.Fatal("Unable to insert agenda")
		}
	},
}
