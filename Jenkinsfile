node('default-node') {
  def app
  def dockerImages = ["yasinah22/order-img", "yasinah22/user-img", "yasinah22/product-img", "yasinah22/category-img", "yasinah22/cart-img", "yasinah22/mail-img"]
  def servicesName = ["order", "user", "product", "category", "cart", "mail"]
  def imageTag = "latest"

  stage('Checkout') {
    git url: 'https://github.com/NetSinx/yconnect-shop', branch: 'master'
  }

  stage('Build') {
    for (int i = 0; i < dockerImages.size(); i++) {
      app = docker.build("${dockerImages[i]}:${imageTag}", "server/${servicesName[i]}/.")
    }
  }

  stage('Deploy') {
    for (int i = 0; i < dockerImages.size(); i++) {
      app = docker.image("${dockerImages[i]}:${imageTag}")
      
      docker.withRegistry('', 'docker-reg') {
        app.push()
      }
    }
  }

  stage('Cleanup') {
    docker.image()
  }
}