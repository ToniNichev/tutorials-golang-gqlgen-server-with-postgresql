# GraphQL servewr with gqlgen library and PostgresQL DB connection in GoLang



query getCustomerLabel {
  getCustomer(customerId:"4") {
    customerId
    metaData
    age
    username
    email
  }
}



query getCustomerLabel {
  getCustomerByMetaData(metaData:"meta_data->>'one' = '1111'") {
    customerId
    metaData
    age
    username
    email
  }
}




