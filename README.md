# Go Reloaded

## Description
**Go Reloaded** est un outil en ligne de commande développé en Golang. Il analyse un texte source et le formate(conversions numériques, correction de la ponctuation, des guillemets et des articles).

Le projet a été conçu selon la méthodologie **TDD (Test-Driven Development)** avec des tests Table-Driven.

---

## Fonctionnalités principales

Le programme traite le texte à travers un pipeline :
1. **Marqueurs de transformation :**
   - Conversion hexadécimal vers décimal : `(hex)`
   - Conversion binaire vers décimal : `(bin)`
   - Modification de la casse : `(up)`, `(low)`, `(cap)`
   - Application sur $n$ mots précédents : `(up, n)` ou `(cap, n)`
2. **Correction de la ponctuation :**
   - Nettoyage des espaces avant les signes (`.,!?:;`)
   - Ajout d'un espace après les signes s'ils sont collés au mot suivant
3. **Gestion des guillemets simples (`'`) :**
   - Associes les guillemets ouvrants et fermants aux mots choisis.
4. **Règle des articles (`a` / `an`) :**
   - Transforme automatiquement `"a"` en `"an"` (et `"A"` en `"An"`) devant un mot commençant par une voyelle ou un `h`.

---

## Structure du projet

* `main.go` : Point d'entrée de l'application (gestion des arguments de la ligne de commande, lecture du fichier d'entrée, orchestration et écriture du fichier de sortie).
* `process.go` : Contient la logique et le pipeline de transformation du texte.
* `process_test.go` : Suite de tests Table-Driven pour tester le projet.

---

## Prérequis

* Avoir [Go](https://golang.org/) installé sur votre machine.

---

## Utilisation

### 1. Lancer les tests
Pour exécuter des tests et vérifier que toutes les règles fonctionnent :
```bash
go test -v
