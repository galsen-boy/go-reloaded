# rechargé

L'outil que vous êtes sur le point de construire recevra comme arguments le nom d'un fichier contenant un texte nécessitant des modifications (l'entrée) et le nom du fichier dans lequel le texte modifié doit être placé (la sortie). Voici une liste des modifications possibles que votre programme doit exécuter :

Chaque instance de (hex) doit remplacer le mot précédent par la version décimale du mot (dans ce cas, le mot sera toujours un nombre hexadécimal). (Ex: "1E (hex) fichiers ont été ajoutés" -> "30 fichiers ont été ajoutés")

Chaque instance de (bin) doit remplacer le mot précédent par la version décimale du mot (dans ce cas, le mot sera toujours un nombre binaire). (Ex: "Ça fait 10 (bin) ans" -> "Ça fait 2 ans")

Chaque instance de (up) convertit le mot précédent avec sa version en majuscule. (Ex: "À vos marques, prêts, partez (en haut) !" -> "À vos marques, prêts, partez !")

Chaque instance de (bas) convertit le mot précédent avec sa version en minuscules. (Ex: "Je devrais arrêter de CRIER (bas)" -> "Je devrais arrêter de crier")

Chaque instance de (cap) convertit le mot précédent avec sa version en majuscule. (Ex : "Bienvenue sur le pont de Brooklyn (cap)" -> "Bienvenue sur le pont de Brooklyn")

Pour (bas), (haut), (majuscule) si un nombre apparaît à côté, comme ceci : (bas, <nombre>) il transforme le nombre de mots précédemment spécifié en minuscules, majuscules ou en majuscules en conséquence. (Ex: "C'est tellement excitant (up, 2)" -> "C'est TELLEMENT EXCITANT")
Chaque occurrence des ponctuations ., ,, !, ?, : et ; doit être proche du mot précédent et avec un espace en dehors du suivant. (Ex: "J'étais assis là-bas, puis BAMM !!" -> "J'étais assis là-bas, puis BAMM !!").

Sauf s'il y a des groupes de ponctuation comme : ... ou !?. Dans ce cas, le programme doit formater le texte comme dans l'exemple suivant : "Je pensais... Tu avais raison" -> "Je pensais... Tu avais raison".
Le signe de ponctuation ' sera toujours trouvé avec une autre instance de celui-ci et ils doivent être placés à droite et à gauche du mot au milieu d'eux, sans aucun espace. (Ex: "Je suis exactement comme ils me décrivent : 'génial'" -> "Je suis exactement comme ils me décrivent : 'génial'")

S'il y a plus d'un mot entre les deux marques ' ', le programme devrait placer les marques à côté des mots correspondants (Ex : "Comme Elton John a dit : 'Je suis l'homosexuel le plus connu au monde'" -> "Comme Elton John l'a dit : 'Je suis l'homosexuel le plus connu au monde'")
Chaque instance de a doit être transformée en an si le mot suivant commence par une voyelle (a, e, i, o, u) ou un h. (Ex: "C'était là. Un rocher incroyable!" -> "C'était là. Un rocher incroyable!").