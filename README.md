# GraphQL servewr with gqlgen library and PostgresQL DB connection in GoLang

# Create Database

mutation createDB {
  createDB
}

# Add users

mutation createUser {
  saveCustomer(input:{
    username:"Toni"
    email:"toni@gmail.com"
    age: 47
    metaData:"{\"firstName\": \"Toni\", \"lastName\":\"Nichev\"}"
  })
}

mutation createUser {
  saveCustomer(input:{
    username:"Jack"
    email:"jack@gmail.com"
    age: 36
    metaData:"{\"firstName\": \"Jack\", \"lastName\":\"Smith\"}"
  })
}

# Add bookmarks


mutation addBookmark {
  addBookmark(
    userId: "1"
    name: "Bookmark two"
    group: "general"
    metaData: """
{
  "url": "https://someurl-two.com",
  "favorite": false,
  "info": [
    {
      "day": "monday",
      "autoOpen": "new-tab"
    },
    {
     "day": "tuesday",
     "autoOpen": "new-tab"
    }
    
  ]
}    
    """
  )
}

# Query user bookmarks

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




