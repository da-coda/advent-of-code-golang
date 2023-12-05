# Trebuchet the second
___
[If you do not care about more goroutine stuff and just want to see stuff about my solution click here](#wait-a-second-you-didnt-even-work-on-the-task-of-part-2-yet)

**A quick warning:** I will show some benchmark results in here. Take them with a grain of salt. I did run the benchmarks
just a few times and for the sake of sensationalism I choose the ones with the biggest improvements. This
is not a scientific paper!

I could just adapt the second part directly in my initial code. But where is the fun in that?
Stuff like AOC is, at least for me, all about learning things and testing stuff. So I decided to find the answer
to one of my questions in part one: Does the code benefit from Goroutines?

# Step 1: rewrite to a sequential approach
___
To test this I decided to first take the solution from Part 1 and to rewrite it to a sequential approach. So we 
are truly going line by line. Then I rewrote some parts so that we can easily throw the go benchmark at the two 
solutions. I just extracted everything from the main func into a run func which has a `path` parameter. When we run
the code via the main func, we just use os.Args[1] as the source of the path. And in the benchmark we can just call
the run func with a different path (because the context during go test is a different, so the path is different as well).

# Step 2: run the benchmark
___ 
This is pretty easy. Because Go is great (most of the time). One thing I like most about go is definitely the "batteries
included" approach and tooling. Just run `go test -v ./day_1_trebuchet/... -bench ^\QBenchmarkMain\E$ -run ^$` and both
benchmarks for part 1 and part 2 are run.

# Step 3: interpreting the result
___
This is probably the difficult part. I will first present the result of a benchmark run:
```bash
/home/daniel/go/go1.21.3/bin/go test -v ./day_1_trebuchet/... -bench ^\QBenchmarkMain\E$ -run ^$
goos: linux
goarch: amd64
pkg: advent-of-code-golang/day_1_trebuchet/part_1
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkMain
BenchmarkMain-8   	     523	   2168029 ns/op
PASS
ok  	advent-of-code-golang/day_1_trebuchet/part_1	1.370s
goos: linux
goarch: amd64
pkg: advent-of-code-golang/day_1_trebuchet/part_2
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkMain
BenchmarkMain-8   	     352	   3089529 ns/op
PASS
ok  	advent-of-code-golang/day_1_trebuchet/part_2	1.435s
```

The Goroutine approach took ~70% as long as the sequential approach (comparing just the ns/op).
That is without a doubt faster. But we are talking about 3.09ms vs 2.17ms. 
But there is one interesting thing that happened. I did run the benchmark more than once (as one should).
While the sequential approach had at most a difference of 0.15ms/op. At the same time the goroutine approach
had a difference from up to 0.3ms/op. This should probably be expected because of the overhead of starting goroutines and
scheduling.

I decided to run a second benchmark with a fixed 10.000 iterations:
```bash
/home/daniel/go/go1.21.3/bin/go test -v ./day_1_trebuchet/... -bench ^\QBenchmarkMain\E$ -run ^$ -benchtime=10000x
goos: linux
goarch: amd64
pkg: advent-of-code-golang/day_1_trebuchet/part_1
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkMain
BenchmarkMain-8   	   10000	   2407253 ns/op
PASS
ok  	advent-of-code-golang/day_1_trebuchet/part_1	24.079s
goos: linux
goarch: amd64
pkg: advent-of-code-golang/day_1_trebuchet/part_2
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkMain
BenchmarkMain-8   	   10000	   3229388 ns/op
PASS
ok  	advent-of-code-golang/day_1_trebuchet/part_2	32.302s
```

And again we see that goroutines are faster. But not by that much. I mean I don't know how many calculations the
elves have to run at a time, but I doubt that it is 10k or more.

# Take away
___
Yes, Goroutines are faster (who would have thought), but at the same time the code becomes worse to read and debug.
The performance impact is not big enough to justify the usage of goroutines.

# Just kidding
___
I am not done yet. Go routines come with an overhead. Especially the creation. So how can we circumvent this overhead?
**WORKER POOLS!**

Why should we create a new goroutine every time? We can just create a worker pool that is as big as we have available cores
and let them handle the calculations. And here are the results (just comparing the normal goroutine approach to worker pool):
```bash
/home/daniel/go/go1.21.3/bin/go test -v ./day_1_trebuchet/... -bench . -run ^$
goos: linux
goarch: amd64
pkg: advent-of-code-golang/day_1_trebuchet/part_1
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkRun
BenchmarkRun-8               	     535	   2244082 ns/op
BenchmarkRunAsWorkerPool
BenchmarkRunAsWorkerPool-8   	     843	   1403423 ns/op
PASS
ok  	advent-of-code-golang/day_1_trebuchet/part_1	2.761s
```

The pool approach takes just ~63% of the time (ns/op) compared to the simple goroutine approach and
~43% of the sequential approach. Again, this is a total overkill and absolutely unnecessary. But it is still amazing to
see! And it is a great learning.

Just for fun I did run it again with 10k iterations:
- Sequential: 30.389s
- Goroutine: 22,236s
- Worker pool: 15,623s

The code for the worker pool is in part 1.

# Wait a second, you didn't even work on the task of part 2 yet!
___
That is right. As so often when one starts to play around everything else becomes more interesting than the actual
task. But now I will work on the actual task. I will choose the sequential approach just because it is more readable.

# No Regex for us
___ 
Well in the end it did bite me in the ass. For the second task we are not just looking for digits in each line but also 
the string form of numbers (one, two, three, etc.). So we cant just throw out all non-digit parts of each line.
Well maybe (probably) there is a way, but we all know that regex can be a pain in the ass and should be used as a last 
resort (Also I found this site which I thought was amazing https://regexlicensing.org/).

# The state of the machine
___
I was always fascinated by state machines. I learned about them in university and I really liked the concept.
In the end Regex is also just a state machine. So I decided to build my own little state machine.

The idea was simple:
- Iterate through a line
- tell the state machine to transition its state according to the current character
- check if the state machine is in an end state
- get the value for the end state and put it as the first value if there is none yet and/or as the current last value
- once we iterated through the string we have a start and an end value (which can be the same)
- startValue * 10 + endValue gives us the result for this line

# The internals of the state machine
___
The idea is pretty simple:
On every transition we are checking the current state which can transition two either the next state or back to an initial
state. Here is one example:
```go
        if unicode.IsDigit(character) {
            sm.isFinishedState = true
            sm.finishedStateValue, _ = strconv.Atoi(string(character))
            sm.currState = ""
            return
        }
	switch sm.currState {
	case "":
		sm.currState = string(character)
	case "o":
		switch character {
		case 'n':
			sm.currState = "on"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	case "on":
		switch character {
		case 'e':
			sm.finishedStateValue = 1
			sm.isFinishedState = true
			sm.currState = "e"
		default:
			sm.currState = sm.currState[1:]
			sm.TransitionState(character)
		}
	}


```
If we found a digit we can just set the end state directly.

If there is no current state we just take the current character as the state. If the current state is `o` the only valid 
state transition would be `n` which would lead to the state `on` which can then only transition to the next state with `e`.
With the `e` transition we end up in an end state. We can check for this state with the methode 
```go 
func (sm *StateMachine) IsFinished() (bool, int)
```
which tells us if we are in an end state and returns the numeric value.

There is also a way to reset the state machine to check for the next number
```go
func (sm *StateMachine) Reset()
```
But the currState is not reset in there. We set the currState to the last letter of the end state value.
This allows to successfully parse stuff like `oneight` and `sevenine`.

# Pit falls and obstacles
___

There was one problem I ran into which took me forever to fix. Why took it so long? Because of all 1k lines in the document
there were only 3 that are breaking. Let's see if you can find the problem (hint: In the code above there is already the
fix for the problem):
```
"kone1ptnkjhks65sixrsseight"
"9hfjjmgrzntssjpxcvbzpvmqzgsd54twonine"
"7kndzrhvcnstgfxjlff9twoninervrknsffmfzmdhtth"
```
How did I find those three? I took the solution from someone else and compared their values against mine. And those three
were flagged as different.

My state machine had problems when there was the beginning of one number which then transitioned into another number.
Let's take a look at the first line:

kone1ptnkjhks65sixrs`seight`

The state machine sees the *se* and expects a `v` as the next transition. There is no `v` so it goes back into an initial 
state and fails to take the `eight` into account.

In the end the fix was not that difficult. In every `fail state` we remove the first character and test the remaining state
recursively. If we have removed all characters we end up in our `case "":` block and everything is fine.

# Final thoughts
___
While part 1 was rather easy, part 2 made stuff a lot more complicated. And after reading some threads about Day 1 online
I seem to not be the only one that thought so. 
And having a bug that only affects 3 out of 1000 sucks so much. I had at one point 50 test cases which I took from my document
and calculated by hand.

But in the end it worked out. I am looking forward to Day 2.