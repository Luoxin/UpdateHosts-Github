package main

import (
	"context"
	"github.com/Luoxin/Eutamias/utils"
	"github.com/alexflint/go-arg"
	"github.com/cloverstd/tcping/ping"
	"github.com/elliotchance/pie/pie"
	"github.com/go-resty/resty/v2"
	"github.com/letsfire/factory"
	dotDns "github.com/ncruces/go-dns"
	"github.com/pterm/pterm"
	"github.com/txn2/txeh"
	"net"
	"net/url"
	"sync"
	"time"
)

type DnsClient struct {
	dnsClientMap  map[string]*net.Resolver
	jsonClientMap map[string]bool
	apiClient     *resty.Client
	_lock         sync.RWMutex
}

func NewDnsClient() *DnsClient {
	return &DnsClient{
		dnsClientMap:  map[string]*net.Resolver{},
		jsonClientMap: map[string]bool{},
		apiClient: resty.New().
			SetTimeout(time.Second * 5).
			SetRetryMaxWaitTime(time.Second * 5).
			SetRetryWaitTime(time.Second),
	}
}

func (p *DnsClient) Added(nameserver string) bool {
	if p.alreadyExist(nameserver) {
		return false
	}

	u, err := url.Parse(nameserver)
	if err != nil {
		if p.tryAddDot(nameserver) {
			return true
		}
	}
	switch u.Scheme {
	case "http", "https":
		if p.tryAddDoh(nameserver) {
			return true
		} else if p.tryAddJSONApi(nameserver) {
			return true
		}
	case "tls", "":
		if p.tryAddDot(nameserver) {
			return true
		}
	default:
		if p.tryAddDoh(nameserver) {
			return true
		} else if p.tryAddDot(nameserver) {
			return true
		} else if p.tryAddJSONApi(nameserver) {
			return true
		}
	}

	return false
}

func (p *DnsClient) alreadyExist(nameserver string) bool {
	p._lock.RLock()
	defer p._lock.RUnlock()
	_, ok := p.dnsClientMap[nameserver]
	if ok {
		return true
	}
	return p.jsonClientMap[nameserver]
}

func (p *DnsClient) tryAddDoh(nameserver string) bool {
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

func (p *DnsClient) tryAddDot(nameserver string) bool {
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

func (p *DnsClient) tryAddJSONApi(nameserver string) bool {
	if len(p.lookupIPWithJsonApi(nameserver, "baidu.com")) == 0 {
		return false
	}

	p._lock.Lock()
	defer p._lock.Unlock()
	p.jsonClientMap[nameserver] = true
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

func (p *DnsClient) lookupIPWithJsonApi(nameserver string, domain string) (ips pie.Strings) {
	if nameserver == "" {
		return
	}

	_, err := p.apiClient.R().SetQueryParams(map[string]string{
		"name":  domain,
		"type":  "1",
		"short": "1",
	}).SetResult(&ips).Get(nameserver)
	if err != nil {
		pterm.Warning.Printfln("lookup err:%v", err)
		return
	}

	return
}

func (p *DnsClient) LookupIPWithNameserver(nameserver, domain string) (ips pie.Strings) {
	p._lock.RLock()
	client := p.dnsClientMap[nameserver]
	p._lock.RUnlock()
	return p.lookupIPWithClient(client, domain)
}

func (p *DnsClient) cloneDnsClient() (ips map[string]*net.Resolver) {
	newClientMap := make(map[string]*net.Resolver)
	p._lock.RLock()
	defer p._lock.RUnlock()
	for k, v := range p.dnsClientMap {
		newClientMap[k] = v
	}
	return newClientMap
}

func (p *DnsClient) cloneJSONApi() (ips []string) {
	p._lock.RLock()
	defer p._lock.RUnlock()
	for k := range p.jsonClientMap {
		ips = append(ips, k)
	}
	return
}

func (p *DnsClient) LookupIPFast(domain string) (ip string) {
	newClientMap := p.cloneDnsClient()
	jsonApiList := p.cloneJSONApi()
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

	lineDnsClient := worker.AddLine(func(i interface{}) {
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

	lineJSONApi := worker.AddLine(func(i interface{}) {
		ips := p.lookupIPWithJsonApi(i.(string), domain)
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
		lineDnsClient.Submit(client)
	}

	for client := range jsonApiList {
		lineJSONApi.Submit(client)
	}

	lineDnsClient.Wait()
	lineJSONApi.Wait()
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
		Protocol: ping.TCP,
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
}

func main() {
	arg.MustParse(&cmdArgs)

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

	client := NewDnsClient()

	lineNameserver := worker.AddLine(func(i interface{}) {
		client.Added(i.(string))
	})

	dnsClientList.Each(func(s string) {
		lineNameserver.Submit(s)
	})
	dnsJsonApiList.Each(func(s string) {
		lineNameserver.Submit(s)
	})

	lineNameserver.Wait()

	lineQuery := worker.AddLine(func(i interface{}) {
		domain := i.(string)
		fastIp := client.LookupIPFast(domain)
		if fastIp != "" {
			pterm.Info.Printfln("fastest: %v\t%v", domain, fastIp)
			hosts.AddHost(fastIp, domain)
		} else {
			pterm.Warning.Printfln("%v not found fast ip", domain)
		}
	})
	githubList.Each(func(domain string) {
		lineQuery.Submit(domain)
	})
	lineQuery.Wait()

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
