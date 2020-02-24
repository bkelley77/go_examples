package main

import (
	"container/list"
	"fmt"
	"time"
)

type MyK8sPod struct {
	pod_ip         time.Time
	cluster_name   string
	name_time_list *list.List
}

type MyK8sNameTime struct {
	time time.Time
	name string
	short_name string
	namespace string
}

func main() {

	pods := make(map[string]MyK8sPod)

	p1 := MyK8sPod {
		pod_ip:       		time.Now(),
		cluster_name:       "cups-1234567890",
		name_time_list:		list.New(),
	}

	pods["1.1.1.1"] = p1

	// insert new entry at the front of the list
	nt1 := MyK8sNameTime {
		time:       time.Now(),
		name:       "cups-1234567890",
		short_name: "cups",
		namespace:  "default",
	}
	p1.name_time_list.PushFront(nt1)

	// insert new entry at the front of the list
	nt2 := MyK8sNameTime {
		time:       time.Now(),
		name:       "cups-db-1234567890",
		short_name: "cups-db",
		namespace:  "default",
	}
	p1.name_time_list.PushFront(nt2)

	for _, pod_element := range pods {
        println("in pod loop...")
		for e := pod_element.name_time_list.Front(); e != nil; e = e.Next() {
			fmt.Println("\nname        : ", e.Value.(MyK8sNameTime).name)
			fmt.Println("short_name  : ", e.Value.(MyK8sNameTime).short_name)
			fmt.Println("time        : ", e.Value.(MyK8sNameTime).time)
			fmt.Println("namespace   : ", e.Value.(MyK8sNameTime).namespace)
		}

	}
}