# Backend

This is the corresponding backend for the Workadventure admin back office.

## Configuration
The configuration is only loaded from the environment variables.
When running inside docker the configuration is loaded through docker compose and passed to the container. See the root README of this repository.  
When you develop on your local machine it is recommended to create a `.env` file which is loaded during the [startup](main.go). For all possible configuration options see the [config.go](config/config.go).

## Deployment
This project was developed natively in docker and is intended to only run inside a docker deployment. All necessary scripts and steps are inside the [Dockerfile](Dockerfile).

## Development
It is recommended to use the Jetbrains Goland IDE for full support. Configure the environment variables using a `.env` file. Start the `main.go` and the Workadventure instance.  
In order to communicate from docker to the local machine development backend use `host.docker.internal` as FQDN with the configured port for the `ADMIN_API`.

### API
The API documentation is done inside the go files method headers using [`swaggo/swag`](https://github.com/swaggo/swag). See the `Getting started` section to install the utility.  
After changing the documentation first start run `swag fmt` and after that `swag init`. Restart the webserver when already running.  
The swagger documentation is available at [`http://localhost:4664/swagger`](http://localhost:4664/swagger). It is OpenAPI v3 compatible and can be used to generate a API client.

### Database
As a database a mariadb 10+ is required. Due to the ORM layer of ent go postgres can be supported when needed.  
To work with ent go the following commands are mandatory.

**Add Entities**
```bash
go run -mod=mod entgo.io/ent/cmd/ent new EntityName
```

**Generate Entities**
```bash
go generate ./ent
```

**Create Versioned migration**
```basg
go run -mod=mod ent/migrate/main.go migrationName
```

**Apply Migrations**  
When developing with a local and outside of docker, the migration isn't done automatically. Basically the following command is needed.
```bash
atlas migrate apply --dir "file://ent/migrate/migrations" --url mysql://root:pass@localhost:3306/ent
```

To also use the environment variables or the `backend/.env` file it is recommended to use the migration script which is also used in the docker entrypoint.
```bash
./migrate.sh
```