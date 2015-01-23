# httpredirect
Http Server written in Golang that make simple redirection. 

it support ssl, vhosts, simple serving files, http redirection
 
Sample config file

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
