# gh-easy-label

GitHub CLI extension that quickly applys predefined label set to your repositories.

## Installation

```sh
gh extension install iced-penguin/gh-easy-label
```

## Usage

**Edit the Configuration File**

```
gh easy-label edit
```

`edit` opens the configuration file (`~/.easy_label.yml`) with your default text editor. Here, you can define label sets, which are collections of related labels. This allows you to efficiently apply multiple labels at once.`

Example:

`~/.easy_label.yml`
```yaml
label-sets:
  default:
    - name: "Type: Documentation"
      description: Documentation
      color: "#0075ca"
    - name: "Type: Bug Fix"
      description: Buf fix
      color: "#B60205"
  priotify:
    - name: "Priority: High"
      description: High priority
      color: "#E99695"
    - name: "Priority: Low"
      description: Low priority
      color: "#BFD4F2"
```


**Show defined label sets**

```
gh easy-label list
```

**Apply label set**

```
gh easly-label apply [label_set]
```

`apply` replaces all existing labels with the labels from the specified label set.`