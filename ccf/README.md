# CREATE CHANGEFEED

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