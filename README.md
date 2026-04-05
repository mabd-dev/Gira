# Gira


[![Download](https://img.shields.io/badge/download-latest%20apk-blue)](../../releases/latest)
[![Changelog](https://img.shields.io/badge/changelog-available-green)](CHANGELOG.md)
![Kotlin](https://img.shields.io/badge/Kotlin-Android-blue)
![Compose](https://img.shields.io/badge/Jetpack%20Compose-UI-brightgreen)
![Architecture](https://img.shields.io/badge/Architecture-MVI-orange)


***

## The Problem -- Why JIRA is painful

Have you ever opened JIRA to check something in the sprint to suddenly find the UI has changed, Yes it happens. Not to mention that the UI is clunky, slow and sometimes does not work. Often while in standup meeting and I try to click on a task to see more details but the popup won't show, clicking other task still nothing. The only solution is to refresh the whole page. THANK YOU JIRA
Another big issue I see in that the the main sprint screen is packed with buttons and actions.

This is personal but JIRA is mouse-heavy and I am a [vim](https://www.vim.org/) user :)


## Key Features

- Vim Keybindings 
- Bulk and repeatable edits
- Navigate between tasks without touching your mouse
- Simple and minimal UI
- Configurable keybindings


## Download and Install

> Currently the app is only supported on MACOS desktop, other platforms will be supported soon

### Macos

Steps:
1. Download DMG


## Quick Start

First thing first, authentication. You need to go to your jira account settings and generate `api token` (you need that later)

### Keybindings Cheatsheet

- `j`, `k`, `gg`, `G`: navigate tasks list
- `Ctrl+h`, `Ctrl+l`: move focus between tasks list and task details panes
- `mt`: move task to `todo`
- `mp`: move task to `in-progress`
- `md`: move task to `done`
- `assign [X]`: assign task to X user
- `/filter @username`: filter all tasks by username

and much more

## Roadmap

### Known issues 

- [ ] Fix tasks list focus management

### Future features

- [ ] Parse and render task description as standard markdown
- [ ] Ability to modify task: title, description, versions, components and custom fields
- [ ] Filter & Command mode: suggestions & auto-complete while typing


### Platform supporting

- Windows
- Linux
- Android
- iOS

## Feedback and Community

If you encounter:
- Bugs
- UI issues
- Feature suggestions

Please open an issue.

Feedback is welcome.

## About

After years of professionaly working as software developer, and recently being scrum-master. I have decided to build this for me and my time. 
