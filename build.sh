#! /bin/bash



if [ $1 == "prod" ]; then
    ps -ef | grep -w yuwenlanzi | awk '{print $2}' | xargs kill -9
    git checkout master
    git fetch origin master
    git merge FETCH_HEAD
    cp ./conf/prod.app.conf ./conf/app.conf
    go install
    go build -o yuwenlanzi main.go
    nohup ./yuwenlanzi > /dev/null 2>&1 &
fi

if [ $1 == "dev" ]; then
    ps -ef | grep -w yuwenlanzi-dev | awk '{print $2}' | xargs kill -9
    git checkout dev
    git fetch origin dev
    git merge FETCH_HEAD
    cp ./conf/dev.app.conf ./conf/app.conf
    go install
    go build -o yuwenlanzi-dev main.go
    nohup ./yuwenlanzi-dev > /dev/null 2>&1 &
fi

exit