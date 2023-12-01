Go Test
=======

In the `rescheduler` directory, there is the bare bones of a simple Lambda function that handles an event that is fired when a user submits a questionnaire.

You can (and should) add any imports as necessary and split into different files as you wish.

You will be primarily interacting with a database and a queue.

## Task One

* Read the requirements for how schedules work in the main readme.md
* Determine if a new questionnaire schedule should be saved to the database.
* If so, save one in the database, and push a new message to SQS that a new schedule has been created.
* If not, push a new message to SQS that the user has completed all of their alloted scheduled questionnaires.

## Task Two

Make the database and queue operations happen asynchronously.
