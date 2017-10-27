package entity

import (
	"fmt"
	"strconv"
)

// Date .
type Date struct {
	Year, Month, Day, Hour, Minute int
}

func (mdate Date) init(tyear, tmonth, tday, thour, tminute int) {
	mdate.Year = tyear
	mdate.Month = tmonth
	mdate.Day = tday
	mdate.Hour = thour
	mdate.Minute = tminute
}

// GetYear .
func (mdate Date) GetYear() int {
	return mdate.Year
}

// SetYear .
func (mdate Date) SetYear(tyear int) {
	mdate.Year = tyear
}

// GetMonth .
func (mdate Date) GetMonth() int {
	return mdate.Month
}

func (mdate Date) SetMonth(tmonth int) {
	mdate.Month = tmonth
}

func (mdate Date) GetDay() int {
	return mdate.Day
}

func (mdate Date) SetDay(tday int) {
	mdate.Day = tday
}

func (mdate Date) GetHour() int {
	return mdate.Hour
}

func (mdate Date) SetHour(thour int) {
	mdate.Hour = thour
}

func (mdate Date) GetMinute() int {
	return mdate.Minute
}

func (mdate Date) SetMinute(tminute int) {
	mdate.Minute = tminute
}

func StringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("String to Int fail")
	}
	return result
}

func IsValid(tdate Date) bool {
	currentYear := tdate.GetYear()
	currentMonth := tdate.GetMonth()
	currentDay := tdate.GetDay()
	currentHour := tdate.GetHour()
	currentMinute := tdate.GetMinute()
	day := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if currentYear > 9999 || currentYear < 1000 {
		return false
	}
	if currentMonth > 12 || currentMonth < 1 {
		return false
	}
	if (currentYear%4 == 0 && currentYear%100 != 0) || currentYear%400 == 0 {
		day[2] = 29
	}
	if currentDay > day[currentMonth] || currentDay < 1 {
		return false
	}
	if currentHour > 23 || currentHour < 0 {
		return false
	}
	if currentMinute > 59 || currentMinute < 0 {
		return false
	}
	return true
}

func StringToDate(tdatestring string) Date {
	var resultDate Date
	if len(tdatestring) != 16 {
		fmt.Println("the len of the ", tdatestring, "change to Date isn't 16")
	}
	count := 0
	for count < len(tdatestring) {
		switch count {
		case 4:
			if tdatestring[4] != '-' {
				fmt.Println("the form of the Date is wrong!")
				return resultDate
			} else {
				ts := tdatestring[0:4]
				tyear, _ := strconv.Atoi(ts)
				resultDate.SetYear(tyear)
			}
		case 7:
			if tdatestring[7] != '-' {
				fmt.Println("the form of the Date is wrong!")
				return resultDate
			} else {
				ts := tdatestring[5:7]
				tmonth, _ := strconv.Atoi(ts)
				resultDate.SetMonth(tmonth)
			}
		case 10:
			if tdatestring[10] != '/' {
				fmt.Println("the form of the Date is wrong!")
				return resultDate
			} else {
				ts := tdatestring[8:10]
				tday, _ := strconv.Atoi(ts)
				resultDate.SetDay(tday)
			}
		case 13:
			if tdatestring[13] != ':' {
				fmt.Println("the form of the Date is wrong!")
				return resultDate
			} else {
				ts := tdatestring[11:13]
				thour, _ := strconv.Atoi(ts)
				resultDate.SetHour(thour)
			}
		case 15:
			if tdatestring[count] < '0' || tdatestring[count] > '9' {
				fmt.Println("the form of the Date is wrong!")
				return resultDate
			} else {
				ts := tdatestring[14:16]
				tminute, _ := strconv.Atoi(ts)
				resultDate.SetMinute(tminute)
			}
		default:
			if tdatestring[count] < '0' || tdatestring[count] > '9' {
				fmt.Println("the form of the Date is wrong!")
				return resultDate
			}
		}
	}
	return resultDate
}

func IntToString(a int) string {
	resultstring := strconv.Itoa(a)
	return resultstring
}

func DateToString(tdate Date) string {
	initTime := "0000-00-00/00:00"
	if !IsValid(tdate) {
		dateString := initTime
		return dateString
	}
	dateString := IntToString(tdate.GetYear()) + "-" + IntToString(tdate.GetMonth()) +
		"-" + IntToString(tdate.GetDay()) + "/" + IntToString(tdate.GetHour()) + ":" + IntToString(tdate.GetMinute())
	return dateString
}

func (mdate Date) CopyDate(tdate Date) Date {
	mdate.SetYear(tdate.GetYear())
	mdate.SetMonth(tdate.GetMonth())
	mdate.SetDay(tdate.GetDay())
	mdate.SetMinute(tdate.GetMinute())
	mdate.SetHour(tdate.GetHour())
	return mdate
}

func (mdate Date) IsSameDate(tdate Date) bool {
	return (tdate.GetYear() == mdate.GetYear() &&
		tdate.GetMonth() == mdate.GetMonth() &&
		tdate.GetDay() == mdate.GetDay() &&
		tdate.GetHour() == mdate.GetHour() &&
		tdate.GetMinute() == mdate.GetMinute())
}

func (mdate Date) MoreThan(tdate Date) bool {
	if mdate.Year > tdate.GetYear() {
		return true
	}
	if mdate.Year < tdate.GetYear() {
		return false
	}
	if mdate.Month > tdate.GetMonth() {
		return true
	}
	if mdate.Month < tdate.GetMonth() {
		return false
	}
	if mdate.Day > tdate.GetDay() {
		return true
	}
	if mdate.Day < tdate.GetDay() {
		return false
	}
	if mdate.Hour > tdate.GetHour() {
		return true
	}
	if mdate.Hour < tdate.GetHour() {
		return false
	}
	if mdate.Minute > tdate.GetMinute() {
		return true
	}
	if mdate.Minute < tdate.GetMinute() {
		return false
	}
	return false
}

func (mdate Date) LessThan(tdate Date) bool {
	if mdate.IsSameDate(tdate) == false && !mdate.MoreThan(tdate) == false {
		return false
	}
	return true
}

func (mdate Date) MoreOrEqual(tdate Date) bool {
	return mdate.IsSameDate(tdate) || mdate.MoreThan(tdate)
}

func (mdate Date) LessOrEqual(tdate Date) bool {
	return !mdate.MoreThan(tdate)
}
