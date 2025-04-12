## TODO FE

- [ ] Use file uploaders for album art in edit/create pages


## TODO BE



## TODO DEPLOY

- [ ] Create Terraform config for deploy on AWS
  - Using ECS Fargate instead of EKS to save money
  - Will need to adjust PSQL DB URL

- [ ] Break up `main.tf` into modules for each AWS section
  - RDS, ECS, networking, etc.

- [ ] Create a Terraform config with EKS for learning

- [x] Create GitLab CI/CD
  - [x] FE tests
  - [x] BE tests
  - [x] Docker test creation
  - [ ] Deploy


## DONE
- [x] Add tests
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
