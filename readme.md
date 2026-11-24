Push-Swap

//Présentation

Push-Swap est un projet réalisé en langage Go. Le but est de trier une liste de nombres dans l'ordre croissant en utilisant deux piles (A et B) et différentes instructions de déplacement.

Le programme push-swap génère les instructions nécessaires pour effectuer le tri. Le programme checker vérifie si le résultat est correct.

//Commandes

Compiler le projet

go build -o push-swap.exe .

Lancer Push-Swap

./push-swap.exe 2 1 3 6 5 8

Compiler le Checker

go build -o checker.exe ./checker

Vérifier le résultat

./push-swap.exe 2 1 3 6 5 8 | ./checker.exe 2 1 3 6 5 8

Résultats du Checker

OK : les nombres sont correctement triés et la pile B est vide.

KO : le résultat du tri est incorrect.

Error : une erreur est détectée dans les arguments ou les instructions.

//Instructaions disponibles

sa, sb, ss : échanger les éléments.

pa, pb : déplacer un élément entre les piles.

ra, rb, rr : effectuer une rotation.

rra, rrb, rrr : effectuer une rotation inverse.

Technologies

Go (Golang)

Visual Studio Code

                              by RAYANE TOURE