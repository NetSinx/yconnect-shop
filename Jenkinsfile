node {
  stage('Cloning Project from Repository') {
    git url: 'https://github.com/NetSinx/yconnect-shop', branch: 'master'
  }

  stage('Build Image') {
    withEnv(['DOCKER_IMAGE=yasinah22/order-img', 'IMAGE_TAG=latest']) {
      checkout scm
      docker.withRegistry('', 'docker-reg') {
        docker.build(DOCKER_IMAGE:IMAGE_TAG).push()
      }
    }
  }
}