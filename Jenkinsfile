node {
  def app1
  def app2
  def app3
  def app4
  def app5
  def app6

  stage('Checkout') {
    git url: 'https://github.com/NetSinx/yconnect-shop', branch: 'master'
  }

  stage('Build') {
    withEnv([
      SERVICES_NAME=["order", "user", "product", "category", "cart", "mail"],
      DOCKER_IMAGE=["yasinah22/order-img", "yasinah22/user-img", "yasinah22/product-img", "yasinah22/category-img", "yasinah22/cart-img", "yasinah22/mail-img"],
      IMAGE_TAG="latest"
    ]) {
      app1 = docker.build("${DOCKER_IMAGE[0]}:${IMAGE_TAG}", "server/${SERVICES_NAME[0]}/.")
      app2 = docker.build("${DOCKER_IMAGE[1]}:${IMAGE_TAG}", "server/${SERVICES_NAME[1]}/.")
      app3 = docker.build("${DOCKER_IMAGE[2]}:${IMAGE_TAG}", "server/${SERVICES_NAME[2]}/.")
      app4 = docker.build("${DOCKER_IMAGE[3]}:${IMAGE_TAG}", "server/${SERVICES_NAME[3]}/.")
      app5 = docker.build("${DOCKER_IMAGE[4]}:${IMAGE_TAG}", "server/${SERVICES_NAME[4]}/.")
      app6 = docker.build("${DOCKER_IMAGE[5]}:${IMAGE_TAG}", "server/${SERVICES_NAME[5]}/.")
    }
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