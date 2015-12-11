# Golang-MongoDB3.0-Example
Connect Go (Golang) app to MongoDB 3.0 example


I spend 3-4 hours in search to find how to connect Go (Golang) app to the new MongoDB 3.0 with its new SCRAM-SHA-1 Authentication Mechanism so I am adding here an example of what it worked for me.
Misleadingly I found around the network that has to be included the "labix.org/v2/mgo" package despite the fact that on the official site http://labix.org/mgo at the time of the reading we have newer and updated information that points to the right package "gopkg.in/mgo.v2" . I spend few hours trying to run the other package. That did not worked at all and when I realized what went wrong decided to upload simple app here just as working example. 
