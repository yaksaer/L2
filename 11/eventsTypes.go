package main

import "time"

func DayEvents(id string, date time.Time) []Info {
	var result []Info
	RWmtx.RLock()
	defer RWmtx.RUnlock()
	if val, ok := Events[id]; ok {
		for _, v := range val.inf {
			if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() && v.Date.Day() == date.Day() {
				result = append(result, v)
			}
		}
	}
	return result
}
func WeekEvents(id string, date time.Time) []Info {
	RWmtx.RLock()
	defer RWmtx.RUnlock()
	var result []Info
	year, week := date.ISOWeek()

	if val, ok := Events[id]; ok {
		for _, v := range val.inf {
			year2, week2 := v.Date.ISOWeek()
			if year == year2 && week == week2 {
				result = append(result, v)
			}
		}
	}
	return result
}

func MonthEvents(id string, date time.Time) []Info {
	RWmtx.RLock()
	defer RWmtx.RUnlock()
	var result []Info

	if val, ok := Events[id]; ok {
		for _, v := range val.inf {
			if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() {
				result = append(result, v)
			}
		}
	}
	return result
}
