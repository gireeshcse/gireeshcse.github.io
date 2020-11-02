### Nx

Nx is a set of extensible dev tools for monorepos, which helps you develop like Google, Facebook, and Microsoft. Nx site

```
npx create-nx-workspace <workspace name>
```
Running this command will give you a set of options that allows you to create the best workspace for your needs.

empty
oss
web components
angular
angular-nest
react
react-express
Then it will also ask you about which CLI should be activated in your workspace

NX - used for all workspaces
Angular CLI - used for Angular workspaces

Finally, you get asked if you want to use the NX cloud, this is the commercial version of NX. There is a free tier which provides better performance and distribution of cache between co-workers.

#### The Nx workspace

The workspace contains different files and folders.

Apps folder - contains all the applications we want to host
Libs folder - contains all the code, usually, applications will link to code inside this folder
A single node_modules folder for the workspace
The package.json serves as the global file for packages for the whole workspace
workspace.json contains all the configuration of your projects and how to build it
nx.json contains the configuration of the workspace

#### Generating a new React app with Nx

One of the powerful capabilities of Nx is it's code generation and scaffolding mechanism. Nx allows you to generate new apps, libraries, components and more, which don't just generate new code but also modify existing one to fit the new situation.

When using the nx list command, we can see all the plugins available to be used in our workspace. We can use plugins for next, nest, node, angular, etc. There are also community plugins that give us more options like support for vue or go - just to name two of the options.

Creating a React App
We can install @nrwl/react to get a react app started.

Using yarn
```
yarn add @nrwl/react
```
Using npx
```
npx i @nrwl/react
```
Now we can run yarn nx list @nrwl/react to see all the possibilities that this plugin gives us.

Create an application
Create a library
Create a component
Create a redux slice for a project
Set up storybook
Generate storybook story
Create a cypress spec

Let's create the React Application
By running the command yarn nx generate @nrwl/react:application <name> you can also use yarn nx g @nrwl/react:application <name> for short.

When running this command, Nx will ask you what kind of stylesheet format you want to use - for example, CSS, SASS, Stylus, etc.

We can also append the --dry-run flag to the previous command to allow us to see all the files that this command would create.

When creating a React Application there are a few configuration files that are automatically created for us.

.eslintrc
babel.config.json
jest.config.js

```
yarn nx list
yarn nx list @nrwl/react
yarn nx generate @nrwl/react:application --help
yarn nx generate @nrwl/react:application store --dryRun
yarn nx generate @nrwl/react:application store
yarn nx run store:serve
yarn nx run store:serve --port=3001
yarn nx run store:lint

yarn nx g @nrwl/react:lib ui-shared --directory=store 
yarn nx g @nrwl/react:component header --project=store-ui-shared
```

You can install the Visual Studio Code extension Nx Console from the VS Code extension tab

This extension allows you to have a better view of your workspace, what projects you have, common Nx commands and also run project commands.

Using Nx Console will help you save time looking for the commands available in Nx by using the CLI