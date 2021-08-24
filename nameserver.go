package main

import (
	"github.com/elliotchance/pie/pie"
)

var dnsClientList = pie.Strings{
	"23.106.151.177",
	"208.67.222.222",
	"4.2.2.3",
	"4.2.2.4",
	"4.2.2.5",
	"4.2.2.6",
	"tls://185.184.222.222:853",
	"tls://185.222.222.222:853",
	"tls://v6.rubyfish.cn:853",
	"tls://2a09::@853",
	"tls://2a09::1@853",
	"114.114.114.114",
	"114.114.115.115",
	"114.114.114.119",
	"114.114.115.119",
	"114.114.114.110",
	"114.114.115.110",
	"223.5.5.5",
	"223.6.6.6",
	"https://223.5.5.5/dns-query",
	"https://223.6.6.6/dns-query",
	"https://dns.alidns.com/dns-query",
	"tls://dns.alidns.com",
	"180.76.76.76",
	"119.29.29.29",
	"182.254.116.116",
	"119.28.28.28",
	"182.254.118.118",
	"https://doh.pub/dns-query",
	"tls://dns.pub",
	"tls://doh.pub",
	"1.2.4.8",
	"210.2.4.8",
	"117.50.11.11",
	"52.80.66.66",
	"117.50.10.10",
	"52.80.52.52",
	"117.50.60.30",
	"52.80.60.30",
	"101.226.4.6",
	"218.30.118.6",
	"123.125.81.6",
	"140.207.198.6",
	"1.1.8.8",
	"1.1.8.9",
	"https://doh.360.cn/dns-query",
	"https://dns.cfiec.net/dns-query",
	"tls://dns.cfiec.net",
	"101.101.101.101",
	"101.102.103.104",
	"https://dns.twnic.tw/dns-query",
	"203.80.96.10",
	"203.80.96.9",
	"218.102.23.228",
	"203.198.7.66",
	"61.10.0.130",
	"61.10.1.130",
	"202.181.240.44",
	"210.0.255.251",
	"202.67.240.222",
	"202.153.97.2",
	"202.153.97.130",
	"202.14.67.4",
	"202.14.67.14",
	"202.76.4.1",
	"202.76.4.2",
	"202.177.2.2",
	"202.177.2.3",
	"202.85.128.32",
	"202.85.128.33",
	"203.194.239.32",
	"202.85.170.89",
	"168.95.192.1",
	"168.95.1.1",
	"112.121.178.187",
	"8.8.8.8",
	"8.8.4.4",
	"https://dns./dns-query",
	"9.9.9.9",
	"149.112.112.112",
	"https://dns.quad9.net/dns-query",
	"https://dns11.quad9.net/dns-query",
	"208.67.220.220",
	"https://doh.opendns.com/dns-query",
	"https://doh.familyshield.opendns.com/dns-query",
	"4.2.2.1",
	"4.2.2.2",
	"168.126.63.1",
	"168.126.63.2",
	"1.1.1.1",
	"1.0.0.1",
	"https://1.1.1.1/dns-query",
	"https://1.0.0.1/dns-query",
	"https://-dns.com/dns-query",
	"37.235.1.174",
	"37.235.1.177",
	"216.146.35.35",
	"216.146.36.36",
	"8.26.56.26",
	"8.20.247.20",
	"64.6.64.6",
	"64.6.65.6",
	"77.88.8.8",
	"77.88.8.1",
	"80.80.80.80",
	"80.80.81.81",
	"176.103.130.130",
	"176.103.130.131",
	"176.103.130.132",
	"176.103.130.134",
	"https://dns.adguard.com/dns-query",
	"https://dns-family.adguard.com/dns-query",
	"156.154.70.1",
	"156.154.71.1",
	"156.154.70.5",
	"156.154.71.5",
	"103.197.104.178",
	"103.197.106.75",
	"203.189.136.148",
	"203.112.2.4",
	"84.200.69.80",
	"84.200.70.40",
	"195.46.39.39",
	"195.46.39.40",
	"109.69.8.51",
	"91.239.100.100",
	"89.233.43.71",
	"81.218.119.11",
	"209.88.198.133",
	"185.222.222.222",
	"185.184.222.222",
	"https://doh.dns.sb/dns-query",
	"185.222.222.222:853",
	"185.184.222.222:853",
	"74.82.42.42",
	"66.220.18.42",
	"104.236.210.29",
	"45.55.155.25",
	"185.228.168.9",
	"185.228.169.9",
	"185.228.168.10",
	"185.228.169.11",
	"185.228.168.168",
	"185.228.169.168",
	"https://doh.cleanbrowsing.org/doh/family-filter/",
	"https://doh.powerdns.org",
	"202.79.32.33",
	"202.79.32.34",
	"https://public.dns.iij.jp/dns-query",
	"https://doh-jp.blahdns.com/dns-query",
	"101.132.183.99",
	"47.98.124.222",
	"123.207.137.88",
	"115.159.220.214",
	"120.77.212.84",
	"101.236.28.23",
	"103.219.29.29",
	"180.97.235.30",
	"115.159.96.69",
	"123.206.21.48",
	"63.223.94.66",
	"115.159.146.99",
	"https://rubyfish.cn/dns-query",
	"https://dns.rubyfish.cn/dns-query",
	"tls://dns.rubyfish.cn",
	"113.205.16.215",
	"140.143.226.193",
	"123.207.22.79",
	"111.230.37.44",
	"110.43.41.122",
	"150.242.98.63",
	"https://i.233py.com/dns-query",
	"https://dns.233py.com/dns-query",
	"https://doh.rixcloud.dev/dns-query",
	"202.141.162.123",
	"202.38.93.153",
	"202.141.176.93",
	"101.6.6.6",
	"61.132.163.68",
	"202.102.213.68",
	"219.141.136.10",
	"219.141.140.10",
	"61.128.192.68",
	"61.128.128.68",
	"218.85.152.99",
	"218.85.157.99",
	"202.100.64.68",
	"61.178.0.93",
	"202.96.128.86",
	"202.96.128.166",
	"202.96.134.33",
	"202.96.128.68",
	"202.103.225.68",
	"202.103.224.68",
	"202.98.192.67",
	"202.98.198.167",
	"222.88.88.88",
	"222.85.85.85",
	"219.147.198.230",
	"219.147.198.242",
	"202.103.24.68",
	"202.103.0.68",
	"222.246.129.80",
	"59.51.78.211",
	"218.2.2.2",
	"218.4.4.4",
	"61.147.37.1",
	"218.2.135.1",
	"202.101.224.69",
	"202.101.226.68",
	"219.148.162.31",
	"222.74.39.50",
	"219.146.1.66",
	"219.147.1.66",
	"218.30.19.40",
	"61.134.1.4",
	"202.96.209.133",
	"116.228.111.118",
	"202.96.209.5",
	"108.168.255.118",
	"61.139.2.69",
	"218.6.200.139",
	"219.150.32.132",
	"219.146.0.132",
	"222.172.200.68",
	"61.166.150.123",
	"202.101.172.35",
	"61.153.177.196",
	"61.153.81.75",
	"60.191.244.5",
	"123.123.123.123",
	"123.123.123.124",
	"202.106.0.20",
	"202.106.195.68",
	"221.5.203.98",
	"221.7.92.98",
	"210.21.196.6",
	"221.5.88.88",
	"202.99.160.68",
	"202.99.166.4",
	"202.102.224.68",
	"202.102.227.68",
	"202.97.224.69",
	"202.97.224.68",
	"202.98.0.68",
	"202.98.5.68",
	"221.6.4.66",
	"221.6.4.67",
	"202.99.224.68",
	"202.99.224.8",
	"202.102.128.68",
	"202.102.152.3",
	"202.102.134.68",
	"202.102.154.3",
	"202.99.192.66",
	"202.99.192.68",
	"221.11.1.67",
	"221.11.1.68",
	"210.22.70.3",
	"210.22.84.3",
	"119.6.6.6",
	"124.161.87.155",
	"202.99.104.68",
	"202.99.96.68",
	"221.12.1.227",
	"221.12.33.227",
	"202.96.69.38",
	"202.96.64.68",
	"221.131.143.69",
	"112.4.0.55",
	"211.138.180.2",
	"211.138.180.3",
	"218.201.96.130",
	"211.137.191.26",
	"223.87.238.22",
	"183.230.98.97",
	"183.230.127.17",
	"2400:3200::1",
	"2400:3200:baba::1",
	"https://2400:3200::1/dns-query",
	"https://2400:3200:baba::1/dns-query",
	"2400:da00::6666",
	"tls://dns.ipv6dns.com",
	"https://dns.ipv6dns.com/dns-query",
	"2001:cc0:2fff:1::6666",
	"2001:cc0:2fff:2::6",
	"2001:de4::101",
	"2001:de4::102",
	"2001:19f0:7001:46b8:5400:01ff:feac:13a7",
	"2408:8262:12bd:1f22::2333",
	"2408:8262:12bd:1f22::3332",
	"tls://v6.rubyfish.cn",
	"https://v6.rubyfish.cn/dns-query",
	"2001:4860:4860::8888",
	"2001:4860:4860::8844",
	"2001:dc7:1000::1",
	"240C::6666",
	"240C::6644",
	"https://[2001:4860:4860::64]/dns-query",
	"https://[2001:4860:4860::6464]/dns-query",
	"2606:4700:4700::1111",
	"2606:4700:4700::1001",
	"https://[2606:4700:4700::1111]/dns-query",
	"https://[2606:4700:4700::1001]/dns-query",
	"https://[2606:4700:4700::64]/dns-query",
	"https://[2606:4700:4700::6464]/dns-query",
	"2a00:5a60::ad1:0ff",
	"2a00:5a60::ad2:0ff",
	"2a00:5a60::bad1:0ff",
	"2001:1608:10:25::1c04:b12f",
	"2001:1608:10:25::9249:d69b",
	"2a00:1508:0:4::9",
	"2620:119:35::35",
	"2620:119:53::53",
	"2610:a1:1018::1",
	"2610:a1:1019::1",
	"2610:a1:1018::5",
	"2620:fe::fe",
	"2620:fe::9",
	"2620:74:1b::1:1",
	"2620:74:1c::2:2",
	"2001:67c:28a4::",
	"2a01:3a0:53:53::",
	"2001:418:3ff::53",
	"2001:418:3ff::1:53",
	"2a09::",
	"2a09::1",
	"2a0d:2a00:1::2",
	"2a0d:2a00:2::2",
	"2a0d:2a00:1::1",
	"2a0d:2a00:2::1",
	"2a0d:2a00:1::",
	"2a0d:2a00:2::",
	"2001:da8::666",
	"2001:da8:8000:1:202:120:2:101",
	"2001:da8:202:10::36",
	"2001:da8:202:10::37",
	"2001:da8:208:10::6",
	"240e:4c:4008::1",
	"240e:4c:4808::1",
	"240e:56:4000:8000::69",
	"240e:56:4000::218",
	"2408:8663::2",
	"2408:8662::2",
	"2408:8000::8",
	"2408:8888::8",
	"2409:8062:2000:1::1",
	"2409:8062:2000:1::2",
	"2409:8028:2000::1111",
	"2409:8028:2000::2222",
	"tls://dns.rubyfish.cn:853",
	"tls://1.0.0.1:853",
	"tls://dns.google:853",
	"tls://dns.alidns.com:853",
	"tls://dns.cfiec.net:853",
	"tls://dns.ipv6dns.com:853",
	"tls://dns.pub:853",
	"tls://doh.pub:853",
	"https://dns.google/dns-query",
	"https://cloudflare-dns.com/dns-query",
	"[2a09::]:853",
	"[2a09::1]:853",
	"94.140.14.14",
	"94.140.15.15",
	"2a10:50c0::ad1:ff",
	"2a10:50c0::ad2:ff",
	"tls://2.dnscrypt.default.ns1.adguard.com",
	"176.103.130.130:5443",
	"tls://2.dnscrypt.default.ns2.adguard.com",
	"[2a00:5a60::ad2:0ff]:5443",
	"tls://dns.adguard.com",
	"94.140.14.15",
	"94.140.15.16",
	"2a10:50c0::bad1:ff",
	"2a10:50c0::bad2:ff",
	"tls://2.dnscrypt.family.ns1.adguard.com",
	"176.103.130.132:5443",
	"tls://2.dnscrypt.family.ns2.adguard.com",
	"[2a00:5a60::bad2:0ff]:5443",
	"tls://dns-family.adguard.com",
	"94.140.14.140",
	"94.140.14.141",
	"2a10:50c0::1:ff",
	"2a10:50c0::2:ff",
	"tls://2.dnscrypt.unfiltered.ns1.adguard.com",
	"94.140.14.140:5443",
	"[2a00:5a60::01:ff]:5443",
	"https://dns-unfiltered.adguard.com/dns-query",
	"tls://dns-unfiltered.adguard.com",
	"2a02:6b8::feed:0ff",
	"2a02:6b8:0:1::feed:0ff",
	"tls://2.dnscrypt-cert.browser.yandex.net",
	"77.88.8.78:15353",
	"77.88.8.88",
	"77.88.8.2",
	"2a02:6b8::feed:bad",
	"2a02:6b8:0:1::feed:bad",
	"77.88.8.3",
	"77.88.8.7",
	"2a02:6b8::feed:a11",
	"2a02:6b8:0:1::feed:a11",
	"tls://cleanbrowsing.org",
	"185.228.168.168:8443",
	"[2a0d:2a00:1::]:8443",
	"tls://family-filter-dns.cleanbrowsing.org",
	"185.228.168.10:8443",
	"[2a0d:2a00:1::1]:8443",
	"https://doh.cleanbrowsing.org/doh/adult-filter/",
	"tls://adult-filter-dns.cleanbrowsing.org",
	"https://doh.cleanbrowsing.org/doh/security-filter/",
	"tls://security-filter-dns.cleanbrowsing.org",
	"tls://2.dnscrypt-cert.shield-2.dnsbycomodo.com",
	"8.20.247.2",
	"2610:a1:1019::5",
	"156.154.70.2",
	"156.154.71.2",
	"2610:a1:1018::2",
	"2610:a1:1019::2",
	"156.154.70.3",
	"156.154.71.3",
	"2610:a1:1018::3",
	"2610:a1:1019::3",
	"156.154.70.4",
	"156.154.71.4",
	"2610:a1:1018::4",
	"2610:a1:1019::4",
	"tls://2.dnscrypt-cert.opendns.com",
	"2620:0:ccc::2",
	"208.67.222.123",
	"208.67.220.123",
	"tls://dns.google",
	"https://dns.cloudflare.com/dns-query",
	"tls://1.1.1.1",
	"1.1.1.2",
	"1.0.0.2",
	"2606:4700:4700::1112",
	"2606:4700:4700::1002",
	"https://security.cloudflare-dns.com/dns-query",
	"tls://security.cloudflare-dns.com",
	"1.1.1.3",
	"1.0.0.3",
	"2606:4700:4700::1113",
	"2606:4700:4700::1003",
	"https://family.cloudflare-dns.com/dns-query",
	"tls://family.cloudflare-dns.com",
	"2620:fe::fe:9",
	"tls://2.dnscrypt-cert.quad9.net",
	"9.9.9.9:8443",
	"[2620:fe::fe]:8443",
	"tls://dns.quad9.net",
	"9.9.9.10",
	"149.112.112.10",
	"2620:fe::10",
	"2620:fe::fe:10",
	"9.9.9.10:8443",
	"[2620:fe::fe:10]:8443",
	"https://dns10.quad9.net/dns-query",
	"tls://dns10.quad9.net",
	"9.9.9.11",
	"149.112.112.11",
	"2620:fe::11",
	"2620:fe::fe:11",
	"9.9.9.11:8443",
	"[2620:fe::11]:8443",
	"tls://dns11.quad9.net",
	"tls://dns.switch.ch",
	"130.59.31.248",
	"2001:620:0:ff::2",
	"https://dns.switch.ch/dns-query",
	"193.58.251.251",
	"92.38.152.163",
	"93.115.24.204",
	"2a03:90c0:56::1a5",
	"2a02:7b40:5eb0:e95d::1",
	"tls://2.dnscrypt-cert.dns.comss.one",
	"94.176.233.93:443",
	"[2a02:7b40:5eb0:e95d::1]:443",
	"https://dns.comss.one/dns-query",
	"tls://dns.comss.one",
	"92.223.109.31",
	"91.230.211.67",
	"2a03:90c0:b5::1a",
	"2a04:2fc0:39::47",
	"https://dns.east.comss.one/dns-query",
	"tls://dns.east.comss.one",
	"149.112.121.10",
	"149.112.122.10",
	"2620:10A:80BB::10",
	"2620:10A:80BC::10",
	"https://private.canadianshield.cira.ca/dns-query",
	"tls://private.canadianshield.cira.ca",
	"149.112.121.20",
	"149.112.122.20",
	"2620:10A:80BB::20",
	"2620:10A:80BC::20",
	"https://protected.canadianshield.cira.ca/dns-query",
	"tls://protected.canadianshield.cira.ca",
	"149.112.121.30",
	"149.112.122.30",
	"2620:10A:80BB::30",
	"2620:10A:80BC::30",
	"https://family.canadianshield.cira.ca/dns-query",
	"185.121.177.177",
	"169.239.202.202",
	"2a05:dfc7:5::53",
	"2a05:dfc7:5353::53",
	"tls://dot-fi.blahdns.com",
	"95.216.212.177",
	"https://doh-fi.blahdns.com/dns-query",
	"tls://2.dnscrypt-cert.blahdns.com",
	"95.216.212.177:8443",
	"2a01:4f9:c010:43ce::1:8443",
	"tls://dot-jp.blahdns.com",
	"139.162.112.47",
	"139.162.112.47:8443",
	"[2400:8902::f03c:92ff:fe27:344b]:8443",
	"tls://dot-de.blahdns.com",
	"159.69.198.101",
	"https://doh-de.blahdns.com/dns-query",
	"159.69.198.101:8443",
	"2a01:4f8:1c1c:6b4b::1:8443",
	"https://fi.doh.dns.snopyta.org/dns-query",
	"95.216.24.230",
	"2a01:4f9:2a:1919::9301",
	"tls://fi.dot.dns.snopyta.org",
	"https://dns-doh.dnsforfamily.com/dns-query",
	"tls://dns-dot.dnsforfamily.com",
	"94.130.180.225",
	"78.47.64.161",
	"2a01:4f8:1c0c:40db::1",
	"2a01:4f8:1c17:4df8::1",
	"tls://dnsforfamily.com",
	"193.17.47.1",
	"185.43.135.1",
	"2001:148f:ffff::1",
	"2001:148f:fffe::1",
	"https://odvr.nic.cz/doh",
	"tls://odvr.nic.cz",
	"180.131.144.144",
	"180.131.145.145",
	"tls://2.dnscrypt-cert.nawala.id",
	"tls://asia.dnscepat.id",
	"172.105.216.54",
	"2400:8902::f03c:92ff:fe09:48cc",
	"https://asia.dnscepat.id/dns-query",
	"tls://eropa.dnscepat.id",
	"5.2.75.231",
	"2a04:52c0:101:98d::",
	"https://eropa.dnscepat.id/dns-query",
	"tls://dot.360.cn",
	"tls://public.dns.iij.jp",
	"https://dns.pub/dns-query",
	"tls://dot.pub",
	"tls://101.101.101.101",
	"174.138.21.128",
	"2400:6180:0:d0::5f6e:4001",
	"tls://2.dnscrypt-cert.dns.tiar.app",
	"https://doh.tiarap.org/dns-query",
	"https://doh.tiar.app/dns-query",
	"quic://doh.tiar.app",
	"tls://dot.tiar.app",
	"172.104.93.80",
	"2400:8902::f03c:91ff:feda:c514",
	"tls://2.dnscrypt-cert.jp.tiar.app",
	"[2400:8902::f03c:91ff:feda:c514]:1443",
	"https://jp.tiarap.org/dns-query",
	"https://jp.tiar.app/dns-query",
	"tls://jp.tiar.app",
	"172.104.237.57",
	"172.104.49.100",
	"51.38.83.141",
	"2001:41d0:801:2000::d64",
	"tls://2.dnscrypt-cert.oszx.co",
	"51.38.83.141:5353",
	"[2001:41d0:801:2000::d64]:5353",
	"https://dns.oszx.co/dns-query",
	"tls://dns.oszx.co",
	"51.38.82.198",
	"2001:41d0:801:2000::1b28",
	"tls://2.dnscrypt-cert.pumplex.com",
	"51.38.82.198:5353",
	"[2001:41d0:801:2000::1b28]:5353",
	"https://dns.pumplex.com/dns-query",
	"tls://dns.pumplex.com",
	"94.130.106.88",
	"2a01:4f8:c0c:83ed::1",
	"https://doh.applied-privacy.net/dns-query",
	"tls://dot1.applied-privacy.net",
	"54.174.40.213",
	"52.3.100.184",
	"104.155.237.225",
	"104.197.28.121",
	"tls://2.dnscrypt-cert.safesurfer.co.nz",
	"176.9.199.158",
	"2a01:4f8:151:11b0::3",
	"tls://2.dnscrypt-cert.DeCloudUs-test",
	"176.9.199.158:8443",
	"[2a01:4f8:151:11b0::3]:8443",
	"https://dns.decloudus.com/dns-query",
	"tls://dns.decloudus.com",
	"51.158.147.50",
	"2001:bc8:2db9:100::853",
	"https://resolver-eu.lelux.fi/dns-query",
	"tls://resolver-eu.lelux.fi",
	"tls://2.dnscrypt-cert.captnemo.in",
	"139.59.48.222:4434",
	"tls://185.222.222.222",
	"176.9.93.198",
	"176.9.1.117",
	"2a01:4f8:151:34aa::198",
	"2a01:4f8:141:316d::117",
	"https://dnsforge.de/dns-query",
	"tls://dnsforge.de",
	"https://kaitain.restena.lu/dns-query",
	"158.64.1.29",
	"2001:a18:1::29",
	"tls://kaitain.restena.lu",
	"tls://2.dnscrypt-cert.dnsrec.meo.ws",
	"185.121.177.177:5353",
	"169.239.202.202:5353",
	"tls://dot.ffmuc.net",
	"https://doh.ffmuc.net/dns-query",
	"tls://2.dnscrypt-cert.ffmuc.net",
	"5.1.66.255:8443",
	"[2001:678:e68:f000::]:8443",
	"https://dns.digitale-gesellschaft.ch/dns-query",
	"185.95.218.42",
	"2a05:fc84::42",
	"tls://dns.digitale-gesellschaft.ch",
	"185.95.218.43",
	"2a05:fc84::43",
	"88.198.92.222",
	"https://doh.libredns.gr/dns-query",
	"https://doh.libredns.gr/ads",
	"tls://dot.libredns.gr.com",
	"116.202.176.26",
	"tls://ibksturm.synology.me",
	"83.77.85.7",
	"https://ibksturm.synology.me/dns-query",
	"178.82.102.190",
	"tls://2.dnscrypt-cert.ibksturm",
	"83.77.85.7:8443",
	"[2a02:1205:5055:de60:b26e:bfff:fe1d:e19b]:8443",
	"tls://getdnsapi.net",
	"185.49.141.37",
	"2a04:b900:0:100::37",
	"tls://dnsovertls.sinodun.com",
	"145.100.185.15",
	"2001:610:1:40ba:145:100:185:15",
	"tls://dnsovertls1.sinodun.com",
	"145.100.185.16",
	"2001:610:1:40ba:145:100:185:16",
	"tls://unicast.censurfridns.dk",
	"2a01:3a0:53:53::0",
	"tls://anycast.censurfridns.dk",
	"tls://dns.cmrg.net",
	"199.58.81.218",
	"2001:470:1c:76d::53",
	"tls://dns.larsdebruin.net",
	"51.15.70.167",
	"tls://dns-tls.bitwiseshift.net",
	"81.187.221.24",
	"2001:8b0:24:24::24",
	"tls://ns1.dnsprivacy.at",
	"94.130.110.185",
	"2a01:4f8:c0c:3c03::2",
	"tls://ns2.dnsprivacy.at",
	"94.130.110.178",
	"2a01:4f8:c0c:3bfc::2",
	"tls://dns.bitgeek.in",
	"139.59.51.46",
	"tls://dns.neutopia.org",
	"89.234.186.112",
	"2a00:5884:8209::2",
	"tls://privacydns.go6lab.si",
	"2001:67c:27e4::35",
	"tls://dot.securedns.eu",
	"146.185.167.43",
	"2a03:b0c0:0:1010::e9a:3001",
	"tls://dnsotls.lab.nic.cl",
	"200.1.123.46",
	"2001:1398:1:0:200:1:123:46",
	"tls://tls-dns-u.odvr.dns-oarc.net",
	"184.105.193.78",
	"2620:ff:c000:0:1::64:25",
	"88.198.91.187",
	"2a01:4f8:1c0c:8233::1",
	"https://doh.centraleu.pi-dns.com/dns-query",
	"tls://dot.centraleu.pi-dns.com",
	"95.216.181.228",
	"2a01:4f9:c01f:4::abcd",
	"https://doh.northeu.pi-dns.com/dns-query",
	"tls://dot.northeu.pi-dns.com",
	"45.67.219.208",
	"2a04:bdc7:100:70::abcd",
	"https://doh.westus.pi-dns.com/dns-query",
	"tls://dot.westus.pi-dns.com",
	"185.213.26.187",
	"2a0d:5600:33:3::abcd",
	"https://doh.eastus.pi-dns.com/dns-query",
	"tls://dot.eastus.pi-dns.com",
	"45.63.30.163",
	"2001:19f0:5801:b7c::1",
	"https://doh.eastau.pi-dns.com/dns-query",
	"tls://dot.eastau.pi-dns.com",
	"66.42.33.135",
	"2001:19f0:7001:225d::1",
	"https://doh.eastas.pi-dns.com/dns-query",
	"tls://dot.eastas.pi-dns.com",
	"https://doh.pi-dns.com/dns-query",
	"45.76.113.31",
	"tls://2.dnscrypt-cert.dns.seby.io",
	"tls://dot.seby.io",
	"139.99.222.72",
	"https://doh-2.seby.io/dns-query",
	"185.235.81.1",
	"185.235.81.2",
	"2a0d:4d00:81::1",
	"2a0d:4d00:81::2",
	"https://doh.dnslify.com/dns-query",
	"tls://doh.dnslify.com",
	"185.235.81.3",
	"185.235.81.4",
	"2a0d:4d00:81::3",
	"2a0d:4d00:81::4",
	"185.235.81.5",
	"185.235.81.6",
	"2a0d:4d00:81::5",
	"2a0d:4d00:81::6",
	"tls://dns.nextdns.io",
	"https://basic.bravedns.com/",
	"76.76.2.0",
	"https://freedns.controld.com/p0",
	"tls://p0.freedns.controld.com",
	"76.76.2.1",
	"https://freedns.controld.com/p1",
	"tls://p1.freedns.controld.com",
	"76.76.2.2",
	"https://freedns.controld.com/p2",
	"tls://p2.freedns.controld.com",
	"76.76.2.3",
	"https://freedns.controld.com/p3",
	"tls://p3.freedns.controld.com",
	"https://doh.mullvad.net/dns-query",
	"tls://doh.mullvad.net",
	"https://adblock.doh.mullvad.net/dns-query",
	"tls://adblock.doh.mullvad.net",
	"tls://2.dnscrypt-cert.dns.arapurayil.com",
	"3.7.156.128:8443",
	"https://dns.arapurayil.com/dns-query",
}

