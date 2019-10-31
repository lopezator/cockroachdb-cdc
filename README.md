# CockroachDB Change Data Capture (CDC)

This repository contains examples of use for the two available [CockroachDB CDC](https://www.cockroachlabs.com/docs/stable/change-data-capture.html) options available:

- [EXPERIMENTAL CHANGEFEED FOR](https://www.cockroachlabs.com/docs/stable/changefeed-for.html) (free version) in `ecf/` folder.
- [CREATE CHANGEFEED FOR](https://www.cockroachlabs.com/docs/stable/create-changefeed.html) (enterprise version) in `ccf/` folder.

## EXPERIMENTAL CHANGEFEED FOR

#### Advantages:

- Prints out directly to stdout from a sqlconn, you don't need any specific sink or http server with everything that entails.

#### Drawbacks:

- Doesn't manage cursor state, you have to materialize it on persistent storage and manage logic to inject it on runtime.

## CREATE CHANGEFEED

#### Advantages:

- Manages cursor state for you. Everytime a request comes into a configurable sink successfully (ack), it updates the `high_water_timestamp` in order to avoid returning already returned results.
- Configurable sinks allow you to get the results directly in the system you want (if available in the sinks list).

#### Drawbacks:

- As there isn't already a GCP PubSub sink available (and we don't want to use Apache Kafka), we need to enable middleware in the form of an http endpoint or a Google Cloud Bucket Storage. 