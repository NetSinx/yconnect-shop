pipeline {
  agent any

  stages {
    stage('Cloning Project from Repository') {
      steps {
        git url: 'https://github.com/NetSinx/yconnect-shop', branch: 'master'
      }
    }
    stage('Build Image') {
      steps {
        sh 'docker build -t user-img server/user/.'
        sh 'docker build -t product-img server/product/.'
        sh 'docker build -t category-img server/category/.'
        sh 'docker build -t cart-img server/cart/.'
        sh 'docker build -t mail-img server/mail/.'
        sh 'docker build -t order-img server/order/.'
      }
    }
    stage('Running Application') {
      steps {
        sh 'docker compose up -d -f server/docker-compose_example.yaml'
      }
    }
  }
}