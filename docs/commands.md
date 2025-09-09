# Description

This markdown file contains the documentation about the commands you can use.

## ```close```
```close``` command allows you to close the selected issues from your GitHub repository. 
**The syntax**:
```
gh-issue close <issues-ids> <owner-username> <name> <token> 
```
The token is required to give write access to the issues of your selected GitHub repository.
In `<issues-ids>`, you can use wildcard `*` to select all issues, or as list (ex. `[2, 3, 4]`)

___

## ```get```
```get``` command outputs the selected issues from your GitHub repository into the console.
**The syntax**:
```
gh-issue get <issues-ids> <owner-username> <name> <token> 
```
The token is required to give write access to the issues of your selected GitHub repository.
In `<issues-ids>`, you can use wildcard `*` to select all issues, or as list (ex. `[2, 3, 4]`)

___

## ```generate-md```
```generate-md``` command generates the issues of your selected GitHub repository
into console.
**The syntax**:
```
gh-issue generate-md <issues-ids> <owner-username> <name> <token> <filename>
```	
Filename can be set to "gh.issue.default", then the filename will be set to ```<owner>/<repo> issues.md```.
In `<issues-ids>`, you can use wildcard `*` to select all issues, or as list (ex. `[2, 3, 4]`)
