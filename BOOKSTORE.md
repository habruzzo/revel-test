This section has the details for the BookStore application. 

models are probably going to end up being structs or something but they ought to be separate from controllers and views (need to figure this out, it will be all Go though)

views themselves dont really require logic to make, they are templates mostly (very little Go needed)

controllers(specifically the actions) are where the routing logic is gonna be (this is all Go)

    actions we will probably need:
        save = saves books after user has edited them
        edit = allows user to edit book entry in db
        display = basically a load function that takes the stuff out of db and shows it all nice
        payment = abstracted, but basically a function that will check inventory, if we have enough of the item, use abstracted module (aka mock/stub this part) to validate customer credit card and take the payment, remove book from inventory and "deliver" it to customer (trigger some sort of log or notification or email or something, mock it etc)

    other things each action will need:
        input validation/sanitization
        logging