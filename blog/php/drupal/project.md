```
mkdir drupal
cd drupal 
curl -sSL https://www.drupal.org/download-latest/tar.gz | tar -xz --strip-components=1


mkdir sites/default/files
chmod a+w sites/default/files

cp ./sites/default/default.settings.php  ./sites/default/settings.php
// Before installing
chmod a+w sites/default/settings.php
// After Installing
chmod go-w sites/default/settings.php
chmod go-w sites/default

mkdir uploads
chmod a+w uploads
// uncomment and change the path
# $settings['file_public_path'] = 'sites/default/files';
// same for the private files
# $settings['file_private_path'] = '';
```

### Upgrading

Take backup of files(sites/default/,modules,themes etc.) and database.

Open settings.php (/sites/default/settings.php) i

```
$settings['update_free_access'] = TRUE;
```
Put your site in maintenance mode
```
mkdir temp
cd temp
curl -sSL https://www.drupal.org/download-latest/tar.gz | tar -xz --strip-components=1
```

In your site hosting directory, delete the core and vendor directories, and all files that are not in a subdirectory, including .htaccess, composer.json, and autoload.php. Don’t delete custom and customized files because you may end up losing the custom functionality stored in them.
Copy the core and vendor directories and the non-custom/non-customized files that you deleted in the preceding step from the temporary directory to your site directory.

Run the update.php script using either of the following: 

Visit http://www.example.com/update.php in your browser (where www.example.com is your site’s URL). Click Continue in the first screen to run the updates and successfully complete the script. 

Open settings.php (/sites/default/settings.php) i

```
$settings['update_free_access'] = FALSE;
```

Disable maintenance mode

Clear cache.

https://www.drupal.org/project/backup_migrate
https://www.drupal.org/docs/contributed-modules/mime-mail/how-email-works-in-drupal