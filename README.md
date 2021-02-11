# revel-test


DOCS: https://revel.github.io/manual/index.html
github link github.com/habruzzo/revel-test/

alright heres how i set up this env:
i already have go installed on my Mac (if you dont, do `brew install go`)
i read the docs for go and they say (paraphrase) basically if you set the env variable "GOPATH" to the directory where i want the go project to be, and say "go get ___" itll get a copy of what i want and put in the directory i asked for.

I opened a new repo on github, and then i cloned it into /Users/holdenabruzzo/proj/revel-test
then i did:
`export GOPATH=~/proj/revel-test` << setting GOPATH so that "go get" will work
`go get github.com/revel/revel` << getting and installing the "revel" git repo
`go get github.com/revel/cmd/revel` << getting and installing the revel command line tool git repo
`export PATH=$PATH:$GOPATH/bin` << adding the revel command line tool to the env path so 

`revel new -a gotestapp` << creating a new directory tree that will be built and run in
`revel run -a gotestapp` << starting the local dev webserver to serve the application
`curl localhost:9000` << just checking that the app works. it should pull down a html page. you also can just go to a browser and go to localhost:9000 and see the page

i also attempted a "go.mod" but i think i fucked it up so dont rely on it or try to use it as is