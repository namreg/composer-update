## Composer update

Dead simple utility that used for updating [composer package manager](https://getcomposer.org) in multiple projects. Every project updates concurrently.

Available options:
```
-basepath string
    Base path for projects. (default "/var/www")
-branch string
    Branch to create. (default "master")
-dep string
    Dependency to update in format "name@master".
-projects string
    Projects to update separated by a comma.
```

How to use:

_Linux_
```bash
./bin/composer-update.am64.linux \
    -projects comments.api,search.api \
    -dep onlinerby/user-sdk@dev-session-reset-time \
    -branch session-reset-time \
    -basepath /var/www
```

_MacOS_
```bash
./bin/composer-update.am64.darwin \
    -projects comments.api,search.api \
    -dep onlinerby/user-sdk@dev-session-reset-time \
    -branch session-reset-time \
    -basepath /var/www
```
