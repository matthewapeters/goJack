# goJack Graphviz #

This docker container will provide the means to run Graphviz without installing it as a local application.
It is included here primarily for illustrative purposes -- teams can ensure that automated processes are
consistent across platforms by leveraging pre-configured environments for particular jobs.

The `build` script checks if there is a `graphviz` image tag locally -- if not present, the script builds and tags the docker image.

These components are used by `graph-game` along with `diagram-game-flow`.

Learn more about [Graphviz here](https://graphviz.org).
