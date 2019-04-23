# quiz

A super simple quiz with a few questions and a few alternatives for each question. With one correct answer.

**Stack**

Backend - Golang

Database - Just in-memory

gRPC API

CLI client https://github.com/spf13/cobra ( as cli framework )

**User stories**

* User should be able to get questions with answers
* User should be able to select just one answer per question.
* User should be able to answer all the questions and then post his/hers answers and get back how many correct answer there was. and that should be displayed to the user.
* User should see how good he/she did compared to others that have taken the quiz , "You where better then 60% of all quizer"

# **Installation**

`go get -u github.com/eu-ga/quiz`

`make vendor`

`make proto_gen`

`make build_all`

After that, need to go to ./bin/service directory and run a binary.

Now go to ./bin/cli and run quiz program.

To see the questions and answer them use the command: `start`.