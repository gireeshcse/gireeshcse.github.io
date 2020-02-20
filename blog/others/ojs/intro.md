https://docs.pkp.sfu.ca/dev/documentation/en/getting-started

Fork and clone the project

    https://github.com/pkp/ojs
    cd ojs
    git submodule update --init --recursive
    cp config.TEMPLATE.inc.php config.inc.php

In config.inc.php

set enable_cdn to off to use local copies of libraries instead of CDNs

set database authentication info.

Install dependencies with composer.

    composer --working-dir=lib/pkp update
    composer --working-dir=plugins/paymethod/paypal update
    composer --working-dir=plugins/generic/citationStyleLanguage update

Install dependencies with NPM.

    npm install
    npm run build

CYPRESS_INSTALL_BINARY=/home/msg/Public/cypress.zip  npm install cypress

To get changes to the application that were made after you forked, add the upstream remote.


git remote add upstream git@github.com:pkp/ojs.git
cd lib/pkp
git remote add upstream git@github.com:pkp/pkp-lib.git
cd ../ui-library
git remote add upstream git@github.com:pkp/ui-library.git
cd ../..

Run the following commands whenever you want to pull the latest changes to your repository.

# Update the app
git checkout master
git pull upstream master
git push

# Update the pkp-lib submodule
cd lib/pkp
git checkout master
git pull upstream master
git push

# Update the ui-library submodule
cd ../ui-library
git checkout master
git pull upstream master
git push

cd ../..

Updates #
When you have pulled down changes from the upstream remote, run the following to sync the lib/pkp and lib/ui-library submodules.

git submodule update --init --recursive
You may need to update dependencies and rebuild the JavaScript package.

composer --working-dir=lib/pkp update
npm install
npm run build

Sometimes a code change will modify the database structure and you will need to run the upgrade script.

php tools/upgrade.php upgrade