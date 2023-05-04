### Add Entities
`go run -mod=mod entgo.io/ent/cmd/ent new Entity`

### Generate Entities
`go generate ./ent`

### Create Versioned migration
`go run -mod=mod ent/migrate/main.go <name>`

### Apply Migrations
`atlas migrate apply --dir "file://ent/migrate/migrations" --url mysql://root:pass@localhost:3306/ent`

play.workadventure.localhost?token=<JWT_TOKEN>
```json
{
  "identifier": "EMAIL/UUID (ANON)",
  "accessToken": null,
  "username": "jwt-test",
}
```