node {
  stage('Cloning Project from Repository') {
    git url: 'https://github.com/NetSinx/yconnect-shop', branch: 'master'
  }

  stage('Build Image') {
    withEnv(['DOCKER_IMAGE=order-img', 'IMAGE_TAG=latest', 'URL_REGISTRY=https://hub.docker.com/repositories/yasinah22/order-img']) {
      checkout scm
      docker.withRegistry('$URL_REGISTRY', '665b56ea-0578-4bbf-a417-10aa6e99abb2') {
        docker.build('$DOCKER_IMAGE:$IMAGE_TAG').push()
      }
    }
  }
}