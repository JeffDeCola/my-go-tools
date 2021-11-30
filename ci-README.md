# CONCOURSE CONTINUOUS INTEGRATION

I use concourse to automatic,

* Copy and edit `README.md` to `/docs/_includes/README.md` for
  [GitHub Webpage](https://jeffdecola.github.io/my-go-tools/)
* **TEST** code
* Alert me of the progress via repo status and slack

## PIPELINE

The concourse
[pipeline.yml](https://github.com/JeffDeCola/my-go-tools/blob/master/ci/pipeline.yml)
shows the entire ci flow. Visually, it looks like,

![IMAGE - my-go-tools concourse pipeline - IMAGE](docs/pics/my-go-tools-pipeline.jpg)

## JOBS, TASKS AND RESOURCE TYPES

The concourse jobs and tasks are,

* `job-readme-github-pages` runs task
  [task-readme-github-pages.yml](https://github.com/JeffDeCola/my-go-tools/blob/master/ci/tasks/task-readme-github-pages.yml)
  that kicks off shell script
  [readme-github-pages.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/ci/scripts/readme-github-pages.sh)
* `job-unit-tests` runs task
  [task-unit-tests.yml](https://github.com/JeffDeCola/my-go-tools/blob/master/ci/tasks/task-unit-tests.yml)
  that kicks off shell script
  [unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/tree/master/ci/scripts/unit-tests.sh)

The concourse resources types are,

* `my-go-tools` users a docker image
  [concourse/git-resource](https://hub.docker.com/r/concourse/git-resource/)
  to **PULL** a repo from github
* `resource-repo-status` users a docker image
  [dpb587/github-status-resource](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit
* `resource-slack-alert` users a docker image
  [cfcommunity/slack-notification-resource](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress

For more information on using concourse for continuous integration & deployment,
refer to my
[concourse-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).