data "okta_brands" "test" {
}

data "okta_themes" "test" {
  brand_id = tolist(data.okta_brands.test.brands)[0].id
}
