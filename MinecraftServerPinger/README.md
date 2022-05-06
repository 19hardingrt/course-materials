This program works by using the minecraft server query setting that is built-in to minecraft servers.
	However, the query setting must be enabled in order for the program to be able to query the server.

	The program will output relevant server information such as:
		- host port
		- host IP address
		- gametype
		- MOTD - server description
		- the game ID
		- server version
		- number of players
		- world name
		- max players
		- plugins


	To use the program:
		- If testing on a local server, ensure that the query setting is enabled in the server.properties file for the minecraft server
		(querying is set to false by default. Refer to minecraft's documentation on how to set up a minecraft server if need be).

		- The program will query a minecraft server that is running a query port on 25565
		(should be set to 25565 by default in the server.properties file, if you need to change the query port that the program looks at,
		CTRL + F "25565" in the code and change it to desired port, is located in the main function.)

		- build the program (go build main.go)
		- run the program (either ./main or go run main.go)
		- enter an IP Address to the terminal
		(If server is running locally, set IP address to "localhost" in server.properties file and then enter
		localhost into the terminal.)

		- The program will timeout after 15 seconds regardless whether or not it successfully queries the server.
		(If the program times out without displaying any server information, the server is either down, doesn't have querying enabled, or the
		defined query port does not match what is in the server.properties.)
