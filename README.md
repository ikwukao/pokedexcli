# Pokédex CLI

A command-line Pokédex application built with **Go** as part of the Boot.dev Backend Development curriculum.

This project interacts with the PokéAPI to explore Pokémon locations, catch Pokémon, inspect their details, and manage a personal Pokédex from the terminal.

> **Note:** This project is currently under active development. The documentation below will be expanded as new features are implemented.

---

## Motivation

The goal of this project is to strengthen practical Go development skills by building a real-world command-line application. Along the way, it explores concepts such as HTTP requests, JSON parsing, caching, CLI design, and clean software architecture.

---

## Features

* Interactive command-line interface (REPL)
* Explore Pokémon world locations
* View Pokémon found in each area
* Catch Pokémon
* Inspect caught Pokémon
* Manage a personal Pokédex
* API response caching for improved performance

---

## Tech Stack

* Go
* PokéAPI
* Go Standard Library (`net/http`, `encoding/json`, etc.)

---

## Project Structure

```text
.
├── main.go
├── internal/
├── pokecache/
├── pokeapi/
└── ...
```

> The project structure may evolve as development progresses.

---

## Quick Start

### Prerequisites

* Go 1.24 or later

### Clone the repository

```bash
git clone https://github.com/ikwukao/pokedexcli.git
cd pokedexcli
```

### Run the application

```bash
go run .
```

---

## Usage

Once the application is running, you will be able to use commands similar to:

```text
help
map
mapb
explore <location>
catch <pokemon>
inspect <pokemon>
pokedex
exit
```

Additional commands and examples will be added as the project grows.

---

## Roadmap

* [ ] Build the REPL
* [ ] Connect to the PokéAPI
* [ ] Browse map locations
* [ ] Explore areas
* [ ] Catch Pokémon
* [ ] Inspect Pokémon
* [ ] Implement API caching
* [ ] Improve error handling
* [ ] Add unit tests
* [ ] Polish documentation

---

## Learning Objectives

This project focuses on learning:

* Go programming
* REST API integration
* JSON parsing
* HTTP clients
* CLI application development
* Caching techniques
* Clean project organization
* Error handling

---

## Contributing

Contributions are not being accepted while this project is under active development as part of the Boot.dev curriculum.

Once the project is complete, contributions, suggestions, and bug reports may be welcomed. In the meantime, feel free to fork the repository and experiment on your own.

---

## License

This project is licensed under the MIT License.
