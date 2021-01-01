[Getting Started](https://docs.pkp.sfu.ca/dev/documentation/en/getting-started)

[Documnetation](https://pkp.sfu.ca/ojs/doxygen/master/html/index.html)

# Project Setup

## Database setup

* login into mysql using root
* Create ojs database
* Create a new mysql User

        CREATE USER 'ojs'@'localhost' IDENTIFIED BY 'ojspassword';
        GRANT ALL PRIVILEGES ON ojs. * TO 'ojs'@'localhost';
        FLUSH PRIVILEGES; # to reload all the privileges

[More info on permissions](https://www.digitalocean.com/community/tutorials/how-to-create-a-new-user-and-grant-permissions-in-mysql)

## Initial Setup

Fork and clone the project

    https://github.com/pkp/ojs
    cd ojs
    git submodule update --init --recursive
    cp config.TEMPLATE.inc.php config.inc.php

In config.inc.php

* Set **enable_cdn** to off to use local copies of libraries instead of CDNs

* Set database authentication info.

To get changes to the application that were made after you forked, add the upstream remote.

    git remote add upstream git@github.com:pkp/ojs.git
    cd lib/pkp
    git remote add upstream git@github.com:pkp/pkp-lib.git
    cd ../ui-library
    git remote add upstream git@github.com:pkp/ui-library.git
    cd ../..

Install dependencies with composer.

    composer --working-dir=lib/pkp update
    composer --working-dir=plugins/paymethod/paypal update
    composer --working-dir=plugins/generic/citationStyleLanguage update

Install dependencies with NPM.

    npm install
    npm run build

If downloading cypress takes times by doing npm install [download cypress](https://www.cypress.io/) and run the following

    CYPRESS_INSTALL_BINARY=/home/msg/Public/cypress.zip  npm install cypress

Initially application access will redirect to install. Create admin account

# To run the application using PHP’s built-in server.

php -S localhost:8000

# Run the following commands whenever you want to pull the latest changes to your repository.

## Update the app
git checkout master
git pull upstream master
git push

## Update the pkp-lib submodule
cd lib/pkp
git checkout master
git pull upstream master or git pull
git push // not required

## Update the ui-library submodule
cd ../ui-library
git checkout master
git pull upstream master
git push // not required

cd ../..

## Updates

When you have pulled down changes from the upstream remote, run the following to sync the lib/pkp and lib/ui-library submodules.

    git submodule update --init --recursive

You may need to update dependencies and rebuild the JavaScript package.

    composer --working-dir=lib/pkp update
    npm install
    npm run build

Sometimes a code change will modify the database structure and you will need to run the upgrade script.

    php tools/upgrade.php upgrade

### XXL format to import users

```
<?xml version="1.0"?>
<PKPUsers  xmlns="http://pkp.sfu.ca" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"  xsi:schemaLocation="http://localhost:8080/ojs/pkp-users.xsd">
    <user_groups xsi:schemaLocation="http://pkp.sfu.ca pkp-users.xsd">
        <user_group>
            <role_id>65536</role_id>
            <context_id>3</context_id>
            <is_default>true</is_default>
            <show_title>false</show_title>
            <permit_self_registration>true</permit_self_registration>
            <permit_metadata_edit>false</permit_metadata_edit>
            <name locale="en_US">Author</name>
            <abbrev locale="en_US">AU</abbrev>
            <stage_assignments>1:3:4:5</stage_assignments>
        </user_group>
    </user_groups>
    <users>
        <user>
            <givenname locale="en_US">Ram</givenname>
            <country>IN</country>
            <email>ram@example.in</email>
            <username>ram</username>
            <password is_disabled="false" must_change="false">
                <value>
                SomePassword
                </value>
            </password>
            <inline_help>true</inline_help>
            <auth_id>1</auth_id>
            <user_group_ref>Author</user_group_ref>
        </user>
    </users>
</PKPUsers>
```