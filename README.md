# Life-Inventory-API


This repository manages the storage and handling of personal entries — such as tasks, routines, reminders, and appointments — grouped by categories.

## Current structure of an Entry
- **Name**  
- **Date**  
- **Status** (done or not done)  

## Features
- Retrieve the list of entries  
- Retrieve individual entries  
- Save new entries  
- Update existing entries  
- Delete entries  

## API Documentation
Swagger UI is available here (the project must be running):  
👉 [http://localhost:8080/docu/index.html](http://localhost:8080/docu/index.html)

## Planned features
- As my app development progresses (link to app repo will be added here), I will continue to extend the functionality.  
- Likely additions:
  - More details per entry  
  - Categories (e.g. routines, habits, reminders, tasks, events)  
  - Repeating entries (daily, weekly, etc.)  

## How to run the project

1. Clone the repository  
    ```bash
    git clone <repository-url>
    cd <project-folder>
    ```

2. Start services  
    ```bash
    docker-compose up
    ```

3. Run the project  
    ```bash
    go run .
    ```

## Requirements
- Go must be installed  
  👉 [Download and install Go](https://go.dev/doc/install)  
  (verify with `go version`)  

- Docker must be installed  
  👉 [Download and install Docker](https://docs.docker.com/get-docker/)  
  (verify with `docker --version`)  

- Optional (to install project dependencies manually)  
    ```bash
    go get .
    ```

## License
[Apache License](LICENSE)