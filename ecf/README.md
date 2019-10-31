# EXPERIMENTAL CHANGEFEED FOR

1. Run `movies.sql` on local CockroachDB server if necessary.

2. Run the collector:

    ```bash
    $> go run main.go
    ```

3. Do some DB changes (see `movies.sql` file) and wait for them to appear on stdout.