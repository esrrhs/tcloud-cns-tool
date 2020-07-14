package main

import (
	"flag"
	"log"
	"net/url"
)

func main() {

	secretId := flag.String("secretId", "", "secretId")
	secretKey := flag.String("secretKey", "", "secretKey")
	domain := flag.String("domain", "", "domain")
	list := flag.Bool("l", false, "list")
	set := flag.Bool("s", false, "set")
	name := flag.String("name", "", "set name")
	value := flag.String("value", "", "set value")
	settype := flag.String("type", "A", "set type")

	flag.Parse()

	if *secretId == "" || *secretKey == "" || *domain == "" {
		flag.Usage()
		return
	}

	if *set && (*name == "" || *value == "") {
		flag.Usage()
		return
	}

	cli := New(*secretId, *secretKey)
	params := url.Values{}
	params.Set("offset", "0")
	params.Set("length", "1000")
	domainsResp, err := cli.DomainList(params)
	if err != nil {
		log.Fatal(err)
	}

	var dst *Domain

	for _, d := range domainsResp.Data.Domains {
		if d.Name == *domain {
			dst = &d
		}
	}

	if dst == nil {
		log.Fatalf("not find %s", *domain)
	}

	log.Printf("find %s", dst.Name)

	records, err := cli.RecordList(*domain, "0", "1000")
	if err != nil {
		log.Fatal(err)
	}

	if *list {
		for _, record := range records {
			log.Println(record.Name + " " + record.Value)
		}
		return
	}

	if *set {
		for _, record := range records {
			if record.Name == *name {
				res, err := cli.RecordDelete(*domain, record.Id)
				if err != nil {
					log.Fatal(err)
				}
				if res.Code != 0 {
					log.Fatalf("RecordDelete fail %s", *name)
				}
			}
		}

		record := Record{Name: *name, Type: *settype, Line: "默认", Value: *value}
		res, err := cli.RecordCreate(*domain, record)
		if err != nil {
			log.Fatal(err)
		}
		if res.Code != 0 {
			log.Println(res.Message)
		} else {
			log.Printf("set ok %s %s", *name, *value)
		}
	}
}
