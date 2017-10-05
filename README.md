# data-rep-go-problems

## Intro
My name is Ray Mannion and I am a 3rd student at the Galway/Mayo Institute of Technology studying Software Development

In this repository you wiil find problems I was asked to complete for a data representation and query module.
The [problems](https://data-representation.github.io/problems/go-fundamentals.html) are done using the Go programming language, which I am currently learning.

In order for you to run these programs you will need to install the [GO](https://www.google.ie/?gws_rd=cr&dcr=0&ei=SQvUWejfHOaXgAaL3JeoBA)
compiler on your computer.

Once you have installed the compiler, and have the files on your computer, open a command line of your choice. Next, navigate your way to the folder. Then run the command "go build" followed by the name of the go file(including the file extension) you wish to run. This will then create an executeable file of the same name, but with an ".exe" ententsion. Then simpily run the .exe file and the GO program will run. 

![alt tag](https://github.com/rayman51/data-rep-go-problems/blob/master/images/goCap.PNG)

## Problems
The Ten programs are
1. Hello World
2. Time Date
3. FizzBuzz
4. Factorial
5. Guessing game
6. High Low
7. Palindrome
8. Merge Sort 
9. Newton Square Root
10. ReverseString

## Useful Links

1. https://golang.org

2. https://en.wikipedia.org/wiki/Go_(programming_language)

3. https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65

## Troubleshooting
I had trouble using the Scanf in some of these programs.
In the guessing game program the loop would trigger twice without user input.
After a quick google search I found out there is a small bug when using Scanf.
For me this problem was solved by adding a \n to the input line e.g fmt.Scanf("%d\n", &guess).
Alternatively, the problem can be solved using fmt.Scan(&guess) or Scanln(&guess). But I found it varies for device.




