package main

import (
	"context"
	"github.com/Luoxin/Eutamias/utils"
	"github.com/alexflint/go-arg"
	"github.com/cloverstd/tcping/ping"
	"github.com/elliotchance/pie/pie"
	"github.com/letsfire/factory"
	dotDns "github.com/ncruces/go-dns"
	"github.com/pterm/pterm"
	"github.com/txn2/txeh"
	"net"
	"strings"
	"sync"
	"time"
)

var text = `1.2.4.8
223.5.5.5
114.114.114.114
8.8.8.8
tls://dns.alidns.com
https://223.5.5.5/dns-query
tls://dns.rubyfish.cn:853
tls://1.0.0.1:853
tls://dns.google:853
tls://dns.alidns.com:853
tls://dns.cfiec.net:853
https://2400:3200:baba::1/dns-query
https://dns.cfiec.net/dns-query
https://223.6.6.6/dns-query
https://dns.alidns.com/dns-query
https://dns.ipv6dns.com/dns-query
tls://dns.ipv6dns.com:853
tls://dns.pub:853
tls://doh.pub:853
https://doh.pub/dns-query
180.76.76.76
119.29.29.29
119.28.28.28
https://dns.google/dns-query
https://dns.quad9.net/dns-query
https://dns11.quad9.net/dns-query
https://dns.twnic.tw/dns-query
https://1.1.1.1/dns-query
https://1.0.0.1/dns-query
https://cloudflare-dns.com/dns-query
https://dns.adguard.com/dns-query
https://doh.dns.sb/dns-query
tls://185.184.222.222:853
tls://185.222.222.222:853
https://doh-jp.blahdns.com/dns-query
https://public.dns.iij.jp/dns-query
https://v6.rubyfish.cn/dns-query
tls://v6.rubyfish.cn:853
https://[2001:4860:4860::6464]/dns-query
https://[2001:4860:4860::64]/dns-query
https://[2606:4700:4700::1111]/dns-query
https://[2606:4700:4700::64]/dns-query
tls://2a09::@853
tls://2a09::1@853
203.112.2.4
9.9.9.9
101.101.101.101
203.80.96.10
218.102.23.228
61.10.0.130
202.181.240.44
112.121.178.187
168.95.192.1
202.76.4.1
202.14.67.4
4.2.2.1
4.2.2.2
4.2.2.3
4.2.2.4
4.2.2.5
4.2.2.6`

type DnsClient struct {
	dnsClientMap map[string]*net.Resolver
	_lock        sync.RWMutex
}

func NewDnsClient() *DnsClient {
	return &DnsClient{
		dnsClientMap: map[string]*net.Resolver{},
	}
}

func (p *DnsClient) Added(nameserver string) bool {
	p._lock.RLock()
	client := p.dnsClientMap[nameserver]
	p._lock.RUnlock()
	if client != nil {
		return false
	}

	if p.tryAddGetDoh(nameserver) {
		return true
	} else if p.tryAddGetDot(nameserver) {
		return true
	}

	return false
}

func (p *DnsClient) tryAddGetDoh(nameserver string) bool {
	client, err := dotDns.NewDoHResolver(nameserver)
	if err != nil {
		//pterm.Error.Printfln("try add %v err:%v", nameserver, err)
		return false
	}

	if len(p.lookupIPWithClient(client, "baidu.com")) == 0 {
		return false
	}

	p._lock.Lock()
	defer p._lock.Unlock()
	p.dnsClientMap[nameserver] = client
	return true
}

func (p *DnsClient) tryAddGetDot(nameserver string) bool {
	client, err := dotDns.NewDoTResolver(nameserver)
	if err != nil {
		//pterm.Error.Printfln("try add %v err:%v", nameserver, err)
		return false
	}

	if len(p.lookupIPWithClient(client, "baidu.com")) == 0 {
		return false
	}

	p._lock.Lock()
	defer p._lock.Unlock()
	p.dnsClientMap[nameserver] = client
	return true
}

func (p *DnsClient) lookupIPWithClient(client *net.Resolver, domain string) (ips pie.Strings) {
	if client == nil {
		return
	}
	ctx := context.TODO()
	ctx, _ = context.WithTimeout(ctx, time.Second*3)
	r, err := client.LookupIP(ctx, "ip4", domain)
	if err != nil {
		pterm.Warning.Printfln("lookup err:%v", err)
		return
	}

	for _, a := range r {
		ips = append(ips, a.To4().String())
	}

	return
}

func (p *DnsClient) LookupIPWithNameserver(nameserver, domain string) (ips pie.Strings) {
	p._lock.RLock()
	client := p.dnsClientMap[nameserver]
	p._lock.RUnlock()
	return p.lookupIPWithClient(client, domain)
}

func (p *DnsClient) cloneClient() (ips map[string]*net.Resolver) {
	newClientMap := make(map[string]*net.Resolver)
	p._lock.RLock()
	defer p._lock.RUnlock()
	for k, v := range p.dnsClientMap {
		newClientMap[k] = v
	}
	return newClientMap
}

func (p *DnsClient) LookupIP(domain string) (ips pie.Strings) {
	newClientMap := p.cloneClient()
	worker := factory.NewMaster(20, 8)
	lineLookup := worker.AddLine(func(i interface{}) {
		ips = append(ips, p.lookupIPWithClient(i.(*net.Resolver), domain)...)
	})

	for _, client := range newClientMap {
		lineLookup.Submit(client)
	}

	return ips.Unique()
}

