# CREATE CHANGEFEED

This folder contains an example of use for the [CockroachDB CDC](https://www.cockroachlabs.com/docs/stable/change-data-capture.html) enterprise version, [CREATE CHANGEFEED](https://www.cockroachlabs.com/docs/stable/create-changefeed.html).

There is an additional example of use, along with [GCP PubSub](https://cloud.google.com/pubsub/) and [CloudEvents](https://cloudevents.io/), located in the subfolder `ce-ps/`.

## Usage instructions

1. Run `movies.sql` on remote CockroachDB server if necessary.

2. Remote port forward to enable Cloud CockroachDB into local http server:

    ```
    $> ssh -R dlopez.serveo.net:80:localhost:8080 serveo.net
    ```

3. Run the http server:

    ```bash
    $> go run main.go
    ```

4. Create a `changefeed` watching the movies tables:

    ```cockroachdb
    CREATE CHANGEFEED FOR TABLE movies INTO 'experimental-http://dlopez.serveo.net/cdc' WITH UPDATED;
    ```

5. Do some DB changes (see `movies.sql` file) and wait for them to appear on stdout.