# WEB SHOP FACULTY PROJECT

---

This is a web shop application developed as a faculty project together with my colleague **Amar Omerović**.  
The project represents a medium-level full-stack web shop with standard e-commerce features.

## Features
- User registration and login
- Product listing and categories
- Shopping cart management
- Order processing
- Admin panel for managing products, categories and orders
- Translation
- ....

---

## Tech Stack used for this Backend Implementation

- Golang (Gin Framework)
- PostgreSQL
- Docker
- Docker compose
- Hetzner Object Storage Bucket (German Hosting Company "Hetzner", like AWS Bucket, used for storing images)

Requirement to run both of these on any PC locally (localhost) is to have installed Docker min "v28.x" version and Docker compose min "v2.x"

To start the docker compose script (which is starting Golang Backend and Database PostgreSQL), you first need to be at the root level of this repository.
After that open the terminal from that root folder and type the command

```bash
docker compose up
```

After that, keep your terminal window open, do not terminate the action, since it will stop the whole backend.
If you wish to use your terminal after starting this project, type the command

```bash
docker compose up -d
```

When you put the **-d** in the command, that means it will run in **detach** mode. After starting successfully it will go back to normal state in terminal (the backend) 
will run in the background.
