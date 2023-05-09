# Open Source Workadventure admin back office

This project is developed as an open source alternative for the Premium Workadventure back office.
It was a requirement to have administration functionality during the "MathAdventure" project at the Hochschule Bochum – Bochum University of Applied Sciences.

This project consist out of two components:
## Backend
The main component for the Admin API as well for the configuration. Authentication is done against this part of the project.
The Workadventure Pusher and Backend is in communication with this part of the back office to get details of the player.

Currently implemented features are:
- IP/UUID bans with expiration
- Map/Room configuration
- User reporting
- Room access control
- Global announcements
- User tags
- Woka/Companion textures with tags

## Frontend
The frontend is in communication with the back office backend. It is basically a UI for the database operations.
It also serves the login page for local authentication or redirects to the SSO provider.

## Features
- 🌌 Modern infrastructure
- 🐳 Infrastructures lives natively in docker
- 🏃 Fast backend thanks to the nature of Go
- 📝 Well documented
- 💡 Modern and responsive UI
- 🔐 Well tested in a production environment
- 👤 Local and *SAML* authentication
- 🔑 Token Login
- 🔥 Support latest version of Workadventure
- ✒️ GDPR compliant

## Testing
A plug-and-play docker configuration is available in [`wadev/`](wadev/).
You can start a complete Workadventure instance with the admin back office by using the following command.

```bash
docker compose up -f compose.yml compose.backoffice.yml
```

You can access your local Workadventure instance at [`play.workadventure.localhost`](http://play.workadventure.localhost).