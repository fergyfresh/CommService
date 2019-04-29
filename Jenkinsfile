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
            sh '''export GOPATH=${\'pwd\'} && export GOBIN=$GOPATH/bin
 && export PATH=$PATH:$GOBIN
 && curl https://glide.sh/get | sh && glide install && go build ./...

 '''
            sh 'echo $GOPATH'
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
}