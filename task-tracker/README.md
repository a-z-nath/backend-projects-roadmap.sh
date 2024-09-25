# Task Tracker

Task tracker is a CLI(command line interface) application to manage your tasks.
this app is to handle some functionality of task such as `add`, `update`, `delete`.
it also handle tasks current status such as `todo`, `in-progress` and `done`.

## Requirements

- [x] add, update, and delete tasks
- [x] mark a task as in progress or done
- [x] list all tasks
- [x] list all tasks that are done
- [x] list all tasks that are not done
- [x] list all tasks that are in progress

## How To Run

Clone the repository and run the following command :

```bash
git clone https://github.com/a-z-nath/backend-projects
cd task-tracker
```

Run the following command to build the task-tracker app :

```bash
go build -o task-tracker
```

To try or run the project you can execute following commands:

```bash
# To see list of available command:
./task-tracker --help

# To add a task
./task-tracker add "Learn Go Lang"

# To delete a task
./task-tracker delete 1

# To update a task
./task-tracker update 1 "Learn Golang"

# To update status of a task
./task-tracker update 1 -s [todo | in-progress | complete]
# or
./task-tracker update 1 --status [todo | in-progress | complete]

# To list all tasks
./task-tracker list

# To list all task with some status
./task-tracker list -s [todo | in-progress | complete]
# or
./task-tracker list --status [todo | in-progress | complete]
```

---

[Task Tracker](https://roadmap.sh/projects/task-tracker)'s idea is implemented with ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style={display:flex,align-items:center}for-the-badge&logo=go&logoColor=white)

All Courtesy Goes to [ROADMAP.SH](https://roadmap.sh)'s Backend Project Ideas.
