pipeline {
  agent any
  stages {
    stage('Set Environment ') {
      steps {
        echo 'Build Step'
        sh 'source /etc/bashrc'
      }
    }
    stage('Build') {
      steps {
        sh 'go build ./...'
      }
    }
    stage('Quality Gates') {
      steps {
        echo 'Build Docker'
      }
    }
  }
}