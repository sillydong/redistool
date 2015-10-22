# RedisTool

As redis-cli can not do some commands like delete by pattern, I wrote this tool to help do such things. Features will be added one by one later.

## Features:

* ### show keys match pattern and delete them

	#### Usage:

	    NAME:
	       ./redistool keys - show all keys matching pattern
	    
	    USAGE:
	       ./redistool keys [command options] [arguments...]
	    
	    OPTIONS:
	       -h "127.0.0.1"       Server hostname (default: 127.0.0.1).
	       -p "6379"            Server port (default: 6379).
	       -a                   Password to use when connecting to the server.
	       -n "0"               Database number.
	       -r "<pattern>"       *               matches all
	                            h?llo           matches hello, hallo and hxllo
	                            h*llo           matches hllo and heeeello
	                            h[ae]llo        matches hello and hallo, but not hillo
	                            h[^e]llo        matches hallo, hbllo, ... but not hello
	                            h[a-b]llo       matches hallo and hbllo
	       -d                   delete keys matching pattern

	#### Example:
	
		#list matched keys
		./redistool keys -h "127.0.0.1" -p 6379 -a "qwer" -n 1 -r "hello*"
		
		#delete matched keys
		./redistool keys -h "127.0.0.1" -p 6379 -a "qwer" -n 1 -r "hello*" -d
