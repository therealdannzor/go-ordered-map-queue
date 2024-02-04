# Ordered Map Queue with RabbitMQ
This is an example of an ordered map queue implementation with goroutines, unbuffered channels, and parallel execution.

### Assumptions
* Instructions are fed into the client application through a CSV text file (see below for examples)
* The solution is optimized for consistency with mutexes surrounding each Update and Delete operation
* It is a prototype and to demonstrate a simple use-case with RabbitMQ, with no advanced redundancy
* Exhaustive list of systems tested on: Apple M1 Pro 14 in. (2021)

### Text File API
* `Update(k, v)`: `ADD,key,value` adds a key and value
* `Delete(k)`: `DELETE,key` deletes a key and value
* `Get(k)`: `LOOKUP,k` gets the key and value, if it exists
* `GetAll`: `GET_ALL` gets all key and value pairs and prints to file

# Quickstart

Make sure you have Docker installed in order to run the message queue:
```
make mq
```

Build the client and server binaries:
```
make
```

Run the unit tests:
```
make test
```

# In Action

Make sure you have built the client and server binaries first.

```
# In terminal 1, start the queue
make mq

# In terminal 2, start the server
./bin/server

# In terminal 3, add 300 kv pairs
./bin/client < input-add-1.txt

# In terminal 4 directly, add an additional 330 kv pairs
./bin/client < input-add-2.text

# Get the status of all kv pairs stored
./bin/client < input-get-all-1.text

# Open the file which should have been produced and check content with editor of choice.
# The name of the file starts with the current timestamp `2024-*-*...`

# Look up a value
./bin/client < input-lookup-1.txt

# Delete the first 430 kv pairs
./bin/client < input-delete-1.txt

# Get the status as above and check that it has been removed
```
