#dir-tree

You can use this application to draw the directory tree in terminal window. Only tested on widnows.

##How to run

Run the application with the directory path as the first argument.

```
dir-tree.exe C:/Users/Joel/
```

Or run the application without any arguments to scan the current directory.

```
dir-tree.exe
```

You can use relative paths too

```
dir-tree.exe ./
```

The second argument is depth. Use it to specify how deep the tree should go.

```
dir-tree.exe C:/Users/Joel/ 2
```

You can also drag and drop folders in the application.

##Sample output

```
PS C:\dev\dir-tree> .\dir-tree.exe C:/Users/Alireza/Documents/
Please run the program as an administrator for better results
.
┣━┓ /CD Projekt Red
┃ ┗━┓ /Cyberpunk 2077
┃   ┗━┓ /benchmarkResults
┃     ┣━┓ /benchmark_2023-01-19_20-21-39
┃     ┃ ┣━━ frames.csv [77.9 KB]
┃     ┃ ┗━━ summary.json [1.3 KB]
┃     ┣━┓ /benchmark_2023-01-19_20-38-48
┃     ┃ ┣━━ frames.csv [85.6 KB]
┃     ┃ ┗━━ summary.json [1.3 KB]
┃     ┣━┓ /benchmark_2023-01-19_20-42-09
┃     ┃ ┣━━ frames.csv [88.4 KB]
┃     ┃ ┗━━ summary.json [1.3 KB]
┃     ┣━┓ /benchmark_2023-01-19_20-48-14
┃     ┃ ┣━━ frames.csv [111.8 KB]
┃     ┃ ┗━━ summary.json [1.3 KB]
┃     ┗━┓ /benchmark_2023-01-21_00-58-44
┃       ┣━━ frames.csv [101.6 KB]
┃       ┗━━ summary.json [1.3 KB]
┣━━ My Music [0.0 KB]
┣━━ My Pictures [0.0 KB]
┣━━ My Videos [0.0 KB]
┣━┓ /Visual Studio 2022
┃ ┣━┓ /Code Snippets
┃ ┃ ┣━┓ /JavaScript
┃ ┃ ┃ ┗━━ /My Code Snippets
┃ ┃ ┣━┓ /TypeScript
┃ ┃ ┃ ┗━━ /My Code Snippets
┃ ┃ ┣━┓ /Visual Basic
┃ ┃ ┃ ┗━━ /My Code Snippets
┃ ┃ ┣━┓ /Visual C#
┃ ┃ ┃ ┗━━ /My Code Snippets
┃ ┃ ┣━┓ /Visual C++
┃ ┃ ┃ ┗━━ /My Code Snippets
┃ ┃ ┣━┓ /Visual Web Developer
┃ ┃ ┃ ┣━━ /My CSS Snippets
┃ ┃ ┃ ┗━━ /My HTML Snippets
┃ ┃ ┗━┓ /XML
┃ ┃   ┗━━ /My Xml Snippets
┃ ┗━┓ /Templates
┃   ┣━┓ /ItemTemplates
┃   ┃ ┣━━ /C#
┃   ┃ ┣━━ /Extensibility
┃   ┃ ┣━━ /JavaScript
┃   ┃ ┣━━ /TypeScript
┃   ┃ ┣━━ /Visual Basic
┃   ┃ ┣━━ /Visual C++
┃   ┃ ┣━━ /Visual C++ Project
┃   ┃ ┗━━ /Visual Web Developer
┃   ┗━┓ /ProjectTemplates
┃     ┣━━ /C#
┃     ┣━━ /Extensibility
┃     ┣━━ /JavaScript
┃     ┣━━ /TypeScript
┃     ┣━━ /Visual Basic
┃     ┣━━ /Visual C++
┃     ┣━━ /Visual C++ Project
┃     ┗━━ /Visual Web Developer
┗━━ desktop.ini [0.4 KB]

size of "C:/Users/Alireza/Documents/" is 0.47 MB

found 36 empty directories

found 0 empty files

0 errors encountered

press enter to exit...
```