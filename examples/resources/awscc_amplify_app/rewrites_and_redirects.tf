resource "awscc_amplify_app" "example" {
  name = "app"

  custom_rules = [
    {
      # Reverse Proxy Rewrite for API requests
      # https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html#reverse-proxy-rewrite
      source = "/api/<*>"
      status = "200"
      target = "https://api.example.com/api/<*>"
    },
    {
      # Redirects for Single Page Web Apps (SPA)
      # https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html#redirects-for-single-page-web-apps-spa
      source = "</^[^.]+$|\\.(?!(css|gif|ico|jpg|js|png|txt|svg|woff|ttf|map|json)$)([^.]+$)/>"
      status = "200"
      target = "/index.html"
    },
  ]
}
