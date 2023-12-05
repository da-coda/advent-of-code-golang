# Cubes and Colo(u)rs
___
This one felt a lot easier than the first one. There is not much to say.

# Refactoring and finding common functionality
___
For this task we again have to read in a file line by line and do stuff. And I got this feeling that we will do this for 
most of the other tasks as well. So I decided to extract this logic into a common function that takes a path and a function
as arguments. The path obviously tells the function which file to parse. The function will be run for every line.
This works greatly in Go because we got first-class functions.

In our common function we call the provided function and provide the current line:
```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
	f(scanner.Text())
}
```
For the calling side all we have to do is provide a function that deals with the line:
```go
var gameList games
common.ReadFile(os.Args[1], func(s string) {
	gameList = append(gameList, parseGame(s))
})
```
The great think here as you can see is that we have access to things defined outside our function. The same think in PHP 
would need to be declared explicitly:
```
$gameList = [];
Common::readFile($path, static function(string $line) use ($gameList) {...});
```

# Trivial but ugly

I prefer mostly simple readable code over some clever code. There is probably some very effective and elegant way to parse
a game line within a single line of code. But who will be able to read this? I would argue that most developers are just skimming
over code and try to get a grasp of what happens. If you stuff a lot of stuff in a single, clever line, one might miss this
and get confused. 

But to be honest, I hate doing stuff on strings. It always feels like I am committing a cardinal sin. Chopping up strings,
replacing stuff here and there, gluing things together again. I feel like Dr. Frankenstein.

In the end all we have to do is parsing a game line into some prepared structs, filter them according to some rules and
reduce them to the sum of their IDs. 

# Final thoughts
___
This task felt a lot more trivial compared to the first task. Especially the second part. The code is not as clean but it 
is readable. But if I have learned something from the first task: The second part is probably more difficult!