# CREATE CHANGEFEED + CLOUDEVENTS + PUBSUB TRANSPORT

The example consists of two folders:

## `sender/`

- Reveives the notification from CRDB CDC onto an http server.
- Generates the cloudevent notification from CRDB CDC received JSON.
- Sends the cloudevent notification over pubsub transport.

## `receiver/`

- Receives the cloudevent notification over pubsub transport.
- Prints it out to stdout.