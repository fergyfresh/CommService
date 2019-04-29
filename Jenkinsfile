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
            sh '''export PATH=$PATH:$GOBIN
'''
            sh 'echo $GOPATH'
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
    GOBIN = "$GOPATH/bin"
  }
}