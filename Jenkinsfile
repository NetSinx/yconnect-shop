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
        sh 'docker build -t user-img yconnect-shop/server/user/.'
        sh 'docker build -t product-img yconnect-shop/server/product/.'
        sh 'docker build -t category-img yconnect-shop/server/category/.'
        sh 'docker build -t cart-img yconnect-shop/server/cart/.'
        sh 'docker build -t mail-img yconnect-shop/server/mail/.'
        sh 'docker build -t order-img yconnect-shop/server/order/.'
      }
    }
    stage('Running Application') {
      steps {
        sh 'docker compose up -d -f yconnect-shop/server/docker-compose_example.yaml'
      }
    }
  }
}