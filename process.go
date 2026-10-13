package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// orchestre les 4 étapes
func Process(s string) string {
	s = appliquerMarqueurs(s)
	s = appliquerPonctuation(s)
	s = appliquerGuillemets(s)
	s = appliquerArticles(s)
	return s
}

// Gestion des marqueurs
func appliquerMarqueurs(texte string) string {
	regexNormalisation := regexp.MustCompile(`\((hex|bin|up|low|cap),\s*(\d+)\)`)
	texte = regexNormalisation.ReplaceAllString(texte, "($1,$2)")
	mots := strings.Fields(texte)
	resultat := []string{}
	regexMarqueur := regexp.MustCompile(`^\((hex|bin|up|low|cap)(?:,(\d+))?\)([.,!?:;]*)$`)
	for _, m := range mots {
		correspondances := regexMarqueur.FindStringSubmatch(m)
		if len(correspondances) == 0 {
			resultat = append(resultat, m)
			continue
		}
		commande := correspondances[1]
		quantite := 1
		if correspondances[2] != "" {
			quantite, _ = strconv.Atoi(correspondances[2])
		}
		for i := len(resultat) - 1; i >= 0 && quantite > 0; i-- {
			resultat[i] = transformerMot(resultat[i], commande)
			quantite--
		}
	}
	return strings.Join(resultat, " ")
}

// Appliques une transformation à un mot
func transformerMot(mot string, commande string) string {
	switch commande {
	case "hex":
		if nombre, erreur := strconv.ParseInt(mot, 16, 64); erreur == nil {
			return strconv.FormatInt(nombre, 10)
		}
	case "bin":
		if nombre, erreur := strconv.ParseInt(mot, 2, 64); erreur == nil {
			return strconv.FormatInt(nombre, 10)
		}
	case "up":
		return strings.ToUpper(mot)
	case "low":
		return strings.ToLower(mot)
	case "cap":
		caracteres := []rune(strings.ToLower(mot))
		if len(caracteres) > 0 {
			caracteres[0] = unicode.ToUpper(caracteres[0])
		}
		return string(caracteres)
	}
	return mot
}

// Ponctuation
func appliquerPonctuation(texte string) string {
	regexAvant := regexp.MustCompile(`\s+([.,!?:;]+)`)
	texte = regexAvant.ReplaceAllString(texte, "$1")
	regexApres := regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)
	texte = regexApres.ReplaceAllString(texte, "$1 $2")
	return texte
}

// Guillemets simples
func appliquerGuillemets(texte string) string {
	mots := strings.Fields(texte)
	resultat := []string{}
	estOuvert := false
	for i := 0; i < len(mots); i++ {
		m := mots[i]
		if m == "'" {
			if !estOuvert {
				if i+1 < len(mots) {
					mots[i+1] = "'" + mots[i+1]
				} else {
					resultat = append(resultat, "'")
				}
				estOuvert = true
			} else {
				if len(resultat) > 0 {
					resultat[len(resultat)-1] = resultat[len(resultat)-1] + "'"
				} else {
					resultat = append(resultat, "'")
				}
				estOuvert = false
			}
		} else {
			resultat = append(resultat, m)
		}
	}
	return strings.Join(resultat, " ")
}

// Articles a et an
func appliquerArticles(texte string) string {
	mots := strings.Fields(texte)
	for i := 0; i < len(mots)-1; i++ {
		m := mots[i]
		if m == "a" || m == "A" {
			motSuivant := mots[i+1]
			motNettoye := strings.TrimPrefix(motSuivant, "'")
			if len(motNettoye) > 0 {
				premiereLettre := unicode.ToLower([]rune(motNettoye)[0])
				if estVoyellePlusH(premiereLettre) {
					if m == "a" {
						mots[i] = "an"
					} else {
						mots[i] = "An"
					}
				}
			}
		}
	}
	return strings.Join(mots, " ")
}

// Voyelles
func estVoyellePlusH(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h'
}
