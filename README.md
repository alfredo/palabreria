# Palabrería

Generates a list of random words in the provided dictionaries.

## Rationale

I was tired of generating passwords from english only lists I wanted to use some words in my mother tonge and other languages that I love.

## Bootstrap

It requies a list of word lists that are meant to be used as dictionaries to be available in the `dict` folder.

In OSX the dictionaries can be generated by:

- Install aspell with brew: `brew install aspell`
- Generate the word list of the languages wanted by running at the root of the repository:
```
aspell -d $LANGUAGE dump master | aspell -l $LANGUAGE expand > dict/en | sed $'s/ /\\\n/g'
```
Where `$LANGUAGE` is the language wanted. e.g.

```
aspell -d en dump master | aspell -l en expand > dict/en | sed $'s/ /\\\n/g'
aspell -d es dump master | aspell -l es expand  > dict/es | sed $'s/ /\\\n/g'
```

This command will generate a master list of words, expand those words with variants then if spaces are present transform them into new lines (OSX syntax).

The available aspell installed dictionaries can be shown with `aspell dicts`.

## Adding more languages

Not all the aspell supported languages are available with the default installation.

Fortunatelly there is an extended list available at https://ftp.gnu.org/gnu/aspell/dict/0index.html

These can be installed with commands similar to these. e.g.
```
wget https://ftp.gnu.org/gnu/aspell/dict/it/aspell-it-0.53-0.tar.bz2
tar xvjf aspell-it-0.53-0.tar.bz2
cd aspell-it-0.53-0
./configure
make install
```

After that the language should be available for the aspell command to generate the word list.


## Compiling

After the dictionaries have been installed. The binary can be compiled with:

```
go build palabreria.go
```

# Generating the passphrase

The binary comes with defaults, so it can be executed with:
```
./palabreria
```

Or arguments can be passed:

```
./palabreria -words=5 -separator=" "
```
