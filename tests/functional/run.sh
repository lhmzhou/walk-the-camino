#!/bin/bash

#-------------------------------------#



#--------------Search Value-----------#
search="$1"
tags="$2"
#updategitProjects="$3"
nodeReport="$3"
rootDir=$(pwd)
sleep 10
isTestPass() {
        if [[ ${retVal} = *"FAIL"* ]]; then
            echo  $1
            echo $retVal
            exit 0
        else
            echo $2
        fi
}

if [ "$#" == 0 ];then
   echo "Must provide search, tags, and bool to update gitProjects."
   exit 0
fi

if [ "$search" == "" ];then
   echo "Must provide search predicate."
   exit 0
fi

if [ "$tags" == "" ];then
  echo "Must provide tags to execute bdd test cases. Multiple values can be given comma separated"
  exit 0
fi

#if [ "$updategitProjects" == "" ];then
#   echo "Must provide true or false for updating test cases and results to gitProjects."
#   exit 0
#fi

for feature in ./*/
do
    sprint="${feature%*/}"
    case $feature in (*"$search"*)
#       echo "${feature}"
        cd "$feature/" || exit 1
        sprint="${feature%*/}"
        currDir=$(pwd)
        echo "currently executing for folder -" $currDir

#       find current sprint iteration from options.json
#       to use in html report generation

        configFilePath="$currDir"/config/options.json
#       echo "$configFilePath"
        while IFS=',' read -r key value
        do
        if [[ ${key} = *"iteration"* || ${key} = *"testName"* ]]; then
            key1=$(echo $key | awk -F \" '{print $2}')
            value1=$(echo $key | awk -F \" '{print $4}')
            eval ${key1}=\${value1}
        fi
        done < "$configFilePath"
#        execute go bdd logic for every feature folder for given tags
#        create node html report
#        and , execute go bdd for gitProjects to update generated test result in gitProjects

        echo "go bdd test is being executed for tags -- " $tags
        if [ "$nodeReport" == false ];then
            retVal=$(go test . -v --godog.format=cucumber --godog.tags="$tags")
        fi
        if [ "$nodeReport" != false ];then
            retVal=$(go test . -v --godog.tags="$tags")
        fi

        if [ "$nodeReport" == false ];then
            echo "go bdd test report generation is in progress -- " $iteration
            echo $testName $currDir $iteration $rootDir
            nodeVal=$(node "$rootDir"/reports/main.js fileName="$testName" directory="$currDir"/ iteration="$iteration" output="$rootDir")
            echo $nodeVal
        fi
#       go run ../cleanJson.go -t="testing:" -p="$currDir"/report/in/ -f=report.json
        echo "Return Value" $retVal
        if [[ $retVal = *exit1* ]];then
            echo "Exiting with exit code 1 due to test failure"
            exit 1
        fi

        cd "../" || exit 1
        currDir=$(pwd)
        echo $currDir
    ;;esac
done