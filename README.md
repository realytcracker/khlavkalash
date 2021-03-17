# khlavkalash
this is an http server that serves only one thing, no matter what you ask for.

it gets its name from the simpsons - https://www.youtube.com/watch?v=4NFv5IGP2uA.

crab juice not included.

no pizza, only khlav kalash.

## installation
`go get -u https://github.com/realytcracker/khlavkalash/...`


## usage
```
 __   ___  __   ___  
|/"| /  ")|/"| /  ") 
(: |/   / (: |/   /  
|    __/  |    __/   [khlavkalash] by [ytcracker]
(// _  \  (// _  \   
|: | \  \ |: | \  \  
(__|  \__)(__|  \__) 
  -f string
        path to your khlav kalash
  -h    prints this helpful garbage
  -l string
        if set, 301 redirect to <parameter>, skip -f
  -m string
        mime type of your khlav kalash (default "image/jpeg")
  -p uint
        port to serve khlav kalash on (80 may require root privs) (default 80)
  -s string
        http server version header (default "nginx/1.17.10")
  -v    print incoming requests to stdout
```
