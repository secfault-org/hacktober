Hacktober 2024
==============

This is a beginner-friendly project to help you get started with your binary exploitation journey.

I selected some simple challenges from [picoCTF](https://picoctf.org/) and [Exploit Education - Phoenix](https://exploit.education/phoenix/).

The Tree main topics are:
* Stack-based Buffer Overflow
* Format String Vulnerability
* Heap-based Buffer Overflow

I built this project to play around [Bubbletea](https://github.com/charmbracelet/bubbletea), [Wish](https://github.com/charmbracelet/wish) and serving a CTF in a nice little TUI.



## ToDos
- [ ] Features
    - [x] List Challenges
        - [x] Using a simple bubbles list
    - [x] footer
      - [x] help
    - [ ] Challenge Details
      - [x] Show challenge details
      - [x] Render Markdown
          - [x] short description
          - [x] source-code
      - [x] State and timer, statusbar?
      - [ ] Add extra actions
          - [ ] Download source & executable
              - [ ] Using scp
                - [ ] cannot be used directly. Use copy command and scp middleware
          - [x] Spawn a container
    - [x] Start Challenge
        - [x] Use a cmd to trigger it
        - [x] using podman bindings
    - [ ] Stop Challenge after time and on session end
    - [ ] Explanation pages
        - [ ] Stack
        - [ ] Format-String explaination
        - [ ] Heap
    - [x] Connect via SSH
      - [ ] User by public key hash
    - [x] ASLR
      - [x] mount bin ro
      - [x] create entrypoint
        - [x] set ASLR
        - [x] write flag to file
    - [ ] Earn stars
        - [ ] Keep starts after relogin
        - [ ] Identify user by publickey
        - [ ] Safe state to database
            - [ ] sqlite
                - [ ] sqlx? / sqlc?
    - [ ] Scoreboard
        - [ ] List all users and their stars
        - [ ] Sort by stars
        - [ ] Let user change their name