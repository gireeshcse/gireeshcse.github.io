### Storybook

* UI development tool.
* To build small atomic components and complex web pages for web applications
* Helps us to document the components for reuse and visually test them to prevent bugs.
* Helps to build UI components isolated from the bussiness logic and context of app.

#### Benefits of component driven development

* Quality

    - Verify that UIs work in different scenarios by building components in isolation and defining their relevant states.

* Durability

    -  Pinpoint bugs down to the detail by testing at the component level. It’s less work and more precise than testing screens.

* Speed

    - Assemble UIs faster by reusing existing components from a component library or design system.

* Efficiency

    - Parallelize development and design by decomposing UI into discrete components then sharing the load between different team members.

#### Setup React Storybook

```
npx create-react-app taskbox
cd taskbox
# Add storybook
npx -p @storybook/cli sb init

# To run test runner(Jest)
yarn test --watchAll

# Start the component explorer on port 9009:
yarn storybook

# Run the frontend app proper on port 3000:
yarn start

```

To download folders from the github **degit**

```
npx degit chromaui/learnstorybook-code/public/font public/font
npx degit chromaui/learnstorybook-code/public/icon public/icon
```

#### Example 

In storybook, there are two basic levels of organisations
    - Component 
    - Component Child Stories

To tell the **Storybook** about the component we are documenting, we create a default export containing

- component
    - The component itself

- title
    - The title to be shown in the sidebar    

- excludedStories
    - the exports in the storybook that should not be rendered as stories by StoryBook

- argTypes 
    - specify **args** behaviour in each story


```
const Template = args => <Task { ...args } />;

export const Default = Template.bind({}) 
```

**Template.bind({})** is technique in javascript to copy a function

```
# src/components/Task.js

import React from "react";

export default function Task({ task: { id, title, state }, onArchiveTask, onPinTask }) {
    return (
      <div className="list-item">
        <input type="text" value={title} readOnly={true} />
      </div>
    );
}

# src/components/Task.stories.js

import React from 'react';
import Task from './Task'
export default {
    component : Task,
    title : "Task"
}
const Template = args => <Task { ...args } />;
export const Default = Template.bind({}) 
Default.args = {
    task : {
        id:1,
        title : "Test Task",
        state:"TASK_INBOX",
        updatedAt: new Date(2020,0,1,9,0)
    }
}
export const Archived = Template.bind({})
Archived.args = {
    task : {
        ...Default.args.task,
        state:"TASK_ARCHIVED"
    }
}
export const Pinned = Template.bind({})
Pinned.args = {
    task : {
        ...Default.args.task,
        state:"TASK_PINNED"
    }
}

# configure the component with the StoryBook
// .storybook/main.js

module.exports = {
  stories: ['../src/components/**/*.stories.js'],
  addons: [
    '@storybook/addon-links',
    '@storybook/addon-essentials',
    '@storybook/preset-create-react-app',
  ],
};

// .storybook/preview.js

import '../src/index.css';

// Configures Storybook to log the actions(onArchiveTask and onPinTask) in the UI.
export const parameters = {
  actions: { argTypesRegex: '^on[A-Z].*' },
};

```

#### Snapshot testing

 The practice of recording the “known good” output of a component for a given input and then flagging the component whenever the output changes in future. 

 **Storyshots addon**

 This is used to create the snapshot test for each of the stories

 ```
 yarn add -D @storybook/addon-storyshots react-test-renderer
 ```

 ```
 // src/storybook.test.js
import initStoryshots from '@storybook/addon-storyshots';
initStoryshots();
```

 We have created snapshot test for each of our Task stories. If we change the implementation of Task, we’ll be prompted to verify the changes.

#### Decorators

* A way to provide arbitrary wrappers to stories
* They can also be used to wrap stories in “providers” –i.e. library components that set React context

```
export default {
  component: TaskList,
  title: 'TaskList',
  decorators: [story => <div style={{ padding: '3rem' }}>{story()}</div>],
};
```

#### Wire in Data

```
yarn add react-redux redux
```
