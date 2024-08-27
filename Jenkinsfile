node {
  environment {
    SERVICES_NAME = ["order", "user", "product", "category", "cart", "mail"]
    DOCKER_IMAGE = ["yasinah22/order-img", "yasinah22/user-img", "yasinah22/product-img", "yasinah22/category-img", "yasinah22/cart-img", "yasinah22/mail-img"]
    IMAGE_TAG = "latest"
  }

  stage('Checkout') {
    git url: 'https://github.com/NetSinx/yconnect-shop', branch: 'master'
  }

  stage('Build') {
    def app1 = docker.build("${env.DOCKER_IMAGE[0]}:${env.IMAGE_TAG}", "server/${env.SERVICES_NAME[0]}/.")
    def app2 = docker.build("${env.DOCKER_IMAGE[1]}:${env.IMAGE_TAG}", "server/${env.SERVICES_NAME[1]}/.")
    def app3 = docker.build("${env.DOCKER_IMAGE[2]}:${env.IMAGE_TAG}", "server/${env.SERVICES_NAME[2]}/.")
    def app4 = docker.build("${env.DOCKER_IMAGE[3]}:${env.IMAGE_TAG}", "server/${env.SERVICES_NAME[3]}/.")
    def app5 = docker.build("${env.DOCKER_IMAGE[4]}:${env.IMAGE_TAG}", "server/${env.SERVICES_NAME[4]}/.")
    def app6 = docker.build("${env.DOCKER_IMAGE[5]}:${env.IMAGE_TAG}", "server/${env.SERVICES_NAME[5]}/.")
  }

  stage('Deploy') {
    docker.withRegistry('', 'docker-reg') {
      app1.push()
      app2.push()
      app3.push()
      app4.push()
      app5.push()
      app6.push()
    }
  }
}