# ToolFormation

**ToolFormation in the early development phase!**

ToolFormation manages the tools used on your machine by defining them in code !

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

# unimplemented
visual-studio-code:
    extension:
        - golang.Go
        - ms-python.python
        - ms-python.vscode-pylance
```
