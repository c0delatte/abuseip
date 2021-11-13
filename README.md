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

### Disclaimer
For my own learning purpose. These codes are messy af. Feel free to contribute so I know how to code properly.
