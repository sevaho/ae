# ae

file and folder structure templater 

The goal is to speed up application creation by having a templater that makes a folder structure like:

```
.
├── [sevaho    ]  app
│   ├── [sevaho     ]  ae_lib.py
│   ├── [sevaho     ]  app.py
│   └── [sevaho     ]  requirements.txt
├── [sevaho    ]  config
├── [sevaho    ]  Dockerfile
├── [sevaho    ]  .dockerignore
├── [sevaho    ]  .gitignore
└── [sevaho    ]  Makefile
```

but also for nodejs, golang and nginx.


# Usage

```txt
USAGE: ae [OPTIONS]... [ARGS]...

    ae, file and folder structure templater

OPTIONS:
   -h, --help       display the help and exit

   TO IMPLEMENT: 
   -d, --debug      debug, makes a debug folder and use that as root folder

ARGUMENTS:

   TO IMPLEMENT: 

   init {type} # initializes a whole project

       types:

       nginx
       python
       go
       nodejs
       bare

   file {type} # initializes only 1 file

       types:

       makefile
       gitignore
       dockerignore
       docker {type}

           types:

           nginx
           python
           go
           nodejs
           bare

       makefile {type}

           types:

           nodejs
           python

EXAMPLES:

   ./ae file python

NOTE:

    the -d flag is not implemented yet
    update is not implemented yet
```
