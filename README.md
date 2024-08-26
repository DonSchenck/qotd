# qotd
## "Quote of the Day" microservice written in Go.

This is used by various tutorials and articles about CodeReady Workspaces 2.0.

For more information, see the Github page.

### PostgreSQL database is optional  
Using a PostgreSQL database in your OpenShift project is optional; the code will automagically provide quotes if a database *is not* specified (or if retrieving data fails for any reason).

If you choose to use a database, note the following:

* OpenShift service name for the database must be put into the environment variable **DBNAME**.

* Database user name must be in the environment variable **DBUSER**.

* Database password must be in the secret supplied qotd-secrets.yaml.

* Database name MUST be '**quotesdb**'.

* Table name must be as follows:

`CREATE TABLE IF NOT EXISTS quotes (
    quoteId   int,
    quotation varchar( 256 ),
    author    varchar(  80 )
);`