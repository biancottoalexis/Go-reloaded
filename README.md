# Go-reloaded

Go-reloaded, c'est un petit correcteur de texte que j'ai codé en Go.
Tu lui donnes un fichier texte, il corrige ce qui ne va pas, et il écrit le résultat dans un autre fichier.

## Comment le lancer

Il faut lui donner deux fichiers : le texte à corriger, puis le fichier où le résultat sera écrit.

```
go run . sample.txt result.txt
```

Si tu oublies un des deux fichiers, le programme te rappelle comment l'utiliser.
Et si le fichier de départ n'existe pas, il te le dit aussi.

## Ce qu'il sait faire

On peut écrire des petites commandes entre parenthèses juste après un mot pour le modifier.

`(hex)` et `(bin)` transforment un nombre en nombre normal (en décimal). Par exemple, `1E (hex)` devient `30`, et `10 (bin)` devient `2`.

`(up)`, `(low)` et `(cap)` changent les majuscules : `go (up)` donne `GO`, `STOP (low)` donne `stop`, et `paris (cap)` donne `Paris`.
On peut aussi ajouter un nombre pour modifier plusieurs mots d'un coup : `so cool (up, 2)` donne `SO COOL`.

Le programme corrige aussi quelques erreurs tout seul :

- `a` devient `an` devant une voyelle ou un h, comme en anglais : `a apple` donne `an apple`.
- La ponctuation (`. , ! ? : ;`) est collée au mot d'avant : `hello ,world !` donne `hello, world!`. Les groupes comme `...` ou `!!` restent ensemble.
- Les apostrophes sont collées aux mots qu'elles entourent : `' hello '` donne `'hello'`.

## Comment le code est rangé

Au début, j'avais tout écrit dans `main.go`. Ça marchait, mais le fichier devenait vraiment long et difficile à relire.
Du coup, j'ai mis chaque règle dans son propre dossier :

```
main.go         lit le fichier, corrige chaque ligne, écrit le résultat
commandes/      (hex) (bin) (up) (low) (cap)
articles/       a -> an
ponctuation/    . , ! ? : ;
apostrophes/    ' mot '
toustest/       les tests
```

`main.go` découpe le texte ligne par ligne, puis chaque ligne passe par les règles, une après l'autre.

## Les tests

J'ai regroupé les tests dans le dossier `toustest` : il y a un test pour chaque règle, et un dernier avec un exemple qui mélange un peu tout.

```
go test ./...
```

## Ce qui m'a donné du fil à retordre

Le `(up, 2)` m'a bloqué un moment. Comme je découpe le texte avec les espaces, la commande se retrouve coupée en deux : `(up,` d'un côté et `2)` de l'autre. Il a fallu aller chercher le nombre dans le morceau suivant.

Pour les apostrophes, la difficulté, c'est de savoir si c'est celle qui ouvre ou celle qui ferme, pour la coller du bon côté.

Et au début, tout le texte ressortait sur une seule ligne. J'ai corrigé ça en traitant le fichier ligne par ligne.