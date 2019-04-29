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
    PATH = "/home/jenkins/firefox:/usr/lib64/qt-3.3/bin:/usr/local/bin:/usr/bin:/usr/local/go/bin:/home/jenkins/workspace/ice_Pipeline_feature_jenkinsfile/bin"

  }
}
