# Pingu
Broadcast messages to all your web users

REST API to publish notifications to browsers via [Server Sent Events](https://developer.mozilla.org/en-US/docs/Server-sent_events) (lightweight one-way WebSockets). 

## Usage
###Creating a channel
Make a POST request to `/session` to obtain the channel ID and an auth token (used to post messages to the channel)

###Subscribing to the channel
Make a GET request to `/broadcast/channelID` to receive notifications on this channel via SSE

###Broadcasting a message to all subscribers
Make a POST request to `/broadcast/channelID` with token and message as parameters
