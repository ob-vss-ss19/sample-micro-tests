pipeline {
    agent none
    stages {
        stage('Setup'){
            agent any
            steps {
                sh 'docker pull ritterho/vss-micro-ci'
            }
        }
        stage('Test') {
            agent {
                docker { 
                    image 'ritterho/vss-micro-ci'                
                }
            }
            steps {
                sh 'go get github.com/t-yuki/gocover-cobertura' // install Code Coverage Tool
                sh 'go test -v -coverprofile=cover.out ./... ' // save coverage info to file
                sh 'gocover-cobertura < cover.out > coverage.xml' // transform coverage info to jenkins readable format                
                publishCoverage adapters: [coberturaAdapter('coverage.xml')] // publish report on Jenkins   
            }
        }
    }
}
