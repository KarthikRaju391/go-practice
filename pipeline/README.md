## Pipelines!

This one I had ~func~ fun with! There's a lot more going on in that program than what meets the eye.

Before I get to the explanation, let's understand the two different types of channels that can be setup in go:
1. Unbuffered channels
2. Buffered channels

### Unbuffered channels
These channels have no size defined which makes them synchronous. Here's how.
So, when we create a channel like this: `make(chan int)`, the channel will only be limited to have written to or read off from **one item at a time** and the process is `blocking` as well.
That means, if there is another **goroutine** reading off the same channel, it will be blocked until something is written to it and the writer will be blocked until the message is read.

### Buffered channels
These channels have a size defined for them which makes them asynchronous.
It's simple, since it has got some latitude to play with now, the writer can write to the channel and not give a func. The reader on the other hand can also read from the channel asynchronously.

We can create `buffered` channels like this: `make(chan int, 4)`. This means there can be **utmost** 4 values in the channel after which it starts `blocking` the process again.

---
Now, if we look at our program, you can immediately notice that in both the helper functions we are only creating `unbuffered` channels.
Despite us having created two different `goroutines` — one to convert the `slice` into a `channel` and another one to `square` the numbers, both the `goroutines` will be in sync because they are interdependent.
Until the item hasn't been added to the `numsChannel` which is the input for the `squareNumsChannel` function, it won't be processed, and once a number is added, it gets immediately squared and immediately gets logged as well in the `main` function because it's dependent on the `unbuffered` channel returned by the `squareNumsChannel` function.
Neat.

>NOTE: If you have gone through the `done_channel` example, you might have also noticed how we are returning the channels in `readonly` mode.
