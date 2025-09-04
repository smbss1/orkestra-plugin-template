Template de Plugin pour Orkestra

Ce dépôt est un modèle pour créer de nouveaux plugins pour l'orchestrateur de workflows Orkestra.

Comment Utiliser ce Template

Créer un nouveau dépôt : Cliquez sur le bouton "Use this template" en haut de cette page pour créer votre propre dépôt de plugin.

Nommage du Dépôt : Nommez votre nouveau dépôt en suivant la convention orkestra-plugin-<nom>, par exemple orkestra-plugin-discord ou orkestra-plugin-aws.

Cloner votre nouveau dépôt : Clonez le dépôt que vous venez de créer sur votre machine locale.

Développement de votre Plugin

1. Modifier le Fichier go.mod

Remplacez le nom du module par celui de votre dépôt.

# Avant
module [github.com/orkestra-plugins/orkestra-plugin-template](https://github.com/orkestra-plugins/orkestra-plugin-template)

# Après (par exemple)
module [github.com/votre-utilisateur/orkestra-plugin-discord](https://github.com/votre-utilisateur/orkestra-plugin-discord)

2. Mettre à jour le Fichier main.go

Ouvrez le fichier main.go et modifiez les deux sections marquées // TODO: :

Mettez à jour la liste des nœuds que votre plugin fournit dans la fonction GetCapabilities.

Ajoutez la logique pour vos nouveaux nœuds dans la fonction Execute.

3. Mettre à jour le Workflow de Release

Ouvrez le fichier .github/workflows/release.yml et modifiez la variable d'environnement PLUGIN_NAME pour qu'elle corresponde au nom de votre plugin.

# ...
env:
  # TODO: Remplacez "template" par le nom de votre plugin (ex: "discord", "aws")
  PLUGIN_NAME: "template"
# ...

4. Développer et Tester

Vous pouvez maintenant développer la logique de vos nœuds. Pour tester localement, compilez votre plugin :

go build -o orkestra-plugin-votre-nom .

Puis placez l'exécutable dans le dossier dist/plugins de votre installation d'Orkestra.

Publier votre Plugin
La publication est entièrement automatisée grâce à GitHub Actions.

Commitez et pushez vos changements :

git add .
git commit -m "feat: Ajout du nœud discord/send-message"
git push origin main

Créez un Tag de Version : La publication est déclenchée par la création d'un tag Git.

git tag v1.0.0
git push origin v1.0.0

C'est tout ! GitHub Actions va automatiquement :

Compiler votre plugin pour Windows, macOS (Intel & ARM), et Linux (x64 & ARM).

Créer des archives .zip avec le nommage correct (orkestra-plugin-<nom>_<os>_<arch>.zip).

Créer une nouvelle "Release" sur GitHub avec votre numéro de tag.

Attacher toutes les archives .zip à cette release.

Votre plugin est maintenant prêt à être installé par n'importe quel utilisateur d'Orkestra via la commande :

./orkestra plugins install votre-utilisateur/votre-nom
