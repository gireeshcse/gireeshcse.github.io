### Adding existing git repository to present git repo

    git submodule add <repo_url>  <path_of_submodule_if_you_want_else_where>
    git submodule add https://github.com/gireeshcse/backend-ms1.git

Adds **.gitmodules** files 

    [submodule "backend-ms1"]
        path = backend-ms1
        url = https://github.com/gireeshcse/backend-ms1

### Track these repos 

    git add backend-ms1

### Update all modules

    git submodule update --remote

### Update specific module

    git submodule update --remote frontend-ms1


