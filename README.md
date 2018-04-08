# Regexp-in-Golang

## Description

This is a regular expression engine implemented in Golang.
It follows Ken Thompson's NFA algorithm invented in the mid-1960.
https://swtch.com/~rsc/regexp/regexp1.html

Shunting yard algorithm used to convert infix into postfix notation
The metacharacters \\*, \\+, and \\? are repetition operators: e1* matches a sequence of zero or more (possibly different) strings, each of which match e1; e1+ matches one or more; e1? matches zero or one.



## Installation

All codes in this repository must be ran via go compiler 
1. First go to https://golang.org/dl/ to download a binary release suitable for your system
2. You can clone your repository to create a local copy on your computer and sync between the two locations 
3. On GitHub, navigate to the main page of the repository
4. Clone or download button Under the repository name, click Clone or download
5. Clone URL button In the Clone with HTTPs section, click to copy the clone URL for the repository. 
6. Open Git Bash. 
7. Change the current working directory to the location where you want the cloned directory to be made.
8. Type git clone, and then paste the URL you copied in Step 2.

git clone https://github.com/YOUR-USERNAME/YOUR-REPOSITORY 

9. Press Enter. Your local clone will be created.

## Run and Test
1.Go to the location you cloned and press regexpByGo file to open Visual Studio code 
2.Open terminal cd location you cloned  
3.First you will be asked to enter regular expression e.g, a.b.c ,every letter must be followed by .(dot) except the last one if quantifier follows a letter .(dot) must follow quantifier 
4.Second prompt will be asked to enter a string to check against the reg exp 
5.Next you can press 1 to continue or -1 to exit programm
after a letter follows 

## License

Educational 