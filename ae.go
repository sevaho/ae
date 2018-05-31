// -----------------------------------------------------------------------------------------------------------------------------
// GENERAL
// -----------------------------------------------------------------------------------------------------------------------------
//
// author: Sebastiaan Van Hoecke
// mail: sebastiaan@sevaho.io
//
// NOTE:
//
// -----------------------------------------------------------------------------------------------------------------------------

package main

import (

    "fmt"
    "strings"
    "os"
    "bufio"
    "log"
    "path/filepath"

)

// -----------------------------------------------------------------------------------------------------------------------------
// CONSTANTS
// -----------------------------------------------------------------------------------------------------------------------------

/*
 *  This maps the key for a file used in lib.go to the real file name for creation
 */

var KEYS = map[string]string {

    "dotenv"                : "app/.env",
    "makefile"              : "Makefile",
    "dockerfile"            : "Dockerfile",
    "nginx_http"            : "config/nginx.conf",
    "nginx_https"           : "config/nginx.conf",
    "app.js"                : "app/app.js",
    "ae_lib.js"             : "app/ae_lib.js",
    "app.go"                : "app/app.go",
    "ae_lib.go"             : "app/ae_lib.go",
    "app.py"                : "app/app.py",
    "ae_lib.py"             : "app/ae_lib.py",
    "requirements.sh"       : "app/requirements.sh",
    "requirements.txt"      : "app/requirements.txt",
    "gitignore"             : ".gitignore",
    "dockerignore"          : ".dockerignore",
    "package.json"          : "app/package.json",

}

// -----------------------------------------------------------------------------------------------------------------------------
// FUNCTIONS
// -----------------------------------------------------------------------------------------------------------------------------

func check (e error) {

    if e != nil {

        panic(e)

    }

}

func create_file (s string, c string) {

    f, err := os.Create(KEYS[s])

    check(err)

    defer f.Close()

    w := bufio.NewWriter(f)
    n, err := w.WriteString(c)
    fmt.Printf("wrote %d bytes \tto %s\n", n, KEYS[s])

    w.Flush();

}

func update () {

    fmt.Println("Updating the binary")
    log.Fatal("NOT IMPLEMENTED YET")

}

// -----------------------------------------------------------------------------------------------------------------------------
// MAIN
// -----------------------------------------------------------------------------------------------------------------------------

var debug = ""

