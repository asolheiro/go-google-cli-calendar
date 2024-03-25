package calendar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const AGENDA = "Estudar"

var (
	ErrAgendaNotFound = errors.New("agenda not found")
	ErrAddAgenda      = errors.New("error to add agenda")
	ErrEventsWeek     = errors.New("error to list week events")
	ErrEventDay       = errors.New("error to list day events")
)

type Calendar struct {
	Service    *gCalendar.Service
	CalendarId string
}

func NewClient() *Calendar {
	ctx := context.Background()
	credentials, err := os.ReadFile("./credentials.json")

	if err != nil {
		log.Fatal("Unable to read JSON credentials.")
	}

	service, err := gCalendar.NewService(
		ctx,
		option.WithCredentialsJSON(credentials),
	)

	if err != nil {
		log.Fatalf("Error creating Google Calendar service:\n%s", err.Error())
	}

	return &Calendar{
		Service: service,
	}
}

func (c *Calendar) InsertAgenda(id string) error {
	entry := &gCalendar.CalendarListEntry{
		Id: id,
	}
	_, err := c.Service.CalendarList.Insert(entry).Do()
	if err != nil {
		return ErrAddAgenda
	}
	return nil
}

func (c *Calendar) GetAgendaID() error {
	list, err := c.Service.CalendarList.List().Do()
	if err != nil {
		log.Fatal("Unable to list calendars linked to account.")
		return ErrAgendaNotFound
	}
	for _, v := range list.Items {
		if v.Summary == AGENDA {
			c.CalendarId = v.Id
		}
	}

	return nil
}

func (c *Calendar) ListWeekAgenda() error {
	now := time.Now()
	weekday := now.Weekday()
	startDate := now.AddDate(0, 0, -int(weekday))
	endDate := startDate.AddDate(0, 0, 7)

	events, err := c.Service.Events.List(c.CalendarId).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()

	if err != nil {
		return ErrEventsWeek
	}

	for _, v := range events.Items {
		fmt.Printf("%s | %s | at  %s.\n", v.Summary, v.Status, v.Start.DateTime)
	}

	return nil
}

func (c *Calendar) ListTodayAgenda() error {
	year, month, day := time.Now().Date()
	startDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 0, 1)

	events, err := c.Service.Events.List(c.CalendarId).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()

	if err != nil {
		return ErrEventDay
	}

	for _, v := range events.Items {
		fmt.Printf("%s | %s | at %s.\n", v.Summary, v.Status, v.Start.DateTime)
	}

	return nil
}
