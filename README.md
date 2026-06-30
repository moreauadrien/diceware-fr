# diceware-fr

Générateur de phrases de passe selon la méthode [Diceware](https://fr.wikipedia.org/wiki/Diceware), avec une liste de 7776 mots français.

## Utilisation

```sh
# Avec Nix
nix run github:moreauadrien/diceware-fr

# Avec Go
go run .
```

```
$ diceware-fr
croustillant-bouilloire-tapisserie-bavardage-squelette-guitare
```

### Options

| Option          | Défaut | Description                  |
| --------------- | ------ | ---------------------------- |
| `--size N`      | `6`    | Nombre de mots               |
| `--separator S` | `-`    | Séparateur entre les mots    |

```sh
$ diceware-fr --size 4 --separator "."
aboyer.pelouse.tremblement.cerceau
```

## Compilation

```sh
go build -o diceware-fr .
```

## Licence

MIT
