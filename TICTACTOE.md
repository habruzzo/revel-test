This section has the details for the TicTacToe application.

The view for tictactoe is almost done, I just made a basic HTML table with 9 buttons to represent each square where someone can move. This application has a very thin client, and the server is going to be doing the heavy lifting on this page.

To implement this feature, first go to "TicTacToe.html" in revel-test/gotestapp/app/views/TicTacToe/Index.html
Look at the table. It's the tic tac toe board. Look at the Javascript. This is a very primitive UI. When the user clicks on a tile, it submits the form embedded in the page with the action "/move/<tile_no>". We will use this as a kind of API to log moves in the game. This will act as our "move endpoint", and we'll use it to update the game state.