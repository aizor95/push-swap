# Push-swap

Ce projet contient deux programmes écrits en Go :

- **push-swap** reçoit une liste d’entiers et affiche des instructions pour trier la pile `a` en ordre croissant.
- **checker** reçoit la même liste et lit les instructions depuis l’entrée standard. Il affiche `OK` si la pile `a` est triée et si la pile `b` est vide, `KO` sinon.

Le premier entier fourni est placé au sommet de la pile `a`.

## Compilation

Depuis le dossier du projet, avec Go installé :

```powershell
go build -o checker.exe .
go build -o push-swap.exe .
```

Les deux exécutables sont construits à partir de `piles.go`. Le programme choisit son comportement en fonction du nom de l’exécutable.

## Utilisation sous Windows PowerShell

Afficher les instructions produites :

```powershell
$numbers = "4 67 3 87 23"
& .\push-swap.exe $numbers
```

Vérifier automatiquement le résultat :

```powershell
$numbers = "4 67 3 87 23"
& .\push-swap.exe $numbers | & .\checker.exe $numbers
```

Le checker affiche `OK` si la suite d’instructions trie correctement la pile.

## Utilisation sous Linux ou macOS

```sh
go build -o checker .
go build -o push-swap .

ARG="4 67 3 87 23"
./push-swap "$ARG" | ./checker "$ARG"
```

## Instructions disponibles

Les instructions manipulent le sommet des piles :

| Instruction | Effet |
|---|---|
| `sa`, `sb` | Échange les deux premiers éléments de `a` ou `b` |
| `ss` | Effectue `sa` et `sb` |
| `pa`, `pb` | Déplace le premier élément de `b` vers `a`, ou de `a` vers `b` |
| `ra`, `rb` | Place le premier élément en bas de la pile correspondante |
| `rr` | Effectue `ra` et `rb` |
| `rra`, `rrb` | Place le dernier élément au sommet de la pile correspondante |
| `rrr` | Effectue `rra` et `rrb` |

## Algorithmes

- Pour **7 nombres ou moins**, `push-swap` utilise une recherche en largeur (BFS) afin de trouver une suite minimale d’instructions.
- Au-delà de **7 nombres**, il utilise un tri radix sur les rangs des valeurs. Cette méthode trie la pile, mais ne garantit pas la suite globalement la plus courte.

## Entrées invalides

Le programme affiche `Error` sur la sortie d’erreur si les arguments ne sont pas des entiers, s’ils contiennent des doublons, ou si le checker reçoit une instruction inconnue ou mal formée. Sans argument, les programmes n’affichent rien.
