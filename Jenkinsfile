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
            node(label: 'QA_AUTO') {
              sh 'go build ./... '
              sh 'whoami'
            }

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
            sh 'sh "echo Deploy to Kubernetes"'
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