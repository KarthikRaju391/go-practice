## What's the purpose of the `done` channel?

Well, `goroutines` can go out of control pretty easily because it's very easy to create one in the first place.
To prevent this, the `done` channel pattern can be used. The most basic implementation of the `done` channel pattern can be found in the `done_channel.go` file.

Now, what the `done` channel lets us do is to kill the child goroutines from the parent. If you observe in the code, our `main` function is the parent function. And the `doWork` function is the child function.

We are creating a `done` channel in the `main` function which is just a channel of boolean. Now we are passing it down to the `doWork` function as a `readonly` argument so that the child process may not write to the channel.

With the access to the `done` channel, it's as simple as adding a `case` to the `select` statement which has a `default` case defined.

Also notice how there is no need to consume a message from the channel specifically. It's a feature, not a bug.
