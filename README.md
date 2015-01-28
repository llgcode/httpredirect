[![Build Status](https://drone.io/github.com/llgcode/httpredirect/status.png)](https://drone.io/github.com/llgcode/httpredirect/latest)
# httpredirect
Http Server written in Golang that make simple redirection. 

it support ssl, vhosts, simple serving files, http redirection

# Installation

from source, first install [golang](http://golang.org/doc/install) and then execute this command
```
go install github.com/llgcode/httpredirect
```

# Run

```
$ httpredirect.exe --help
Usage of httpredirect.exe:
  -log="": Define log file
  -router="router.json": Define router file
  -sslCert="": Path to certificate file
  -sslKey="": Path to key file
```

# router.json
```json
{
 "Port": 80,
 "Redirections": [
   {
       "Path": "songbook.llgmusic.net/",
       "URL": "http://127.0.0.1:8081/"
   }, 
   {
       "Path": "www.llgmusic.net/",
       "URL": "http://127.0.0.1:8082/"
   }, 
   {
       "Path": "llgmusic.net/",
       "URL": "http://127.0.0.1:8082/"
   },
   {
       "Path": "/mywebapp/",
       "URL": "/opt/mywebapp"
   }
]
}
```
