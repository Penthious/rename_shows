## What
This is a project to help out in the process of renaming a collection of episodes for a given show

## Folder structure for shows
The project assumes your folder structure consists of a root parent, followed by seasons, then episodes

└── Show Name\
         ├── season 01\
         │   ├── episode1.xyz\
         │   ├── episode2.xyz\
         │   └── episode3.xyz\
         └── season 02\
                ├── episode1.xyz\
                ├── episode2.xyz\
                └── episode3.xyz
 
## How
`go run . -path="/path/to/your/shows/directory"`

Prompt will ask for you to input a show name, IE: rick and morty (this name should match what your folder name is that contains all of your rick and morty shows)