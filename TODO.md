## TODO FE

- [ ] Add tests

- [ ] Use file uploaders for album art in edit/create pages


## TODO BE



## TODO DEPLOY

- Keep Docker Compose as dev, and K8s as prod

- [ ] Adjust K8s deploy for prod
  - Env vars and passwords should be more secure

- [ ] Create Terraform config for K8s config on AWS
  - Will need to adjust PSQL DB URL

- [ ] Create GitLab CI/CD
  - [ ] FE tests
  - [ ] BE tests
  - [ ] Docker test creation
  - [ ] Deploy


## DONE
- [x] Refactor existing tests
- [x] Add tests for everything new
- [x] Convert to using classes instead of static methods in files?
- [x] Edit README
- [x] `Cross-site POST form submissions are forbidden` in K8s deploy this time
  - Need to set `ORIGIN` env var in frontend to internal IP of pod
- [x] Create K8s config containers
- [x] Push code to GitHub as secondary remote
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
