

<img height="250" src="https://github.com/kswope/restjester/blob/master/assets/jester.png" />

# restjester BETA

## REST API mocking server


* You don't like mocking API REST calls, you want the real thing, sorta.
* Single golang binary to just download and run, no installation.
* Same solution for all languages.
* Because you have to use an external API if you are testing in a browser.
* Can run as system service or command line.


### Install

Download the server compiled for your architecture

[Windows](https://github.com/kswope/restjester/blob/sync/releases/windows/amd64/restjester?raw=true)
|
[OSX](https://github.com/kswope/restjester/blob/master/server/bin/restjester?raw=true)
|
[Linux](https://github.com/kswope/restjester/blob/master/server/bin/restjester?raw=true)
|
[ARM](https://github.com/kswope/restjester/blob/master/server/bin/restjester?raw=true)
|
[ARM64](https://github.com/kswope/restjester/blob/master/server/bin/restjester?raw=true)



Run with output to terminal
```
shell> ./restjester
Starting server at port 5351
```

Run as a daemon (requires deamon)
```
shell> deamon --name restjester ./restjester
```

### If you want/need to compile your own binary 

Install [golang](https://golang.org/), download this repo, run 'make', the binary will be in server/bin/


### ruby example using rest-client and rspec

```
it "can install and GET resource" do

  # install resource on restjester
  RestClient.post 'localhost:5351', { path:'/users/1', data: {username: 'kswope'}.to_json }

  # GET resource
  response = RestClient.get 'localhost:5351/users/1' 
  expect( JSON.parse( response.body ) ).to eql( { 'username'=>'kswope' } )

end
```

### Installing endpoint ruby examples
```
RestClient.post 'localhost:5351', { method:'GET'     path:'/users/1', data: {username: 'kswope'}.to_json }
RestClient.post 'localhost:5351', { method:'PUT',    path:'/users/1', status:200 }
RestClient.post 'localhost:5351', { method:'POST',   path:'/users/1', status:200 }
RestClient.post 'localhost:5351', { method:'DELETE', path:'/users/1', status:403 }
```

### Install endpoint parameters
* path ( required )
* method ( optional, default is GET )
* status ( optional, default is 200 )
* data ( optional, string (probably json), body of response, default is "" )


### GET all endpoints ruby example
```
RestClient.get 'localhost:5351'
```

### Query string parameter order doesn't matter
```
RestClient.post 'localhost:5351', { method:'GET' path:'/users/1?a=1&b=2', data: {username: 'kswope'}.to_json }

response = RestClient.get 'localhost:5351/users/1?a=1&b=2' 

# same as
response = RestClient.get 'localhost:5351/users/1?b=2&a=1' 
```


### Viewing all endpoints in browser

Just to go [localhost:5351](http://localhost:5351)

Tip: install a JSON viewer plugin like [JSONView](https://chrome.google.com/webstore/detail/jsonview/chklaanhfefbnpoihckbnefhakgolnmc)

### Clearing all installed endpoints ruby example
```
RestClient.delete 'localhost:5351'
```

### javascript example
```
TODO
```

### perl example
```
TODO LATER
```

### INSTALL:

Not sure of platform?  
> uname --hardware-platform

Run with output to terminal
> ./restjester

Run as a daemon (requires deamon)
> deamon --name restjester /usr/sbin/restjester

Run on system startup, copy init script to appropriate place


