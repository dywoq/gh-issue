| Command              | Action                      | Input Method                 | Notes                                            |
| -------------------- | --------------------------- | ---------------------------- | ------------------------------------------------ |
| `close`              | Close issues                | CLI arguments                | `<issues-ids> <owner> <repo> <token>`            |
| `close-config`       | Close issues                | JSON config file + issue IDs | `<config> <issues-ids> <filename>`               |
| `get`                | Get issues                  | CLI arguments                | `<issues-ids> <owner> <repo> <token>`            |
| `get-config`         | Get issues                  | JSON config file + issue IDs | `<config> <issues-ids> <filename>`               |
| `generate-md`        | Generate Markdown of issues | CLI arguments                | `<issues-ids> <owner> <repo> <token> <filename>` |
| `generate-md-config` | Generate Markdown of issues | JSON config file + issue IDs | `<config> <issues-ids> <filename>`               |

**Note about JSON config file:**  
The JSON config file used in `*-config` commands must include the following fields:

```json
{
  "owner": "your-github-username",
  "repository": "your-repo-name",
  "token": "your-github-token"
}
