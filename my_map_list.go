package main

import (
	"fmt"
	"time"
)

type K8sService struct {
	name string
	cluster_ip string
	namespace string
	selector [2] string
	pods [2] K8sPod
}

type K8sNameTime struct {
	time time.Time
	name string
	short_name string
	namespace string
}

type K8sPod struct {
	pod_ip	string
	cluster_name string
	labels [] string
	name_time_list [] K8sNameTime
}


func main() {

	fmt.Println("maps")
	services := make(map[string]K8sService)
	s := K8sService{name: "ten", cluster_ip: "10.10.10.10", namespace: "default"}
	services["service1"] = s
	fmt.Println("services = ", services)

	s2 := K8sService{name: "twenty", cluster_ip: "20.20.20.20", namespace: "default" }
	services["service2"] = s2

	s3 := K8sService{name: "thirty", cluster_ip: "30.30.30.30", namespace: "default" }
	services["service3"] = s3

	fmt.Println("services = ", services)
	pods := make(map[string] K8sPod)
	p := K8sPod{pod_ip: "99.99.99.99", cluster_name: "cluster99"}
	pods["99.99.99.99"] = p


	nt := K8sNameTime {
		time:       time.Now(),
		name:       "cups-1234567890",
		short_name: "cups",
		namespace:  "default",
	}

	p.name_time_list = append(p.name_time_list, nt)
	fmt.Println("p = ", p)

	p2 := K8sPod{pod_ip: "88.88.88.88", cluster_name: "cluster88"}
	pods["88.88.88.88"] = p2

	nt2 := K8sNameTime{
		time:       time.Now(),
		name:       "cups_db-1234567890",
		short_name: "cups_db",
		namespace:  "default",
	}

	p2.name_time_list = append(p2.name_time_list, nt2)
	fmt.Println("p2 = ", p2)

	key := "1.1.1.1"
	fmt.Println("Is key ",key, " present.")

	if val, ok := pods[key]; ok {
		fmt.Println("Found ", val)
	} else {
		fmt.Println("Not Found")
	}
}