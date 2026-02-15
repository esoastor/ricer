# ricer (Work in Progress)
A CLI tool for managing Linux ricing through centralized theme definitions.

`ricer` allows you to define and manage configuration changes across multiple configuration files using a single theme definition.
Instead of manually editing individual files, describe all theme-related changes in one place and apply them with one command.

## Overview
`ricer` operates as follows:
1. Compares the `current` theme with a selected target theme.
2. Generates a changemap describing the differences.
3. Applies the resulting changes to the configured files.

Example:

```sh
# list all themes
ricer themes
# set theme with name nord as current
ricer set nord
```

## Installation
Clone project and run the provided installation script:

```sh
./install.sh
```

The installation script:
- Builds the `ricer` binary.
- Installs it to `/usr/local/bin`.
- Creates an empty configuration file at:

```
~/.config/ricer/config.yaml
```
### Requirements
- Go ≥ 1.25

## Configuration
Edit the configuration file:

```
~/.config/ricer/config.yaml
```

| key          | type   | description                                                                    | required |
|--------------|--------|--------------------------------------------------------------------------------|----------|
| themesPath   | string | Directory containing theme files                                               | yes      |
| subjectsPath | string | Directory containing configuration files to be modified                        | yes      |
| exclude      | list   | Files or directories (relative to `subjectsPath`) excluded from processing     | no       |
| afterCommand | list   | Sh command that will be executed after theme switching. \["command", ..."args"\] | no       |
Ensure that all configured paths exist and are accessible.

## Current Theme
Inside `themesPath`, create a file named `current`.

This file must describe the exact current state of the theme values present in your configuration files.  
It serves as the baseline for generating changemaps when switching to another theme.

See the **Themes** section for formatting details.

## Creating Themes
After creating the `current` theme, use its structure to define additional themes.

All themes must share the same key structure.  
Keys that do not exist in `current` are ignored.

> After initial creation, `current` is managed automatically by `ricer`.  
> To preserve the original state, copy `current` to another file (for example, `initial`) before applying changes.

## Themes
A theme is a file without an extension containing entries in the following format:

```
key = value
```
### Key
A mnemonic identifier (for example, `colorMain`).  
Keys exist for readability and semantic clarity.

### Value
In the `current` theme, `value` must exactly match the string currently present in your configuration files.
In other themes, `value` represents the string that will replace the corresponding `current` value when the theme is applied.

When applying a target theme:
- Each value defined in `current` is replaced with the corresponding value from the selected theme.
- Replacement is performed using exact string matching.
- No structural parsing of configuration files is performed.

### File Tag
To restrict changes to a specific file, use the following block syntax:

```
[file <path>]
...
[endfile]
```

`<path>` must be relative to `subjectsPath`.

### Comments
Single-line comments are supported using:

```
// comment
```
### Example Theme
```
radius = 15px
font = CoolFont

[file coolSoft/styles.css]
se-background = #fabb11 // comment
wp-background = #777faf
[endfile]

[file lalala/hmhmhm.conf]
active_border = #777faf
inactive_border = #891bbf
rw-active_border = #c26480
rw-inactive_border = #fabb11
[endfile]

[file racoon/theme]
background = #777faf
[endfile]
```
## Usage
After configuring the application and creating `current` and additional themes, use the following commands:

### `ricer themes`
Lists all themes found in `themesPath`.

### `ricer subjects`
Lists all files in `subjectsPath`, excluding those defined in `exclude`.  
These files are eligible for modification when applying a theme.

### `ricer changemap <name>`
Displays the changes that would be applied if the theme with the specified `<name>` were set as `current`.

### `ricer set <name>`
Sets the specified theme as `current` and applies the computed changemap to the configured files.

Todo:  
[ ] post install action  
[ ] vars support
