package notification

import (
	"errors"
	"time"
)

const CronAny = -9999

type CronDay = int
type CronHour = int
type CronMin = int

func IsTimeToCron(time time.Time, day CronDay, hour CronHour, minute CronMin) (bool, error) {
	if day != CronAny && day < 1 || day > 31 {
		return false, errors.New("Day should be between 1 and 31")
	}
	if hour != CronAny && hour < 0 || hour > 23 {
		return false, errors.New("Hour should be between 0 and 23")
	}
	if minute != CronAny && minute < 0 || minute > 59 {
		return false, errors.New("Minute should be between 0 and 59")
	}

	var isDay = day == CronAny || time.Day() == day
	var isHour = hour == CronAny || time.Hour() == hour
	var isMinute = minute == CronAny || time.Minute() == minute
	return isDay && isHour && isMinute, nil
}
