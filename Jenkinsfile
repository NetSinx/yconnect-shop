node {
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
    for (int i = 0; i < servicesName.size(); i++) {
      "app"+i} = docker.build("${dockerImages[i]}:${imageTag}", "server/${servicesName[i]}/.")
    }
  }

  stage('Deploy') {
    docker.withRegistry('', 'docker-reg') {
      for (int i = 0; i < servicesName.size(); i++) {
        "app"+i.push()
      }
    }
  }
}