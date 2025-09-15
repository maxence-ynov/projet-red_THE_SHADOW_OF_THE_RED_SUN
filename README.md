# 🏮 The Shadow of the Red Sun

## 📖 Présentation  
**The Shadow of the Red Sun** est un **jeu d’aventure en ligne de commande (CLI)** dans un **Japon post-apocalyptique**.  
Vous incarnez un **samouraï solitaire** équipé d’un **bras mécanique** et d’un **katana ancestral**.  
Votre mission : **gravir la Tour du Soleil Rouge**, affronter des créatures mythologiques japonaises et découvrir ce qu’il reste à sauver d’un monde en ruine.
---

## ✨ Fonctionnalités principales  
- ⚔ **Combat tour par tour** : affrontez des ennemis avec des schémas d’attaque variés.  
- 💰 **Économie et PNJ** : échangez avec un marchand itinérant, achetez des potions et des ressources.  
- 🛠 **Système de crafting** : fabriquez des équipements améliorant vos statistiques.  
- 🧙 **Progression RPG** : gagnez de l’expérience, débloquez des sorts et améliorez vos capacités.  
- 🎭 **Ambiance narrative** : dialogues et messages immersifs entre les combats.  

---
## 🛠 Installation  
### 1️⃣ Cloner le dépôt  
```bash
git clone https://github.com/maxence-ynov/projet-red_THE_SHADOW_OF_THE_RED_SUN
cd projet-red_ShadowOfTheRedSun/src
```

---
### ▶ Lancer le jeu

Exécutez la commande suivante dans le dossier src :
```bash
go run main.go
```

---
## 🎮 Comment jouer

### -Menu principal :

-📜 Afficher les informations du personnage.

- 🎒 Gérer l’inventaire.

- 🏪 Visiter le marchand.

- ⚔ Lancer un combat d’entraînement.

- ❌ Quitter le jeu.

### -Combat tour par tour :

- Attaquez avec votre katana ou un sort.

- Ouvrez votre inventaire pour utiliser une potion.

- Observez les patterns des ennemis pour planifier vos coups.

- Gagnez de l’expérience et progressez vers des ennemis plus puissants.

---
## 📁 Structure du projet

```bash
projet-red_ShadowOfTheRedSun/
├── docs/
│   └── gestion_projet.md      # Suivi et organisation du projet
├── src/
│   ├── main.go           # Point d’entrée du jeu
│   ├── character.go           # Gestion du personnage
│   ├── inventory.go           # Inventaire et objets
│   ├── merchant.go            # Marchand et crafting
│   ├── combat.go              # Système de combat
│   └── monsters.go            # Données des monstres et patterns
│   └── blacksmith.go          # Forgeron
└── README.md                  # Présentation et instructions
```
---

## 👥 Contribution
| Nom   | Rôle                 | Contributions principales     |
| ----- | -------------------- | ----------------------------- |
| Nom 1 | Développeur          | Combat, intégration générale. |
| Nom 2 | Développeur          | Marchand, économie, crafting. |
| Nom 3 | Scénariste / Dev     | Narration, menus, tests.      |

---



