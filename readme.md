Thoughts while building
==============

## Queue service
I'm not a fan of this, I would rather have a direct client but without knowing the structure of the queue or how it fits into the architecture I wasn't sure what to do, so I created a package to separate it out.

I would rather keep everything within the package that relates to it's function to follow golang "guidelines" but without knowing if the sqs messages are part of the scheduling logic or the rescheduling logic I didn't want to commit to either.

## Testing
Nothing crazy here, it's a simple application so I've just created some basic mocks and check whether the functions behave correctly with them.

## "should it be rescheduled"
I'm not sure if I'm missing something here but other than the remaining completions I can't see any other logic behind whether a new schedule is needed.

Initially I thought I'd need to compare the number of attempts but I can't seem to find a way to know the current participants attempt count. Because I'm unsure I've bunkered the function away and it's pretty trivial to adjust it if needs be.

## asynch
The async stuff is within the rescheduler ScheduleNewQuestionnaire function. While writing this I'm certain there is a better way to organise this but I wanted to keep it pretty readable at a higher level.
I'm not a fan of the ScheduleNewQuestionnaire function but the messiness of it's pretty self contained so can be adjusted without effecting the other code.

Maybe some sort of event based system might work, maybe there's a design pattern I didn't think of, but I dunno, I'm limited on time :).

## Lambda folder
I have chosen to move the main file into a lambda folder, in my experience it had been much easier to organise and build the zip files for each lambda if they're contained within they're own folder and main file like above.

This is obviously just personal preference, I'm just not a fan of coupling my packages to the aws lambda package.
