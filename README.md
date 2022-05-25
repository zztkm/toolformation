# ToolFormation

ToolFormation は自分のPCにインストールするツール類をコードで定義することで管理することができます

```yaml
package-manager: homebrew
homebrew:
	formula:
		- node
		- r
		- pyenv
	cask:
		- visual-studio-code
		- rstudio
		- iterm2

visual-studio-code:
	extension:
		- golang.Go
		- ms-python.python
		- ms-python.vscode-pylance
```
