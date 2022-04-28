package utils

import (
	"approvefeishu/structs"
	"fmt"
)

func Add(uuid string) bool {
	s := structs.Uuids
	if !in(uuid) {
		if len(s) >= 50 {
			s = s[1:]
			s = append(s, uuid)
			fmt.Println(s)
		} else {
			s = append(s, uuid)
		}
		return true
	}
	fmt.Printf("重复的uuid：%s", uuid)
	return false
}

func in(uuid string) bool {
	for _, v := range structs.Uuids {
		if v == uuid {
			return true
		}
	}
	return false
}
