package notification

import (
	"testing"
	"time"
)

func TestChronForDaysAndHourAndMinutes(t *testing.T) {
	time := time.Date(2024, 6, 2, 0, 0, 5, 0, time.UTC)

	isCron, _ := IsTimeToCron(time, 2, 0, 0)
	if !isCron {
		t.Fatalf("Cron should be true")
	}

	isCron2, _ := IsTimeToCron(time, 2, 1, 0)

	if isCron2 {
		t.Fatalf("Cron should be false")
	}

	isCron3, _ := IsTimeToCron(time, 2, 0, 1)

	if isCron3 {
		t.Fatalf("Cron should be false")
	}

	isCron4, _ := IsTimeToCron(time, 2, 0, 1)

	if isCron4 {
		t.Fatalf("Cron should be false")
	}

	isCron5, _ := IsTimeToCron(time, 23, 0, 0)

	if isCron5 {
		t.Fatalf("Cron should be false")
	}
}

func TestForAnyHourAndMinute(t *testing.T) {
	time := time.Date(2024, 6, 2, 0, 0, 5, 0, time.UTC)

	isCron, _ := IsTimeToCron(time, 2, CronAny, 0)
	if !isCron {
		t.Fatalf("Cron should be true")
	}

	isCron1, _ := IsTimeToCron(time, 2, 0, CronAny)
	if !isCron1 {
		t.Fatalf("Cron should be true")
	}

	isCron2, _ := IsTimeToCron(time, 2, 1, CronAny)

	if isCron2 {
		t.Fatalf("Cron should be false")
	}

	isCron3, _ := IsTimeToCron(time, 2, CronAny, 1)

	if isCron3 {
		t.Fatalf("Cron should be false")
	}

	isCron5, _ := IsTimeToCron(time, CronAny, 0, 0)

	if !isCron5 {
		t.Fatalf("Cron should be TRUE")
	}

	isCron4, _ := IsTimeToCron(time, CronAny, CronAny, CronAny)

	if !isCron4 {
		t.Fatalf("Cron should be true")
	}
}

func TestHourShouldBeBetween0And23(t *testing.T) {
	time := time.Date(2024, 6, 2, 0, 0, 5, 0, time.UTC)

	_, err := IsTimeToCron(time, CronAny, 24, 0)
	if err == nil {
		t.Fatalf("Cron should err when hour is greater than 23")
	}

	_, err0 := IsTimeToCron(time, CronAny, 23456765432, 0)
	if err0 == nil {
		t.Fatalf("Cron should err when hour is greater than 23")
	}

	_, err1 := IsTimeToCron(time, CronAny, -1, 0)
	if err1 == nil {
		t.Fatalf("Cron should err when hour is less than 0")
	}

	_, err2 := IsTimeToCron(time, CronAny, -98674753, 0)
	if err2 == nil {
		t.Fatalf("Cron should err when hour is less than 0")
	}

	_, err3 := IsTimeToCron(time, CronAny, CronAny, 0)
	if err3 != nil {
		t.Fatalf("Cron should not err when hour is Any")
	}
}

func TestMinuteShouldBeBetween0And59(t *testing.T) {
	time := time.Date(2024, 6, 2, 0, 0, 5, 0, time.UTC)

	_, err := IsTimeToCron(time, CronAny, 0, 60)
	if err == nil {
		t.Fatalf("Cron should err when hour is greater than 23")
	}

	_, err0 := IsTimeToCron(time, CronAny, 0, 23456765432)
	if err0 == nil {
		t.Fatalf("Cron should err when hour is greater than 23")
	}

	_, err1 := IsTimeToCron(time, CronAny, 0, -1)
	if err1 == nil {
		t.Fatalf("Cron should err when hour is less than 0")
	}

	_, err2 := IsTimeToCron(time, CronAny, 0, -98674753)
	if err2 == nil {
		t.Fatalf("Cron should err when hour is less than 0")
	}

	_, err3 := IsTimeToCron(time, CronAny, 0, CronAny)
	if err3 != nil {
		t.Fatalf("Cron should not err when hour is Any")
	}
}

func TestDaysShouldBeBetween1And31(t *testing.T) {
	time := time.Date(2024, 6, 2, 0, 0, 5, 0, time.UTC)

	_, err := IsTimeToCron(time, 32, CronAny, CronAny)
	if err == nil {
		t.Fatalf("Cron should err when days is greater than 31")
	}

	_, err0 := IsTimeToCron(time, 34567654, CronAny, CronAny)
	if err0 == nil {
		t.Fatalf("Cron should err when days is greater than 31")
	}

	_, err1 := IsTimeToCron(time, 0, CronAny, CronAny)
	if err1 == nil {
		t.Fatalf("Cron should err when hour is less than 1")
	}

	_, err2 := IsTimeToCron(time, -2341, CronAny, CronAny)
	if err2 == nil {
		t.Fatalf("Cron should err when hour is less than 0")
	}

	_, err3 := IsTimeToCron(time, CronAny, CronAny, CronAny)
	if err3 != nil {
		t.Fatalf("Cron should not err when hour is Any")
	}
}
