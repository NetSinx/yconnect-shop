node {
  label 'default-node'

  def app1
  def app2
  def app3
  def app4
  def app5
  def app6
  def dockerImages = ["yasinah22/order-img", "yasinah22/user-img", "yasinah22/product-img", "yasinah22/category-img", "yasinah22/cart-img", "yasinah22/mail-img"]
  def servicesName = ["order", "user", "product", "category", "cart", "mail"]
  def imageTag = "latest"

  stage('Checkout') {
    git url: 'https://github.com/NetSinx/yconnect-shop', branch: 'master'
  }

  stage('Build') {
    app1 = docker.build("${dockerImages[0]}:${imageTag}", "server/${servicesName[0]}/.")
    app2 = docker.build("${dockerImages[1]}:${imageTag}", "server/${servicesName[1]}/.")
    app3 = docker.build("${dockerImages[2]}:${imageTag}", "server/${servicesName[2]}/.")
    app4 = docker.build("${dockerImages[3]}:${imageTag}", "server/${servicesName[3]}/.")
    app5 = docker.build("${dockerImages[4]}:${imageTag}", "server/${servicesName[4]}/.")
    app6 = docker.build("${dockerImages[5]}:${imageTag}", "server/${servicesName[5]}/.")
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