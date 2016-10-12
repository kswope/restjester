

<img height="250" src="https://github.com/kswope/restjester/blob/master/assets/jester.png" />

# restjester BETA

## REST API mocking server


* You don't like mocking API REST calls, you want the real thing, sorta.
* Single golang binary to just download and run, no installation.
* Same solution for all languages.
* Because you have to use an external API if you are testing in a browser.
* Can run as system service or command line.
* Easy to setup API backend for frontend development - load up restjester with data and go to work.


* [Windows](https://github.com/kswope/restjester/blob/master/releases/windows/amd64/restjester?raw=true)
* [OSX](https://github.com/kswope/restjester/blob/master/server/bin/restjester?raw=true)
* [Linux](https://github.com/kswope/restjester/blob/master/server/bin/restjester?raw=true)
* [ARM64](https://github.com/kswope/restjester/blob/master/server/bin/restjester?raw=true)
* [ARM](https://github.com/kswope/restjester/blob/master/server/bin/restjester?raw=true)


### ruby example using rest-client and rspec

```
shell> ./restjester
Starting server at port 5351
```

```
it "can install and GET resource" do

  # install resource on restjester
  RestClient.put 'localhost:5351', { path:'/users/1', data: {username: 'kswope'}.to_json }

  # GET resource
  response = RestClient.get 'localhost:5351/users/1' 
  expect( JSON.parse( response.body ) ).to eql( { 'username'=>'kswope' } )

end
```

Get root will view all endpoints
Delete root will clear all endpoints

### Installing endpoint ruby examples
```
RestClient.put 'localhost:5351', { path:'/users/1', data: {username: 'kswope'}.to_json }
RestClient.put 'localhost:5351', { method:'PUT', path:'/users/1', data: {username: 'kswope'}.to_json }
RestClient.put 'localhost:5351', { method:'DELETE', path:'/users/1', status:403 }
```


### Install endpoint parameters
* path ( required )
* method ( optional, default is GET )
* status ( optional, default is 200 )
* data ( optional, body of response )


### javascript example
```
TODO
```

### perl example
```
TODO LATER
```

### INSTALL:

restjester is currently distributed as a single compile golang binary.  No external dependancies.

Find the server compiled for your architecture and put on your path, like /usr/sbin
Not sure of platform?  
> uname --hardware-platform

Run with output to terminal
> ./restjester

Run as a daemon (requires deamon)
> deamon --name restjester /usr/sbin/restjester

Run on system startup, copy init script to appropriate place


