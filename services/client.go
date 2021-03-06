package services

import (
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

// CreateNewClient creates a new client application entry
func CreateNewClient(name, description, secret, scopes, canonicalURI, redirectURI string) models.Client {
    var client models.Client = models.Client{
        Name: name,
        Description: description,
        Secret: secret,
        Scopes: scopes,
        CanonicalURI: canonicalURI,
        RedirectURI: redirectURI,
        Type: models.ConfidentialClient,
    }

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Create(&client)
    return client
}

// FindOrCreateClient attempts to find a client application by its name; otherwise, it creates a new one
func FindOrCreateClient(name string) models.Client {
    var client models.Client

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Where("name = ?", name).First(&client)
    if dataStoreSession.NewRecord(client) {
        client = models.Client{
            Name: name,
            Secret: models.GenerateRandomString(64),
            CanonicalURI: "localhost",
            RedirectURI: "/",
            Scopes: models.PublicScope,
            Type: models.PublicClient,
        }
        dataStoreSession.Create(&client)
    }
    return client
}

// FindClientByKey gets a client application by its key
func FindClientByKey(key string) models.Client {
    var client models.Client

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Where("key = ?", key).First(&client)
    return client
}

// FindClientByUUID gets a client application by its UUID
func FindClientByUUID(uuid string) models.Client {
    var client models.Client

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Where("uuid = ?", uuid).First(&client)
    return client
}

// ClientAuthentication gets a client application by its key-secret pair
func ClientAuthentication(key, secret string) models.Client {
    var client models.Client

    client = FindClientByKey(key)
    if client.ID != 0 && client.Authentic(secret) {
        return client
    }
    return models.Client{}
}
