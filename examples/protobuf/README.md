## Compile the proto

```sh
cd ..
protoc --go_out=. *.proto 
```

## Run the program

```sh
go run main.go
```

### O/p

```
Actual message:  name:"Vishwas" age:24 followers:{youtube:1400 twitter:2330}
Serialised message:  [10 7 86 105 115 104 119 97 115 16 24 26 6 8 248 10 16 154 18]
My name is Vishwas and my age is 24
I have 1400 followers in Youtube and 2330 in Twitter
```


