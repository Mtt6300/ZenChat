# ZenChat
simple chat service with Go 

- you can create a lot of Chatroom and server can accept multi client.

- easy to share your room just need sent config file and Put it on the same path with client.exe and room.exe 

- simple config

# Install go 
visit go website and install it on your Server			

- link : https://golang.org/doc/install#install

# Customization Config
after installing go download source from github and edit ```config.json``` file:			
```
{
    "mainServer": {
        "host": "",
        "port": ":2323",
        "enabled": true
    },
    "gpServer": {
        "GroupName":"Localhost",
        "Bio":"Pls dont SPAM and\n be Friendly with us :D",
        "Admin":"Admin",
        "host": "",
        "port": ":3232",
        "info":"this is server info",
        "enabled": true
    }
}
```
important : mainsServer and gpServer host shoud be same and Their port should not be similar !!

# Build Client and Room
after changes, compile sources		
```go build client.go && go build room.go```		

you can build ```server.go``` but it's not recommended.

# Share room
Whenever you want to shere your room just send ```room.exe``` and ```client.exe``` and ```config.json``` to your friends 

# usage				

first you need to run server.go on server  ```go run server.go```

if you open ```room.exe``` you can watch chat room 		

if you want type chats, you can open ```client.exe``` and chat with your Friends
