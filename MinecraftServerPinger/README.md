This program works by using the minecraft server query setting that is built-in to Java Edition minecraft servers. The querying functionality runs off of the UDP protocol.
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
		
		- NOTE THAT THIS ONLY WORKS ON SERVERS RUNNING ON THE JAVA EDITION OF MINECRAFT, BEDROCK EDITION SERVERS UTILIZE DIFFERENT PROTOCOLS.
		
	Breakdown of time spent working on project (since I was making changes/running tests entirely on my local machine instead of pushing all of my changes to 	  GitHub)
		- Planning: 1-2 Hours (Coming up with idea/pseudocoding)
		- Researching: 3-4 hours (Initial research/reading on querying documentation was about an hour, the remaining research time was spent looking at 		 specific documentation needed for code functionality such as the "time package" for implementing timeouts)
		- Getting baseline code written: ~13 hours (Not counting the pauses for reading code documentation)
		- Testing/Debugging: ~15 hours (This took a long time as initially my code was not outputting any server information when the server was running. Only 		       to find out this was because I did not have querying enabled on my local server *facepalm*
		- Total: 34-35 hours (Not counting my stupidity during the debug/testing process it took about 20 hours)
		
		
	External Resources:
		https://wiki.vg/Query (Documentation on the query protocol built into minecraft servers)
		https://pkg.go.dev/time (Documentation of time package, used for implementing timeouts for the program)
		https://gobyexample.com/timeouts (Specifically for implementing timeouts)
		https://www.w3schools.com/go/go_input.php (For accepting a user input rather than just pre-assigning a variable).
		https://help.minecraft.net/hc/en-us/articles/360058525452-How-to-Setup-a-Minecraft-Java-Edition-Server (Documentation for setting up a minecraft 		 server)
		
		
