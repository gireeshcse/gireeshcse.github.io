
### Creating App with the Expo

```
npx create-expo-app@latest
```

To test the app, we will be using expo app from the playstore. We will be testing the application directly on the mobile.


### Starting a development server

```
npx expo start
```

### File Structure

* **app**

    - Contains app's navigation. 

* **assets**
* **components**
* **constants**
* **hooks**
* **scripts**

    - contains reset-project.js
    - **npm run reset-project** moves the **app** directory to **app-example** and creates a new app directory with an **index.tsx**
* **app.json**

    - Contains configuration options for the project and is called the **app** config. 

### Expo CLI

Development tool which is designed to help us move faster during the app development phase.

### EAS CLI

Used to log into our Expo account and compile our app using different EAS services us Build, Update or Submit. We can also use this tool to:

- Publish our app to the app stores.
- Create a development, preview, or production build of our app.
- Create over-the-air(OTA) updates.
- Manage our app credentials
- Create an ad hoc provisioning profile for an IOS device.

```
npm install -g eas-cli
```
### Expo Doctor

Used to diagonse issues in our Expo project.

```
npx expo-doctor
```

Performs checks and analyzes our project's codebase for common issues in app config and package.json files, dependency compatibility, configuration files, and the overall health of the project.

### [ReactNative Directory](https://reactnative.directory/)

### File based Routing

* Expo Router

    - Routing framework
    - Allows us to manage navigation between screens in our app.
    - Uses a file-based method to determine routes inside our app.

    - **App** directory

        - any file we add to this directory becomes a route 

    ```
        app
            index.tsx   // matches '/'
            details.tsx   // matches '/details'
            _layout.tsx
        settings
            index.tsx   // matches '/settings'

    ```

    - **_layout** file

        - Used to define shared UI elements such as headers, tab bars so that persist between different routes.

        ```
        import { Stack } from "expo-router";

        export default function RootLayout() {
            return (
                <Stack>
                <Stack.Screen name="index" />
                <Stack.Screen name="settings" />
                <Stack.Screen name="details" />
                </Stack>
            );
        }
        ```

        ``` <Stack.Screen name={routeName} /> ``` allows defining routes in a stack  

    - Groups

        - Created to organize similar routes or a section of the app.
        - Each group has a layout file, and the grouped directory requires a name inside parentheses. (group)

        ```
            app  
                _layout.tsx  // Root Layout
            (home)
                index.tsx   // matches /
                details.tsx  // matches '/details'
                _layout.tsx  // Home Layout
        ``` 

        - app/_layout.tsx

        ```
        import { Stack } from 'expo-router';

        export default function RootLayout() {
        return (
            <Stack>
            <Stack.Screen name="(home)" />
            </Stack>
        );
        }
        ```

        - app/(home)/_layout.tsx

        ```
        import { Stack } from 'expo-router';

        export default function HomeLayout() {
            return (
                <Stack>
                <Stack.Screen name="index" />
                <Stack.Screen name="details" />
                </Stack>
            );
        }
        ```

    

     



    - Tab navigator

        - Common pattern to navigate between different sections of an app using a tab bar.

        ```
        import { Tabs } from 'expo-router';

        export default function TabLayout() {
        return (
            <Tabs>
                <Tabs.Screen name="(home)" />
                <Tabs.Screen name="settings" />
            </Tabs>
        );
        }
        ```

        ```
        app
            _layout.tsx
            (tabs)
            _layout.tsx
            (home)
                index.tsx
                details.tsx
                _layout.tsx
            settings.tsx
        ```

        app/_layout.tsx

        ```
        import { Stack } from 'expo-router';

        export default function RootLayout() {
        return (
            <Stack>
            <Stack.Screen name="(tabs)" />
            </Stack>
        );
        }
        ```

    - Not found routes

        - To handle routes that are 404s

        - +not-found.tsx

        - This route file matches all unmatched routes from a nested level.

        ```
        import { Link, Stack } from 'expo-router';

        import { View, StyleSheet } from 'react-native';

        export default function NotFoundScreen(){
            return (
                <>
                    <Stack.Screen options={{title: " Oops! This screen doesn't exist."}} />
                    <View>
                        <Link href={"/"}>Go to home screen</Link>
                    </View>
                </>
            )
        }
        ```


