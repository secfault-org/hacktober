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
    - [ ] List Challenges
        - [ ] Using a simple bubbles list
        - [ ] split layout with details
            - [ ] Preview and on enter select it
    - [ ] Challenge Details
        - [ ] Render Markdown
            - [ ] short description
            - [ ] source-code
            - [ ] State and timer
        - [ ] Add extra actions
            - [ ] Download source & executable
                - [ ] Using scp
            - [ ] Spawn a container
    - [ ] Start Challenge
        - [ ] Use a cmd to trigger it
        - [ ] using podman bindings
    - [ ] Stop Challenge after time and on session end
    - [ ] Explanation pages
        - [ ] Stack
        - [ ] Format-String explaination
        - [ ] Heap
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