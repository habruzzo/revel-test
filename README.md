# revel-test
Sorry about the formatting in github, if you read it as a normal text file it has better formatting.

The purpose of this tutorial is to teach about Go and basic web applications. In this tutorial we will be creating a web application with several features. It will use the Revel application framework (see DOCS below for info on Revel).
It will have a simple tictactoe page, and a more complex bookstore application.

DOCS: https://revel.github.io/manual/index.html
github link: github.com/habruzzo/revel-test/

Here's how I set up the dev environment:

    1. I already have go installed on my Mac (if you dont, do `brew install go`)
    2. I read the docs for go and they say (paraphrase) basically if you set the env variable "GOPATH" to the directory where you want the go project to be, and say "go get ___" itll get a copy of what you want and put in the directory you asked for.
    3. I made a new repo /Users/holdenabruzzo/proj/revel-test. you can clone it at the github link above.
    4. Then I did:
        `export GOPATH=~/proj/revel-test` << setting GOPATH so that "go get" will work
        `go get github.com/revel/revel` << getting and installing the "revel" git repo
        `go get github.com/revel/cmd/revel` << getting and installing the revel command line tool git repo
        `export PATH=$PATH:$GOPATH/bin` << adding the revel command line tool to the env path so you can use the command locally

        `revel new -a gotestapp` << creating a new directory tree that will be built and run in
        `revel run -a gotestapp` << starting the local dev webserver to serve the application
        `curl localhost:9000` << just checking that the app works. it should pull down a html page. you also can just go to a browser and go to localhost:9000 and see the page

    i also attempted a "go.mod" but i think i fucked it up so dont rely on it or try to use it. ill figure that out later.

Important concepts of web applications:
    MVC (what are alternatives?)
    "front-end" ("client")
    "back-end" ("server")
    security
    network
    context
    static content
    cookies/sessions
    db(sql vs nosql)

Tic Tac Toe?
    Features:
        tic tac toe board that updates with every "move" the user makes
    Use case:
        user goes to localhost:9000/tictactoe
        user sees a blank tic tac toe board
        user clicks a tile on the board to make a move
        the page refreshes, this time with a responding move by the computer
        the user responds until someone wins or the game is tied
        the screen announces the game is over and who won, if anyone

    TODO:
        - make view (done)
        - register router
        - create controller
        (front end only, basically)

    See architecture and implementation notes in TICTACTOE.md

Book Storefront?
    Features:
        book = title(string), author(string), publish date(date), date added(date), price(double), quantity(int)
        1. search for books based on title/author
        2. store books in database
        3. add books to database
        4. remove books from database
        5. display books for sale
        6. handle payment transaction
    Use cases:
        the user comes to the website localhost:9000/books
        the user is able to see the books that are available for sale
        the user is able to click on a book to navigate to the "edit" page
        here, the user can edit the book information or purchase the book
        if the user edits the book information, they have the option to save their work
            choosing to save will refresh the page and return them to the display screen
        if the user chooses to purchase the book, they will go to a payment page
            the payment page will take their card information and return a confirmation or error, and a link to return them to the display screen

    TODO:
        - make views (display view, payment view, edit view(add/delete))
        - register routers (display, edit, save, payment)
        - create controllers (display, edit, save, payment)
        - connect database (install revel db module --it exists i checked)

    1) Possible extension = admin access for book editing? (install revel auth module --it exists i checked)
    2) Other possible extension = using a performant deployment server (caddy server https://caddyserver.com/docs/ should be simple enough)
    3) Other possible extension = csrf tokens (basic anti-mitm, install revel csrf module --it exists i checked)
    4) Other possible extension = a search functionality (search view, search controller, etc)

    See architecture and implementation notes in BOOKSTORE.md


What is MVC? https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller
M = model
V = view
C = controller

It's a way of splitting up your application, and it's pretty popular though not the only way (and an argument could be made that it's not the best way) but definitely a simple and effective enough way, so we might as well use it for learning. Plus, it's commonly adopted in enterprise situations because of its longevity in the industry, so you may see it at companies in the future still.

The model holds the data and deals with it.
The view presents the data to the user.
The controller directs traffic depending on how a user interacts with the view.

A basic problem with this is that models tend to get very bloated, and that will probably happen in this case as well. It's okay to flaunt the standard as long as you know the standard and you have a reason for doing what you're doing.