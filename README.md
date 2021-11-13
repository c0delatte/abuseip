# abuseip
Check the report history of any IP address to see if anyone else has reported malicious activities using API by [abuseipdb.com](https://www.abuseipdb.com/).

### Installation
```
$ go get github.com/abaykan/abuseip@v1.0
```
- Put your `API_KEY` in `~/.config/ipdb.txt`:
```
$ echo "<YOUR_API_KEY>" > ~/.config/ipdb.txt
```

### Usage
Usage example:
```
$ echo "230.246.11.2" | abuseip
```
or
```
$ cat iplist.txt | abuseip > logs.txt
```
or
```
$ awk '{print $1}' /var/log/apache2/access.log | abuseip 
```

### Output
```
$ sudo awk '{print $1}' /var/log/apache2/access.log | abuseip 
[172.69.170.123] Code: US | Score: 0 | Total Reports: 0
[172.70.130.250] Code: US | Score: 0 | Total Reports: 1
[172.69.71.109] Code: US | Score: 20 | Total Reports: 3
[108.162.246.40] Code: US | Score: 19 | Total Reports: 6
[209.141.56.209] Code: US | Score: 100 | Total Reports: 2077
[172.69.69.124] Code: US | Score: 0 | Total Reports: 1
[205.185.124.100] Code: US | Score: 100 | Total Reports: 1163
[172.69.69.124] Code: US | Score: 0 | Total Reports: 1
[104.206.128.62] Code: US | Score: 100 | Total Reports: 892
[172.69.170.123] Code: US | Score: 0 | Total Reports: 0
[172.69.170.103] Code: US | Score: 0 | Total Reports: 0
[128.14.209.162] Code: US | Score: 100 | Total Reports: 1134
[172.70.35.40] Code: US | Score: 0 | Total Reports: 1
```
### Disclaimer
For my own learning purpose. These codes are messy af. Feel free to contribute so I know how to code properly.