var dnsJsonApiList = pie.Strings{
	"https://223.5.5.5/resolve",
	"https://223.6.6.6/resolve",
	"https://dns.alidns.com/resolve",
	"https://doh.pub/resolve",
	"https://doh.360.cn/resolve",
	"https://dns.cfiec.net/resolve",
	"https://dns.twnic.tw/resolve",
	"https://dns./resolve",
	"https://dns.quad9.net/resolve",
	"https://dns11.quad9.net/resolve",
	"https://doh.opendns.com/resolve",
	"https://doh.familyshield.opendns.com/resolve",
	"https://1.1.1.1/resolve",
	"https://1.0.0.1/resolve",
	"https://-dns.com/resolve",
	"https://dns.adguard.com/resolve",
	"https://dns-family.adguard.com/resolve",
	"https://doh.dns.sb/resolve",
	"https://public.dns.iij.jp/resolve",
	"https://doh-jp.blahdns.com/resolve",
	"https://rubyfish.cn/resolve",
	"https://dns.rubyfish.cn/resolve",
	"https://i.233py.com/resolve",
	"https://dns.233py.com/resolve",
	"https://doh.rixcloud.dev/resolve",
	"https://2400:3200::1/resolve",
	"https://2400:3200:baba::1/resolve",
	"https://dns.ipv6dns.com/resolve",
	"https://v6.rubyfish.cn/resolve",
	"https://[2001:4860:4860::64]/resolve",
	"https://[2001:4860:4860::6464]/resolve",
	"https://[2606:4700:4700::1111]/resolve",
	"https://[2606:4700:4700::1001]/resolve",
	"https://[2606:4700:4700::64]/resolve",
	"https://[2606:4700:4700::6464]/resolve",
	"https://dns.google/resolve",
	"https://cloudflare-dns.com/resolve",
	"https://dns-unfiltered.adguard.com/resolve",
	"https://dns.cloudflare.com/resolve",
	"https://family.cloudflare-dns.com/resolve",
	"https://dns10.quad9.net/resolve",
	"https://dns.switch.ch/resolve",
	"https://dns.comss.one/resolve",
	"https://dns.east.comss.one/resolve",
	"https://private.canadianshield.cira.ca/resolve",
	"https://protected.canadianshield.cira.ca/resolve",
	"https://family.canadianshield.cira.ca/resolve",
	"https://doh-fi.blahdns.com/resolve",
	"https://doh-de.blahdns.com/resolve",
	"https://fi.doh.dns.snopyta.org/resolve",
	"https://dns-doh.dnsforfamily.com/resolve",
	"https://asia.dnscepat.id/resolve",
	"https://eropa.dnscepat.id/resolve",
	"https://dns.pub/resolve",
	"https://doh.tiarap.org/resolve",
	"https://doh.tiar.app/resolve",
	"https://jp.tiarap.org/resolve",
	"https://jp.tiar.app/resolve",
	"https://dns.oszx.co/resolve",
	"https://dns.pumplex.com/resolve",
	"https://doh.applied-privacy.net/resolve",
	"https://dns.decloudus.com/resolve",
	"https://resolver-eu.lelux.fi/resolve",
	"https://dnsforge.de/resolve",
	"https://kaitain.restena.lu/resolve",
	"https://doh.ffmuc.net/resolve",
	"https://dns.digitale-gesellschaft.ch/resolve",
	"https://doh.libredns.gr/resolve",
	"https://ibksturm.synology.me/resolve",
	"https://doh.centraleu.pi-dns.com/resolve",
	"https://doh.northeu.pi-dns.com/resolve",
	"https://doh.westus.pi-dns.com/resolve",
	"https://doh.eastus.pi-dns.com/resolve",
	"https://doh.eastau.pi-dns.com/resolve",
	"https://doh.eastas.pi-dns.com/resolve",
	"https://doh.pi-dns.com/resolve",
	"https://doh-2.seby.io/resolve",
	"https://doh.dnslify.com/resolve",
	"https://doh.mullvad.net/resolve",
	"https://adblock.doh.mullvad.net/resolve",
	"https://dns.arapurayil.com/resolve",
}
