#!/usr/bin/env groovy

def utils = new localhost.Utils()

node('Node1') {
    println "Pipeline to run - ${utils.pipelineToRun}"
    def branch = env.BRANCH_NAME
    println "Branch - ${branch}"
    def imagename = ""
    stage("Checkout") {
        scmCheckout {
            deleteWorkspace = 'true'
        }
    }

    try {
        if (branch == "master") {
            // clean up Docker environment
            stage("Clean up") {
                sh 'docker images prune'
                sh 'docker-compose -f dev-compose.yml down'
            }
            // create Testing environment with Compose
            stage("Unit Tests") {
                // Don't wipe environment after this stage
                scmCheckout {
                    deleteWorkspace = 'false'
                }
                println "Starting Consul"
                sh 'mkdir -p .coverage'
                sh 'chmod 777 .coverage'
                println "Starting Tests"
                println "GOPATH is ${GOPATH}"
                sh 'docker-compose -f dev-compose.yml up --exit-code-from tests --build tests'
                sh 'cp .coverage/cover.out ./cover.out'
                println "Clean up"
                sh 'docker-compose -f dev-compose.yml down'
            }
            // check if it is runnable with Compose
            stage("Functional Tests") {
                // don't wipe environment after this stage
                scmCheckout {
                    deleteWorkspace = 'false'
                }
                sh 'sleep 10'
                println "Starting BDD tests"
                sh 'docker-compose -f dev-compose.yml up --exit-code-from functional --build functional'
                println "Clean up"
                sh 'docker-compose -f dev-compose.yml down'
            }
            // build and deploy changes to higher env
            stage("E1 Deployment"){
               sh 'sleep 10'
               println "Application deployed successfully to E1"
            }

        }
        // TODO Add steps to build and deploy to specific env
    }
    finally {
        println "Clean up"
        sh 'docker-compose -f dev-compose.yml down'
        println 'current build status'
        println "${currentBuild.result}"
    }
}
