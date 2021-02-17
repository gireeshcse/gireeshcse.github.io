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

# OJS Pre-Installation

```
sudo chmod -R 777 cache/
sudo chmod -R 777 public/
sudo chmod -R 777 config.inc.php
sudo mkdir /opt/lampstack-7.3.26-0/apache2/files
sudo chmod 777 /opt/lampstack-7.3.26-0/apache2/files
sudo chmod -Rv  777 plugins/
```
* Create a journal and upload apiLoginPlugin and run the following database query

```
INSERT INTO `auth_sources` (`auth_id`, `title`, `plugin`, `auth_default`, `settings`) VALUES
(1, 'API Login', 'apiLogin', 0, NULL);
```
To customize the basic information in adding the author such as making contact info optional when adding collaborator(modified according to our requirement and fork and create according to our requirement)
```
https://github.com/ewhanson/authorRequirements
```
To update the locale and messages. After enabling this locale tab is present in the website settings
```
https://github.com/pkp/customLocale
```

Important files and options present are

```
lib/pkp/locale/en_US/common.po - about.contact,   
user.username,
author.submit.metadata
author.submit.journalSectionDescription
submission.submit.coverNote

lib/pkp/locale/en_US/locale.po -  section.section, section.policy
```

Import users
config.inc.php - cdn : off , smtp settings
create sections in journal settings 
workflow- checklists modification 

#### Upgrade 

* Download new version and extract the files
* Copy the config.inc.php, public, and files_dir(if inside ojs directory) to newly extracted verison
* Also copy any plugins installed from previous verision to new version
* diff config.inc.php config.TEMPLATE.inc.php and check the config info and modify according if any changes are required
* Run the following commands

```
sudo chmod -R 777 cache/
sudo chmod -R 777 public/
sudo chmod -R 777 config.inc.php
sudo chmod -Rv  777 plugins/
```

##### 1. Command-line

- Edit config.inc.php and change "installed = On" to "installed = Off"
- Run the following command from the OJS directory (not including the $):
	`$ php tools/upgrade.php upgrade`
- Re-edit config.inc.php and change "installed = Off" back to
	 "installed = On"

##### 2. Web

If you do not have the PHP CLI installed, you can also upgrade by running a
web-based script. To do so:

- Edit config.inc.php and change "installed = On" to "installed = Off"
- Open a web browser to your OJS site; you should be redirected to the
	installation and upgrade page
- Select the "Upgrade" link and follow the on-screen instructions
- Re-edit config.inc.php and change "installed = Off" back to
	 "installed = On"

## Help

Help Manual

1. Change section policy at

Settings->Journal Settings->Section


If co-authors exists please submit the co-authors consent as well

2. To change the label names or fields information

Settings->Website Settings->Locale

lib/pkp/locale/en_US/common.po

about.contact   = ORG Contact
common.prefix = Prefix(Optional)
common.prefixAndTitle.tip = 
common.subtitle = 

user.po

user.username = PC No

author.po

author.submit.metadata = Report Information
author.submit.journalSectionDescription =  Please select the type of Report
submission.submit.coverNote = Write any comments to Editor(Optional)

locale/en_US/locale.po
section.section = Type of Report
section.policy = Report Policy
reviewer.article.decision.resubmitElsewhere 
reviewer.article.decision.decline
reviewer.article.decision.seeComments
reviewer.article.decision.resubmitHere
reviewer.article.decision.pendingRevisions
reviewer.article.decision.accept

lib/pkp/locale/en_US/submission.po
submission.queries.submission = Pre-Review Discussions
submission.queries.review = Review Discussions

lib/pkp/locale/en_US/grid.po 
grid.action.addQuery = Add discussion


3. Workflow settings

Settings-> Workflow Settings

We can update type of files to be submitted (Author Consent, Article Text etc..)
Add Submission Checklist
Add Author Guidelines
Also update the following
Metadata settings
Reviewer Guidelines
Review Forms,
Review Mode
Single Click Review Access from email 
Default Response dealines
Email Templates

### Important Links

* [Submission Acknowledgement email is not being sent](https://forum.pkp.sfu.ca/t/ojs-3-1-2-the-submission-acknowledgement-email-is-not-being-sent-to-author/54864/9)

Solution:

```
I switched force_dmarc_compliant_from, force_default_envelope_sender and allow_envelope_sender to “On” and uncomment.
```

