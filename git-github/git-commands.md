# Git Commands

<table>
   <thead>
      <tr>
         <th>Command</th>
         <th>Description</th>
      </tr>
   </thead>
   <tbody>
      <tr>
         <td>git init [repository_name]</td>
         <td>Start a new repository</td>
      </tr>
      <tr>
         <td>git clone [url]</td>
         <td>Bring a repository that is hosted somewhere like Github into a folder on your local machine</td>
      </tr>
      <tr>
         <td>git add [file]</td>
         <td>
            - Adds a file to the staging area <br/>
            - Track your files and changes in Git
         </td>
      </tr>
      <tr>
         <td>git commit -m [commit_message]</td>
         <td>
            - Save your files in Git <br>
            - Records or snapshots the file permanently in the version history
         </td>
      </tr>
      <tr>
         <td>git commit -a</td>
         <td>Commits any files you've changed since then and commits any files you've added</td>
      </tr>
      <tr>
         <td>git diff</td>
         <td>Shows the file differences which are not yet staged</td>
      </tr>
      <tr>
         <td>git diff --staged</td>
         <td>Differences between the files in the staging area and the latest version present</td>
      </tr>
      <tr>
         <td>git diff [first branch] [second branch]</td>
         <td>Differences between two branches mentioned</td>
      </tr>
      <tr>
         <td>git reset [file]</td>
         <td>Unstages the file, but it preserves the file contents</td>
      </tr>
      <tr>
         <td>git reset [commit]</td>
         <td>Undoes all the commits after the specified commit and preserves the changes locally</td>
      </tr>
      <tr>
         <td>git reset --hard [commit]</td>
         <td>Discards all history and goes back to the specifed commit</td>
      </tr>
      <tr>
         <td>git status</td>
         <td>Command lists all the files that have to be committed</td>
      </tr>
      <tr>
         <td>git rm [file]</td>
         <td>Deletes the file from your working directory and stages the deletion</td>
      </tr>
      <tr>
         <td>git log</td>
         <td>Used to list the version hsitory for the current branch</td>
      </tr>
      <tr>
         <td>git log --follow [file]</td>
         <td>Lists version history for a file, including the renaming of files also</td>
      </tr>
      <tr>
         <td>git show [commit]</td>
         <td>Shows the metadata and content changes of the specified commit</td>
      </tr>
      <tr>
         <td>git tag [commit_id]</td>
         <td>Used to give tags to the specifed commit</td>
      </tr>
      <tr>
         <td>git branch</td>
         <td>Lists all the local branches in the current repository</td>
      </tr>
      <tr>
         <td>git branch [branch_name]</td>
         <td>Creates a new branch</td>
      </tr>
      <tr>
         <td>git branch -d [branch_name]</td>
         <td>deletes the feature branch</td>
      </tr>
      <tr>
         <td>git checkout [branch_name]</td>
         <td>Used to switch from one branch to another</td>
      </tr>
      <tr>
         <td>git checkout -b [branch_name]</td>
         <td>Creates a new branch and also switches to it</td>
      </tr>
      <tr>
         <td>git merge [branch_name]</td>
         <td>Merges the specified branch's history into the current branch</td>
      </tr>
      <tr>
         <td>git remote add [variable_name] [remote_server_linke]</td>
         <td>Used to connect your local repository to the remote server</td>
      </tr>
      <tr>
         <td>git push [variable_name] master</td>
         <td>Sends the committed changes of master branch to your remote repository</td>
      </tr>
      <tr>
         <td>git push [variable_name] [branch]</td>
         <td>Sends the branch commits to your repository</td>
      </tr>
      <tr>
         <td>git push --all [variable_name]</td>
         <td>Pushes all branches to your remote repository</td>
      </tr>
      <tr>
         <td>git push [variable_name]:[branch_name]</td>
         <td>Deletes a branch on your remote repository</td>
      </tr>
      <tr>
         <td>git pull [repository_link]</td>
         <td>
            - Fetches and merges changes on the remote server to your working directory <br>
            - Download changes from remote repository to your local machine, the opposite of push
         </td>
      </tr>
      <tr>
         <td>git stash save</td>
         <td>Stores all the modified tracked files</td>
      </tr>
      <tr>
         <td>git stash pop</td>
         <td>Restores the most recently stashed files</td>
      </tr>
      <tr>
         <td>git stash list</td>
         <td>Lists all stashed changesets</td>
      </tr>
      <tr>
         <td>git stash drop</td>
         <td>Discards the most recently stashed changeset</td>
      </tr>
   </tbody>
</table>

## Great Read
* [Oh Shit, Git!?!](https://ohshitgit.com/)
* [Github Git Cheat Sheet](https://training.github.com/downloads/github-git-cheat-sheet/)
* [Git Cheat Sheet – 50 Git Commands You Should Know](https://www.freecodecamp.org/news/git-cheat-sheet/)