func (p *DnsClient) LookupIPFast(domain string) (ip string) {
	newClientMap := p.cloneClient()
	worker := factory.NewMaster(8, 2)

	var _lock sync.Mutex

	alreadyMap := map[string]bool{}

	var minIp string
	var minDelay = time.Duration(-1)

	linePing := worker.AddLine(func(i interface{}) {
		ip := i.(string)
		delay := p.Ping(ip)
		pterm.Info.Printfln("%v\t%v", i, delay)
		_lock.Lock()
		defer _lock.Unlock()
		if delay <= minDelay && minDelay >= 0 {
			return
		}
		minDelay = delay
		minIp = ip
	})

	lineLookup := worker.AddLine(func(i interface{}) {
		ips := p.lookupIPWithClient(i.(*net.Resolver), domain)
		ips.Each(func(ip string) {
			_lock.Lock()
			ok := alreadyMap[ip]
			if !ok {
				alreadyMap[ip] = true
				pterm.Info.Printfln("%v\t%v", domain, ip)
			}
			_lock.Unlock()
			if ok {
				return
			}
			linePing.Submit(ip)
			pterm.Info.Printfln("add %v to test connectivity", ip)
		})
	})

	for _, client := range newClientMap {
		lineLookup.Submit(client)
	}

	lineLookup.Wait()
	linePing.Wait()

	return minIp
}

func (p *DnsClient) Ping(ip string) time.Duration {
	target := ping.Target{
		Timeout:  time.Second,
		Interval: time.Second,
		Host:     ip,
		Port:     80,
		Counter:  1,
		Protocol: ping.HTTPS,
	}

	pinger := ping.NewTCPing()
	pinger.SetTarget(&target)
	pingerDone := pinger.Start()
	<-pingerDone
	if pinger.Result().Failed() >= 1 {
		return -1
	}

	return pinger.Result().Avg()
}

var githubList = pie.Strings{
	"alive.github.com",
	"live.github.com",
	"github.githubassets.com",
	"central.github.com",
	"desktop.githubusercontent.com",
	"assets-cdn.github.com",
	"camo.githubusercontent.com",
	"github.map.fastly.net",
	"github.global.ssl.fastly.net",
	"gist.github.com",
	"github.io",
	"github.com",
	"github.blog",
	"api.github.com",
	"raw.githubusercontent.com",
	"user-images.githubusercontent.com",
	"favicons.githubusercontent.com",
	"avatars5.githubusercontent.com",
	"avatars4.githubusercontent.com",
	"avatars3.githubusercontent.com",
	"avatars2.githubusercontent.com",
	"avatars1.githubusercontent.com",
	"avatars0.githubusercontent.com",
	"avatars.githubusercontent.com",
	"codeload.github.com",
	"github-cloud.s3.amazonaws.com",
	"github-com.s3.amazonaws.com",
	"github-production-release-asset-2e65be.s3.amazonaws.com",
	"github-production-user-asset-6210df.s3.amazonaws.com",
	"github-production-repository-file-5c1aeb.s3.amazonaws.com",
	"githubstatus.com",
	"github.community",
	"media.githubusercontent.com",
}

var cmdArgs struct {
	HostsFile string `arg:"-h,--hosts" help:"hosts file path"`
	Action    string `arg:"-a,--action" help:""`
}

var client *DnsClient

func Init() {
	worker := factory.NewMaster(8, 2)

	client = NewDnsClient()

	lineNameserver := worker.AddLine(func(i interface{}) {
		client.Added(i.(string))
	})

	list := pie.Strings(strings.Split(text, "\n"))
	list.Each(func(s string) {
		lineNameserver.Submit(s)
	})

	lineNameserver.Wait()
}

func UpdateHosts() {
	var err error
	var hosts *txeh.Hosts
	if utils.FileExists(cmdArgs.HostsFile) {
		hosts, err = txeh.NewHosts(&txeh.HostsConfig{
			ReadFilePath:  cmdArgs.HostsFile,
			WriteFilePath: cmdArgs.HostsFile,
		})
		if err != nil {
			pterm.Error.Printfln("err:%v", err)
			return
		}
	} else {
		hosts, err = txeh.NewHostsDefault()
		if err != nil {
			pterm.Error.Printfln("err:%v", err)
			return
		}
	}

	worker := factory.NewMaster(8, 2)

	var _lock sync.Mutex
	hostMap := map[string]string{}
	lineQuery := worker.AddLine(func(i interface{}) {
		domain := i.(string)
		fastIp := client.LookupIPFast(domain)
		if fastIp != "" {
			pterm.Info.Printfln("fastest: %v\t%v", domain, fastIp)
			hosts.AddHost(fastIp, domain)
			_lock.Lock()
			hostMap[domain] = fastIp
			_lock.Unlock()
		} else {
			pterm.Warning.Printfln("%v not found fast ip", domain)
		}
	})
	githubList.Each(func(domain string) {
		lineQuery.Submit(domain)
	})
	lineQuery.Wait()

	for doamin, ip := range hostMap {
		pterm.Printfln("%s\t%s", ip, doamin)
	}

	if cmdArgs.HostsFile != "" {
		err = hosts.SaveAs(cmdArgs.HostsFile)
	} else {
		err = hosts.Save()
	}
	if err != nil {
		pterm.Error.Printfln("save hosts err:%v", err)
		return
	}
}

func main() {
	arg.MustParse(&cmdArgs)

	switch cmdArgs.Action {
	case "run":
		Init()
		UpdateHosts()
		for {
			select {
			case <- time.After():

			}
		}
	default:
		Init()
		UpdateHosts()
	}
}
