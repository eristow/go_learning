## TODO FF

- [ ] Use file uploaders for album art in edit/create pages


## TODO BE

- [ ] Refactor existing tests

- [ ] Add tests for everything new

- [ ] Convert to using classes instead of static methods in files?


## TODO DEPLOY

- Keep Docker Compose as dev, and K8s as prod

- [ ] Push code to GitHub as secondary remote

- [ ] Create K8s config containers

- [ ] Create Terraform config for K8s config

- [ ] Create GitLab CI/CD
  - [ ] FE tests
  - [ ] BE tests
  - [ ] Docker test creation
  - [ ] Deploy


## DONE
- [x] Dockerize FE
- [x] Co-locate `+page.server.ts` actions to a single location using named actions
- [x] Use env var for backend URL
- [x] Add CRUD functions for albums
- [x] Create PUT route for updates
- [x] Create simple FE
- [x] Dockerize BE
- [x] Change to using auto IDs for albums
- [x] Use PSQL DB for BE instead of in-memory data


## Frontend libraries
- prettier
- eslint
- vitest
- tailwindcss
	- forms
- sveltekit-adapter
	- node
- drizzle
  - PSQL
