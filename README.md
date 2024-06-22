# The Classic Todo List
## Description
This is a simple todo list application that allows users to add, edit, and delete tasks. The application is built using React and Redux. The application is also responsive and can be used on mobile devices.

## Idea of Design on V0
https://v0.dev/r/ighKwI6iTCD

## Features planned
- [X] Create a task
- [X] Define a priority for a task
- [X] Mark a task as completed
- [X] Define a due date for a task
- [X] Persist tasks in local storage
- [X] Create a Docker image
- [X] send a notification when a task is due
- [X] delete a task
- [ ] edit a task
- [ ] revisit cron design

side quest:
- [ ] Implement a CRON library to send notifications

primus_sucks


## How to deploy

```batch
docker build -t todo-app .

docker tag todo-app europe-west1-docker.pkg.dev/grand-radio-333810/cloud-run-source-deploy/todo-app

docker push europe-west1-docker.pkg.dev/grand-radio-333810/cloud-run-source-deploy/todo-app:latest


gcloud beta run deploy todo-app \
    --image europe-west1-docker.pkg.dev/grand-radio-333810/cloud-run-source-deploy/todo-app:latest \
    --region europe-west1 --allow-unauthenticated --port 8080 \
    --min-instances 1 \
    --no-cpu-throttling  --project grand-radio-333810 --tag init

```