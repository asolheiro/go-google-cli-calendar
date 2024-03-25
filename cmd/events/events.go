package cmd

import (
	"fmt"
	"log"

	"github.com/rmndvngrpslhr/go-cli-calendar/internal/calendar"
	"github.com/spf13/cobra"
)

func init() {
	EventsCmd.AddCommand(
		EventsWeekCmd,
		EventsTodayCmd,
	)
}

var EventsCmd = &cobra.Command{
	Use:   "events",
	Short: "list all your calendar events",
	Long:  "Check all your scheduled events",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.GetAgendaID()

		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Agenda ID: ", c.CalendarId)
	},
}
