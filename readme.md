Main Code Test
==============

## Code Test Overview

The goal of this test is to see how well you approach situations and to gauge your level with programming languages in general.  
* It is **not** to produce a working piece of software that covers every edge case or the vast array of libraries out there.  
* We are not looking for completeness and typos are allowed :) 
* It is not designed to catch you out
* It doesn't need to run!  (Leave pseudocode describing your intentions if in doubt)
* Pseudocode what you are about to do before going into the detail of the code (so that if you don't complete something, we can see your direction of travel)
* Show your working / Be generous with your comments
* **Use Google!** - do not be afraid to look up references!
* How long you put into the test is up to yourself, but we recommend only a few hours in total.
* Details of what you should do are in task.md


**Read all READMEs carefully before coding and heed the notes and advice in each!**


## Abstract

One way of capturing data from a participant in a study is to get them to fill in a Questionnaire.

## Database Structure

You will be working with these tables:

```mysql
CREATE TABLE participants (
    id VARCHAR(128) PRIMARY KEY NOT NULL,
    name VARCHAR(128) NOT NULL
);

CREATE TABLE questionnaires (
    id VARCHAR(128) PRIMARY KEY NOT NULL,
    study_id VARCHAR(128) NOT NULL,
    name VARCHAR(128) NOT NULL,
    questions JSON NOT NULL,
    max_attempts INT,
    hours_between_attempts INT DEFAULT 24
);

CREATE TABLE scheduled_questionnaires (
    id VARCHAR(128) PRIMARY KEY NOT NULL,
    questionnaire_id VARCHAR(128) NOT NULL,
    participant_id VARCHAR(128) NOT NULL,
    scheduled_at DATETIME NOT NULL,
    status ENUM(pending,completed)
);

CREATE TABLE questionnaire_results (
    id VARCHAR(128) NOT NULL,
    answers JSON NOT NULL,
    questionnaire_id VARCHAR(128) NOT NULL,
    participant_id VARCHAR(128) NOT NULL,
    questionnaire_schedule_id VARCHAR(128),
    completed_at DATETIME
);
```

### Participants

This is your basic user table of participants of a study.

### Questionnaires (`questionnaires`)

This is a table that holds various questionnaires that can be completed and the questions within them.

* `questions` holds the configuration for the questions for this questionnaire.
* The `max_attempts` column will contain the maximum number of times a participant can fill in a given questionnaire - if this is null, then there is no limit to the number of times they can fill it in.
* `hours_between_attempts` is an integer of the number of hours in the future the next questionnaire should be scheduled for. For the sake of this test, this should be the number of hours from the time the questionnaire was filled in.

This table is general / global - that is, it is an abstract table about the questionnaires themselves, and nothing is linked to any specific participants.

### Scheduled Questionnaires (`scheduled_questionnaires`)

This is a table that holds a list of scheduled questionnaires for specific participants to fill in.

* The `scheduled_at` column denotes when a questionnaire should become available to a participant, and the status column denotes the status of the scheduled questionnaire.

For example, participant X might have questionnaire Y scheduled for 6pm today and questionnaire Z scheduled for 9am tomorrow. This would be represented by two rows in this table.

### Questionnaire Results (`questionnaire_results`)

When a participant fills in a questionnaire, the results end up in here.

It must be linked to the participant and questionnaire

If it was part of a schedule, then it will be linked to the schedule as well through the `questionnaire_schedule_id` column. Some questionnaires can be filled in outside of a schedule, in which case this column will be null.

## Scheduling

The `Questionnaire` model represents the abstract questionnaire configuration. A `Questionnaire Schedule` represents a specific request for a participant to fill in a `Questionnaire`.
