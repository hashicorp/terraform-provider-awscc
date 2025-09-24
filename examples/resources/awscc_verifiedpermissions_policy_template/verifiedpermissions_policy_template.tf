resource "awscc_verifiedpermissions_policy_template" "example" {
  policy_store_id = awscc_verifiedpermissions_policy_store.example.policy_store_id
  description     = "example"
  statement       = "permit (principal in ?principal, action in PhotoApp::Action::\"listPhotos\", resource == ?resource) unless { resource.private };"
}


resource "awscc_verifiedpermissions_policy_store" "example" {
  description = "Example verified permissions policy store"
  validation_settings = {
    mode = "STRICT"
  }
  schema = {
    cedar_json = jsonencode({
      "PhotoApp" : {
        "commonTypes" : {
          "PersonType" : {
            "type" : "Record",
            "attributes" : {
              "age" : {
                "type" : "Long"
              },
              "name" : {
                "type" : "String"
              }
            }
          },
          "ContextType" : {
            "type" : "Record",
            "attributes" : {
              "ip" : {
                "type" : "Extension",
                "name" : "ipaddr",
                "required" : false
              },
              "authenticated" : {
                "type" : "Boolean",
                "required" : true
              }
            }
          }
        },
        "entityTypes" : {
          "User" : {
            "shape" : {
              "type" : "Record",
              "attributes" : {
                "userId" : {
                  "type" : "String"
                },
                "personInformation" : {
                  "type" : "PersonType"
                }
              }
            },
            "memberOfTypes" : [
              "UserGroup"
            ]
          },
          "UserGroup" : {
            "shape" : {
              "type" : "Record",
              "attributes" : {}
            }
          },
          "Photo" : {
            "shape" : {
              "type" : "Record",
              "attributes" : {
                "account" : {
                  "type" : "Entity",
                  "name" : "Account",
                  "required" : true
                },
                "private" : {
                  "type" : "Boolean",
                  "required" : true
                }
              }
            },
            "memberOfTypes" : [
              "Album",
              "Account"
            ]
          },
          "Album" : {
            "shape" : {
              "type" : "Record",
              "attributes" : {}
            }
          },
          "Account" : {
            "shape" : {
              "type" : "Record",
              "attributes" : {}
            }
          }
        },
        "actions" : {
          "viewPhoto" : {
            "appliesTo" : {
              "principalTypes" : [
                "User",
                "UserGroup"
              ],
              "resourceTypes" : [
                "Photo"
              ],
              "context" : {
                "type" : "ContextType"
              }
            }
          },
          "createPhoto" : {
            "appliesTo" : {
              "principalTypes" : [
                "User",
                "UserGroup"
              ],
              "resourceTypes" : [
                "Photo"
              ],
              "context" : {
                "type" : "ContextType"
              }
            }
          },
          "listPhotos" : {
            "appliesTo" : {
              "principalTypes" : [
                "User",
                "UserGroup"
              ],
              "resourceTypes" : [
                "Photo"
              ],
              "context" : {
                "type" : "ContextType"
              }
            }
          }
        }
      }
    })
  }
}