node('master') {
    try {
        stage('build') {
            checkout scm
            sh "echo Doing the build (test)"
        }
        stage('testing') {
            sh "echo Testing the build"
        }

        stage("deploying") {
            sh "echo Deploying the build"
        }

        } catch {

    } finally {

    }
}