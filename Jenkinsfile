pipeline {
  agent {
    node {
      label 'QA_AUTO'
    }

  }
  stages {
    stage('Build ') {
      parallel {
        stage('Build ') {
          steps {
            sh 'echo $GOPATH'
            sh 'mkdir bin '
            sh ' export GOBIN=$GOPATH/bin'
            sh ' curl https://glide.sh/get | sh '
            sh 'glide install '
            sh 'go build ./...'
          }
        }
        stage('Slack Message') {
          steps {
            echo 'Send slack message here '
          }
        }
      }
    }
    stage('Docker Container') {
      steps {
        echo 'Build Docker Container'
      }
    }
    stage('Deploy to Kubernetes') {
      parallel {
        stage('Deploy to Kubernetes') {
          steps {
            echo 'Deploy to Kubernetes'
          }
        }
        stage('Slack Message ') {
          steps {
            echo 'Send slack with deploy '
          }
        }
      }
    }
  }
  environment {
    GOPATH = "${WORKSPACE}"
    GOBIN = "${WORKSPACE + '/bin'}"
  }
}