package main

import "time"

type DnsLookupResult struct {
	Ip    string
	Delay time.Duration
}

//go:generate pie DnsLookupResultList.*
type DnsLookupResultList []*DnsLookupResult
