Setup instructions:

1) install postgres on system
    for MAC users you can follow this link to download the executable: 
        https://github.com/PostgresApp/PostgresApp/releases/download/v2.5.6/Postgres-2.5.6-10-11-12-13-14.dmg

2) Command to run migrations:

    - you should be in the root directory of the project, which is /bd_test_task_three
    - replace <database_name> with your DB name, remove angle brackets too. I set it to 'postgres' by default
    migrate -source file://internal/pkg/db/migrations/mysql -database "postgres://postgres:postgres@localhost:5432/<database_name>?sslmode=disable" up

3) change the name of the DB in schema.resolver.go file to whatever DB you selected (check the comments for the query to run on graphql)
4) in the root directory run the command go run ./server.go and go to localhost:8080 from browser