func main () {

    if (len(os.Args) < 2) {

        fmt.Println(get_content("usage"))
        os.Exit(0)

    }

    if (os.Args[1] == "-d") {

        os.Args = append(os.Args[:1], os.Args[1+1:]...)
        debug = "debug/"

    }

    if (len(os.Args) < 3) {

        fmt.Println(get_content("usage"))

    } else  {

        cwd, err := os.Getwd()

        check(err)

        hostname, err := os.Hostname()

        check(err)

        cwd = filepath.Base(cwd)

        fmt.Println("Working directory: ", cwd)

        switch (os.Args[1]) {

            case "completion":

                fmt.Println(get_content("completion"))
                os.Exit(0)

                break

            case "update":

                update()

                break

            case "init":

                os.Mkdir("config", 0755)
                os.Mkdir("app", 0755)

                // create_file("gitignore", get_content("gitignore"))
                // create_file("dockerignore", get_content("dockerignore"))

                switch (os.Args[2]) {

                    case "go":

                        makefile := get_content("makefile")
                        makefile = strings.Replace(makefile, "<<<REPLACE>>>", cwd, -1)
                        create_file("makefile", makefile)

                        dockerfile := get_content("dockerfile_go")
                        dockerfile = strings.Replace(dockerfile, "<<<REPLACE>>>", hostname, -1)
                        create_file("dockerfile", dockerfile)

                        create_file("requirements.sh", get_content("deploy"))
                        create_file("ae_lib.go", get_content("ae_lib.go"))
                        create_file("app.go", get_content("app.go"))

                        break

                    case "nodejs":

                        makefile := get_content("makefile")
                        makefile = strings.Replace(makefile, "#npm start --prefix app # nodejs", "npm start --prefix app # nodejs", -1)
                        makefile = strings.Replace(makefile, "#rm -rf app/etc # nodejs", "rm -rf app/etc # nodejs", -1)
                        makefile = strings.Replace(makefile, "#rm -rf app/node_modules # nodejs", "rm -rf app/node_modules # nodejs", -1)
                        makefile = strings.Replace(makefile, "#npm install --prefix app # nodejs", "npm install --prefix app # nodejs", -1);
                        makefile = strings.Replace(makefile, "#COUNT=$(npm audit --prefix app| grep 'High' | wc -l); [ $COUNT -gt 1 ] && exit # nodejs", "#COUNT=$(npm audit --prefix app| grep 'High' | wc -l); [ $COUNT -gt 1 ] && exit # nodejs", -1)
                        makefile = strings.Replace(makefile, "<<<REPLACE>>>", cwd, -1)

                        create_file("makefile", makefile)

                        dotenv := get_content("dotenv")
                        create_file("dotenv", dotenv)

                        dockerfile := get_content("dockerfile_nodejs")
                        dockerfile = strings.Replace(dockerfile, "<<<REPLACE>>>", hostname, -1)
                        create_file("dockerfile", dockerfile)

                        create_file("app.js", get_content("app.js"))
                        create_file("ae_lib.js", get_content("ae_lib.js"))

                        c := get_content("package.json")
                        c = strings.Replace(c, "<<<NAME>>>", cwd, -1)
                        c = strings.Replace(c, "<<<AUTHOR>>>", hostname, -1)
                        create_file("package.json", c)

                        break

                    case "python":

                        makefile := get_content("makefile")
                        makefile = strings.Replace(makefile, "#pip install --user -r app/requirements.txt", "pip install --user -r app/requirements.txt", -1)
                        makefile = strings.Replace(makefile, "#python app/app.py # python", "python app/app.py # python", -1)
                        makefile = strings.Replace(makefile, "<<<REPLACE>>>", cwd, -1)
                        create_file("makefile", makefile)


                        dockerfile := get_content("dockerfile_python")
                        dockerfile = strings.Replace(dockerfile, "<<<REPLACE>>>", hostname, -1)
                        create_file("dockerfile", dockerfile)

                        create_file("requirements.txt", get_content("requirements.txt"))
                        create_file("ae_lib.py", get_content("ae_lib.py"))
                        create_file("app.py", get_content("app.py"))

                        break

                    case "nginx":

                        makefile := get_content("makefile")
                        makefile = strings.Replace(makefile, "<<<REPLACE>>>", cwd, -1)
                        create_file("makefile", makefile)

                        scanner := bufio.NewReader(os.Stdin)

                        fmt.Println("1 for http, 2 for https")

                        http, _, err := scanner.ReadRune()

                        scanner = bufio.NewReader(os.Stdin)

                        fmt.Println("Give the server name: ")

                        server_name, err := scanner.ReadString('\n')
                        server_name = strings.TrimSuffix(server_name, "\n")

                        if err != nil {

                            log.Fatal(err)
                        }

                        // readrune will give you the ascii character which means a '1' is 49 in ascii. so I minify by 48 ('0') 
                        // to get the value as int
                        http2 := int(http - '0')

                        if (http2 == 1) {

                            nginx_http := get_content("nginx_http")
                            nginx_http = strings.Replace(nginx_http, "<<<REPLACE>>>", server_name, -1)
                            create_file("nginx_http", nginx_http)

                        } else  if (http2 == 2) {

                            nginx_https := get_content("nginx_https")
                            nginx_https = strings.Replace(nginx_https, "<<<REPLACE>>>", server_name, -1)
                            create_file("nginx_https", nginx_https)


                        } else {

                            fmt.Println(get_content("usage"))
                            os.Exit(1)

                        }

                        dockerfile := get_content("dockerfile_nginx")
                        dockerfile = strings.Replace(dockerfile, "<<<REPLACE>>>", hostname, -1)
                        create_file("dockerfile", dockerfile)

                        break

                    case "bare":

                        makefile := get_content("makefile")
                        makefile = strings.Replace(makefile, "<<<REPLACE>>>", cwd, -1)
                        create_file("makefile", makefile)

                        dockerfile := get_content("dockerfile_bare")
                        dockerfile = strings.Replace(dockerfile, "<<<REPLACE>>>", hostname, -1)
                        create_file("dockerfile", dockerfile)

                        break

                    default:

                        fmt.Println(get_content("usage"))
                        break

                }

                break

            case "file":

                switch(os.Args[2]) {

                    case "makefile":

                    if (len(os.Args) < 4) {

                        makefile := get_content("makefile")
                        makefile = strings.Replace(makefile, "<<<REPLACE>>>", cwd, -1)
                        create_file("makefile", makefile)

                    } else  {

                        switch(os.Args[3]) {

                    case "nodejs":

                        makefile := get_content("makefile")
                        makefile = strings.Replace(makefile, "#npm start --prefix app # nodejs", "npm start --prefix app # nodejs", -1)
                        makefile = strings.Replace(makefile, "#rm -rf app/etc # nodejs", "rm -rf app/etc # nodejs", -1)
                        makefile = strings.Replace(makefile, "#rm -rf app/node_modules # nodejs", "rm -rf app/node_modules # nodejs", -1)
                        makefile = strings.Replace(makefile, "#npm install --prefix app # nodejs", "npm install --prefix app # nodejs", -1);
                        makefile = strings.Replace(makefile, "#COUNT=$(npm audit --prefix app| grep 'High' | wc -l); [ $COUNT -gt 1 ] && exit # nodejs", "#COUNT=$(npm audit --prefix app| grep 'High' | wc -l); [ $COUNT -gt 1 ] && exit # nodejs", -1)
                        makefile = strings.Replace(makefile, "<<<REPLACE>>>", cwd, -1)
                        create_file("makefile", makefile)

                        break

                    case "python":

                        makefile := get_content("makefile")
                        makefile = strings.Replace(makefile, "#pip install --user -r app/requirements.txt", "pip install --user -r app/requirements.txt", -1)
                        makefile = strings.Replace(makefile, "#python app/app.py # python", "python app/app.py # python", -1)
                        makefile = strings.Replace(makefile, "<<<REPLACE>>>", cwd, -1)
                        create_file("makefile", makefile)

                        break

                            default:

                                fmt.Println(get_content("usage"))

                                break

                        }

                    }

                    case "gitignore":

                        create_file("gitignore", get_content("gitignore"))

                        break

                    case "dockerignore":

                        create_file("dockerignore", get_content("dockerignore"))

                        break

                    case "docker":

                    if (len(os.Args) < 4) {

                        fmt.Println(get_content("usage"))

                    } else  {

                        switch(os.Args[3]) {

                        case "bare":

                        dockerfile := get_content("dockerfile_bare")
                        dockerfile = strings.Replace(dockerfile, "<<<REPLACE>>>", hostname, -1)
                        create_file("dockerfile", dockerfile)

                            break

                        case "go":

                        dockerfile := get_content("dockerfile_go")
                        dockerfile = strings.Replace(dockerfile, "<<<REPLACE>>>", hostname, -1)
                        create_file("dockerfile", dockerfile)

                            break

                        case "nginx":

                        dockerfile := get_content("dockerfile_nginx")
                        dockerfile = strings.Replace(dockerfile, "<<<REPLACE>>>", hostname, -1)
                        create_file("dockerfile", dockerfile)

                            break

                        case "python":

                        dockerfile := get_content("dockerfile_python")
                        dockerfile = strings.Replace(dockerfile, "<<<REPLACE>>>", hostname, -1)
                        create_file("dockerfile", dockerfile)

                            break

                        case "nodejs":

                        dockerfile := get_content("dockerfile_nodejs")
                        dockerfile = strings.Replace(dockerfile, "<<<REPLACE>>>", hostname, -1)
                        create_file("dockerfile", dockerfile)

                            break

                            default:

                                fmt.Println(get_content("usage"))

                                break

                        }

                    }

                    default:

                        fmt.Println(get_content("usage"))

                        break

                }

                break


            default:

                fmt.Println(get_content("usage"))

                break

        }

    }

}
