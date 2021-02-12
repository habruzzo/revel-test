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
    `export PATH=$PATH:$GOPATH/bin` << adding the revel command line tool to the env path so you can use the command locally

    `revel new -a gotestapp` << creating a new directory tree that will be built and run in
    `revel run -a gotestapp` << starting the local dev webserver to serve the application
    `curl localhost:9000` << just checking that the app works. it should pull down a html page. you also can just go to a browser and go to localhost:9000 and see the page

i also attempted a "go.mod" but i think i fucked it up so dont rely on it or try to use it as is.

important concepts
    MVC (what are alternatives?)
    front-end
    back-end
    security
    network
    context
    static files
    cookies/sessions
    db(sql vs nosql)

tic tac toe?

    TODO:
        -make view
        -register router
        -create controller
        (front end only, basically)

book storefront?
    features:
        book = title(string), author(string), publish date(date), date added(date), price(double), quantity(int)
        1.search for books based on title/author
        2.store books in database
        3.add books to database
        4.remove books from database
        5.display books for sale
        6.handle payment transaction

    TODO:
        -make views (display view, search view, payment view, edit view(add/delete))
        -register routers (display, search, edit, save, payment)
        -create controllers (display, search, edit, save, payment)
        -connect database (install revel db module --it exists i checked)

possible extension= admin access for book editing? (install revel auth module --it exists i checked)
other possible extension= using a real server (caddy server https://caddyserver.com/docs/ should be simple enough)
other possible extension im less interested in=csrf tokens (basic anti-mitm, install revel auth module --it exists i checked)

views themselves dont really require logic to make, templates mostly
controllers(specifically the actions) are where the real logic is gonna be

    actions we will probably need:
        save = saves books after user has edited them
        edit = allows user to edit book entry in db
        display = basically a load function that takes the stuff out of db and shows it all nice
        search = search book entries in db and display results
        payment = abstracted, but basically a function that will check inventory, if we have enough of the item, use abstracted module (aka mock/stub this part) to validate customer credit card and take the payment, remove book from inventory and "deliver" it to customer (trigger some sort of log or notification or email or something, mock it etc)

    other things each action will need:
        input validation/sanitization
        logging


