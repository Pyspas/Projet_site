package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, `<!doctype html>
<html lang="fr">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Mon site Go</title>
  <style>body{font-family:Arial,Helvetica,sans-serif;line-height:1.6;margin:2rem;}h1{color:#2c3e50;}a{color:#3498db;}</style>
</head>
<body>
  <h1>Bienvenue sur mon site Go</h1>
  <p>Ce site est servi par un serveur Go simple.</p>
  <p>Pour démarrer, exécute <code>go run main.go</code>.</p>
  <p><a href="/about">À propos</a></p>
</body>
</html>`)
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, `<!doctype html>
<html lang="fr">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>À propos</title>
  <style>body{font-family:Arial,Helvetica,sans-serif;line-height:1.6;margin:2rem;}h1{color:#2c3e50;}a{color:#3498db;}</style>
</head>
<body>
  <h1>À propos</h1>
  <p>Ceci est une page de démonstration pour un site internet en Go.</p>
  <p><a href="/">Retour à l'accueil</a></p>
</body>
</html>`)
    })

    addr := ":8080"
    fmt.Printf("Serveur lancé sur http://localhost%s\n", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}
