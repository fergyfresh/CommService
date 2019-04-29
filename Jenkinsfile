pipeline {
  agent any
  stages {
    stage('Build') {
      parallel {
        stage('Build') {
          steps {
            echo 'Build Step'
            sh '''source ~/.bash_profile



&& 
go build ./...'''
          }
        }
        stage('Communication Step ') {
          steps {
            echo 'Comm Step '
          }
        }
      }
    }
    stage('Qualtiy Gates') {
      steps {
        echo 'Checking Quality Gate'
      }
    }
    stage('Build Docker') {
      steps {
        echo 'Build Docker'
      }
    }
  }
}