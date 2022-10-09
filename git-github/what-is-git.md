# What is Git?

![git](https://www.freecodecamp.org/news/content/images/2020/05/image-163.png)

Git is a version control system which lets you track changes you make to your files over time. With **Git**, you can revert to various states of your files (like a time traveling machine). You can also make a copy of your file, make changes to that copy, and then merge these changes to the original copy. You are not limited to using Git just for source code files — you can also use it to keep track of text files or even images. This means that Git is not just for developers — anyone can find it helpful.

## What is a Version Control?
A **version control** is basically a way that we as programmers track our code changes, we basically save an initial version of our code into Git. And then when we update the code, we can save it into Git again, and again and again and throughout time as our code continues to change, we can look back at all of the changes we have made over time. This help us to see and understand what we did when, as well as track down bugs, or go back to a previous version of the code if we need to.

## Terminologies
<table>
   <thead>
      <tr>
         <th>Terms</th>
         <th>Description</th>
      </tr>
   </thead>
   <tbody>
      <tr>
         <td>Directory</td>
         <td>Folder</td>
      </tr>
      <tr>
         <td>Terminal or Command Line</td>
         <td>Interface for Text Commands</td>
      </tr>
      <tr>
         <td>CLI</td>
         <td>Command Line Interface</td>
      </tr>
      <tr>
         <td>Directory</td>
         <td>Folder</td>
      </tr>
      <tr>
         <td>cd</td>
         <td>Change Directory</td>
      </tr>
      <tr>
         <td>Code Editor</td>
         <td>Word Processor for Writing Code</td>
      </tr>
      <tr>
         <td>Git</td>
         <td>A tool that trackes the changes in your code overtime</td>
      </tr>
      <tr>
         <td>Github</td>
         <td>A website to host your repositories online</td>
      </tr>
      <tr>
         <td>Repository</td>
         <td>Project, or the folder/place where your project is kept</td>
      </tr>
   </tbody>
</table>

## Install Git

### Windows
1. Download the latest version on [Git for Windows Installer](https://gitforwindows.org/)
2. Follow the instructions as provided in the **Git Setup** wizard screen until the installation is complete
3. Open the windows command prompt and type `git version` to verify Git was installed

### MacOS
1. Download the latest version on [macOS Git Installer](https://sourceforge.net/projects/git-osx-installer/files/git-2.23.0-intel-universal-mavericks.dmg/download?use_mirror=autoselect)
2. Follow the instructions as provided until the installation is complete
3. Open the command prompt "terminal" and type `git version` to verify Git was installed

### Linux
1. To install Git, run the following command:

```bash
dev@dev:~$ sudo apt-get update
dev@dev:~$ sudo apt-get install git
```

2. You can verify the installation by typing:

```bash
dev@dev:~$ git version
```

## Configure Git
The next thing you'll need to do is set your username and email address. Git will use this information to identify who made specific changes to files.

To set your username and email address, type and execute these commands:
```bash
git config --global user.name "YOUR_USERNAME"
git config --global user.email "YOUR_EMAIL"
```

**NOTE**: Make sure to replace `"YOUR_USERNAME"` and `"YOUR_EMAIL"` with the values you choose.

## Reference
* [Install Git](https://github.com/git-guides/install-git)