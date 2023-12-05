# Trebuchet
___
One of the main selling points of Go is the concurrency pattern. And I am at least 90% sure that I just confused 
concurrency for parallelism. Or maybe not. This is definitely something that I should learn.

I have used WaitGroups, Channels and Goroutines before. But not very often and I have to check up the syntax for 
channels every time I do. And every time I do I find the syntax to be very intuitive. Your function needs a channel which
is read only? ```ch <-chan string```. You only want to write in a channel? ```ch chan<- string```. It makes sense! And I
still forget the syntax every time...

So when I saw that I have to read in a document and handle it line by line I knew what I had to do. 

# Reading in the document
___
Sure, I could just read in the whole file and iterate over it in go. But files can become large and I have to
calculate the calibration values line by line anyway. So instead I just open the file, read a line, throw this line into 
a channel and then read the next line. After difficult calculations and absolutely scientific decision-making I came up 
with a channel size of 20. (Add.: In the end I decided to use the runtime.NumCPU() function to decide how many workers
I will spawn)

# Calculating the calibration values
___
Cool, so now I have a channel filled with single lines from the base document. The calculation of the calibration
values is luckily pure. There is no side effect. All we have to do is take this string, find the numbers, put them 
together and return the number. This is a prime candidate for a goroutine. We can have many of those just calculate 
the value for a single line and throw the result again in another channel (I feel like I am building the next Venice).

Instead of traversing the string from both sides to find the first/last digit I just decided to regex away every character
so that we only have a string full of digits left. Now just take the first and last char and we are done.
This will surely not bite me in the ass later on in part 2, right?

# Summing up the values
___
Finally, all we have to do is sum up the values and print the resulting sum. For that we start a new goroutine that 
consumes the calibration value channel and adds the values up. Why do we do this in a goroutine? While this goroutine does
its addition magic, the main thread waits for the last few calculation goroutines to finish up and close the channel behind them.
The closing of the channel is the signal for our sum up goroutine to stop the consumption and to print the result.

# But this is overkill!
___
I know. And I am not really sure if this task even benefits from goroutines. 
But as I mentioned in the beginning I did not do it this way because I think that it is the best way,
but because I really wanted to play around with Goroutines and learn a bit more about them. And one main take away I got 
is that I need to get a better understanding about concurrency, parallelism and when to use them.

# Final thoughts
___ 
I really liked the task. Especially as the first task of the AOC. And I obviously fucked up my channels and goroutines more
than once. And debugging them is really not fun. But overall it took me ~30 min, and I am happy with the result. 