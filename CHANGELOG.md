# Change Log

All notable changes to this project will be documented in this file.
This project follows Semantic Versioning.


## 1.5.0 - 2026-04-25

### ✨ New

- Render task description to markdown [#1](https://github.com/mabd-dev/Gira/issues/1). We support:
  - Heading
  - Text with marks like (strong, stroke, underline, code, link)
  - Paragraph (combination of texts)
  - Rule (line separator)
  - Bullet List, Ordered List, TaskList, Nested Lists
  - CodeBlock
  - Table

### 🔧 Updates

- Add api call to change jira task assignee [#2](https://github.com/mabd-dev/Gira/issues/2)
- Show `Ends Today` on sprints ending today, instead of showing `0 days remaining`


## 1.4.2 - 2026-04-10

### ✨ New

- assign to `self` command (:assign self)
- filter tasks by `self` (/@self)
- **Open Task in browser**: press `<super>+enter`


### 🔧 Updates

- **Home Screen Focus Loss**: retain focused pain when clicking outside of taskList/taskDetail panes


## 1.4.1 - 2026-04-06

### ✨ New

- **Build Flavor**: configure app dependencies, appID and behavior based on flavor
- NEW **Light Theme Selector**
- NEW **Color Theme Selector**. Supported themes: `solarized blue`, `catppuccin`
- **Open Task in Browser**: using `<super>+Enter` keybinding. <super> is `cmd` or `windows` key for `mac` or `windows` repectively


## [v1.4.0] - 2026-04-03

### ✨ New

- NEW (behavior): auto navigate to auth screen when token is expired 
- NEW **Settings Screen**


### 🔧 Updates

- Implemented `AES-256-GCM` encryption for jira credentials on **Android** and **JVM** platforms


