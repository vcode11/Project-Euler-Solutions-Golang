package main

import "fmt"

func main(){
	month := 1
	year := 1901
	day_of_week := 1
	days_in_month := map[int]int{
			1:31,
			3:31,
			4:30,
			5:31,
			6:30,
			7:31,
			8:31,
			9:30,
			10:31,
			11:30,
			12:31,
	}
	sundays:=0
	for year !=  2001{
		if day_of_week == 6{
			sundays+=1
		}
		if val, ok := days_in_month[month]; ok{
			day_of_week = (day_of_week + val)%7
		} else{
			days := 28
			if year%4 == 0{
				if year%400 == 0{
					days = 29
				} else if year%100 == 0{
					days = 28
				} else {
					days = 29
				}
			} else {
				days = 28
			}
			day_of_week = (day_of_week + days)%7
		}
		month++
		if month == 13{
			month = 1
			year++
		}

	}
	fmt.Println(sundays)
}
