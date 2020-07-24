[![Build Status](https://dev.azure.com/achoudhary3/zeronet/_apis/build/status/redestro.zeronet?branchName=master)](https://dev.azure.com/achoudhary3/zeronet/_build/latest?definitionId=5&branchName=master)

# Zeronet
> The future of AI Games

## Introduction

Zeronet is an AI-powered tic-tac-toe game that presents four modes to the user to cover every possible way to play the game. The game is also enriched with features such as move recommendation, exclusive AI vs AI game mode, and a sleek UX. The AI is powered by the minimax algorithm and the architecture surrounding is developed using high-end OOP practices.

## The Application Of Object Oriented Programming

> Foreword: Since Go does not have the concept of classes in-built, but instead uses structs to emulate that behavior, so we have also decided to follow the same route.

### Session

The basic entity of our game is a Session. Whenever a user starts a game, a new session is created as per the game mode and this session is stored in our in-memory database.
A session has the following fields:

* Token(string)
* Moves(Array(9)int)
* Board(Array(9)string)
* MoveCount(int)
* Player1(*Player)
* Player2(*Player)

The session has its own interface functions:

* Create: Creates a new session instance depending on the parameters provided.
* Update: Updates the Game(Board and Moves array) status in the session. The two arrays(Board and Moves) will be referred to combined as Game from here on.

### Player

The secondary entity of our architecture is a Player. A player can be of two types: Human or AI. A player struct has the following fields:

* token(string) -> same as the session
* Kind(string)
* symbol(string)

The player struct has its own interface functions:

* Init: Creates a new Player instance depending on the parameters provided.
* Play: The interface call invokes our AI algorithm and returns the move played by the AI. The same interface is also used for procuring recommendations for the next player.

## Architecture Flow

### Start Session

The first step is to start the session. There are 4 kinds of sessions that are supported:

* Human vs AI (Human is Player1)
* Human vs Human
* AI vs Human (AI is Player1)
* AI vs AI

Once an appropriate call is made to the API, a random token is generated and then a corresponding token is created and add to the in-memory database.
The session creation involves creating appropriate players with the correct Kind and symbol and also initializing the Game(Board and Moves arrays).
Play Game
The play API is simple and the only requirement is the token and the current player position(whether he is player1 or player2) and the move he played. The rest of the computation is made using the information already stored within the session.
The Play function is called via the interface which returns the next move played by the AI and also returns the recommendation by switching the players and updating the board.
There is an additional API to fetch the first move by the API in scenarios where AI is the first player.


### End Session
Whenever the game ends(either in a win, loss, or draw), a delete API is called to remove the session from the memory and the in-memory database.

## Frontend Flow
The first step involves choosing the appropriate game mode, upon which the relevant API is called to create the session. The token received is committed to the store.
The frontend contains 2 basic entities:

* Board
* Cell

Whenever a move is made on the cell, it emits an event that informs the board to update the state and calls the play API with the relevant payload to fetch the next move/recommendation.
In a scenario where the game ends the delete API is called to remove the session and free up the memory and also resets the store.

The A-team:
[@AyanChoudhary](https://github.com/AyanChoudhary)
[@shubhamgupta2956](https://github.com/shubhamgupta2956)