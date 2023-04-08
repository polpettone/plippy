# Plippy
- Polpettone Command Line Clipboard Manager

## Features
- remembers entries from system clip board
- bring back old entries to the system clip board
- interactive selection list with search function

## Installation
### Prequisits 
- go environment
### Steps
- clone repo
- ```go install```

## Usage

- show usage 
  - ```plippy --help```
- run plippy to track your clipboard: 
  - ```plippy start &```
  - *Attention*: running plippy will track all your copies to clipboard, also sensitive ones like credentials
    - all entries will go to $HOME/.config/plippy/contents.yaml
    - contents.yaml will contain copied credentials in plain text
- choose an entry 
  - ```plippy list``` opens an interactive list where you can choose an entry
  - hit / to search for an entry in the list
  - hit enter to select an entry, this entry will go to your current system clipboard

- stop plippy
  - kill the plippy process


## Roadmap
- fine tuning
- proper handling of duplicates
- handle critical values like credentials
- statistics
- run as unix daemon

