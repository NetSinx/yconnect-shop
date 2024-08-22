pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        sh "git clone https://github.com/NetSinx/yconnect-shop"
        sh "docker build -t user-img yconnect-shop/server/user/."
      }
    }
  }
}